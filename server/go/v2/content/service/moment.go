package service

import (
	"context"
	"fmt"
	"github.com/liov/hoper/go/v2/content/client"
	"github.com/liov/hoper/go/v2/protobuf/utils/request"
	httpi "github.com/liov/hoper/go/v2/utils/net/http"
	stringsi "github.com/liov/hoper/go/v2/utils/strings"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"google.golang.org/grpc/metadata"
	"net/http"
	"unicode/utf8"

	"github.com/liov/hoper/go/v2/content/conf"
	"github.com/liov/hoper/go/v2/content/dao"
	"github.com/liov/hoper/go/v2/content/model"
	"github.com/liov/hoper/go/v2/protobuf/content"
	"github.com/liov/hoper/go/v2/protobuf/user"
	"github.com/liov/hoper/go/v2/protobuf/utils/empty"
	"github.com/liov/hoper/go/v2/protobuf/utils/errorcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type MomentService struct {
	content.UnimplementedMomentServiceServer
}

func (*MomentService) Service() (describe, prefix string, middleware []http.HandlerFunc) {
	return "瞬间相关", "/api/moment", nil
}

func (*MomentService) Info(ctx context.Context, req *request.Object) (*content.Moment, error) {
	ctxi, span := user.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	_, err := ctxi.GetAuthInfo(AuthWithUpdate)
	if err != nil {
		return nil, err
	}
	contentDao := dao.GetDao(ctxi)
	err = contentDao.LimitRedis(dao.Dao.Redis, &conf.Conf.Customize.Moment.Limit)
	if err != nil {
		return nil, err
	}
	db := dao.Dao.GetDB(ctxi.Logger)
	var moment content.Moment
	err = db.Table(model.MomentTableName).
		Where(`id = ?`, req.Id).First(&moment).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError, err, "First")
	}
	contentTags, err := contentDao.GetContentTagDB(db, content.ContentMoment, []uint64{moment.Id})
	if err != nil {
		return nil, err
	}
	var tags = make([]*content.TinyTag, len(contentTags))
	for i := range contentTags {
		tags[i] = &contentTags[i].TinyTag
	}
	moment.Tags = tags

	_,comments,err:=contentDao.GetCommentsDB(db,content.ContentMoment,req.Id,0,0,0)
	if err != nil{
		return nil,err
	}
	moment.Comments = comments
	var userIds []uint64

	for i := range comments{
		userIds = append(userIds,comments[i].UserId)
		userIds = append(userIds,comments[i].RecvId)
	}
	userIds = append(userIds,moment.UserId)
	_,span = trace.StartSpan(ctx,"UserClient.BaseList")
	ctx = metadata.AppendToOutgoingContext(ctx,httpi.GrpcTraceBin,stringsi.ToString(propagation.Binary(span.SpanContext())),
		httpi.GrpcInternal,httpi.GrpcInternal)
	userList,err:=client.UserClient.BaseList(ctx,&user.BaseListReq{Ids:userIds})
	if err != nil{
		return nil,err
	}
	span.End()
	var m = make(map[uint64]*user.UserBaseInfo)
	for _,u:=range userList.List{
		m[u.Id] = u
	}
	for i := range comments{
		comments[i].RecvUser = m[comments[i].RecvId]
		comments[i].User = m[comments[i].UserId]
	}
	moment.User = m[moment.UserId]
	return &moment, nil
}

func (m *MomentService) Add(ctx context.Context, req *content.AddMomentReq) (*empty.Empty, error) {

	if utf8.RuneCountInString(req.Content) < conf.Conf.Customize.Moment.MaxContentLen {
		return nil, errorcode.InvalidArgument.Message(fmt.Sprintf("文章内容不能小于%d个字", conf.Conf.Customize.Moment.MaxContentLen))
	}

	ctxi, span := user.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := ctxi.GetAuthInfo(AuthWithUpdate)
	if err != nil {
		return nil, err
	}
	contentDao := dao.GetDao(ctxi)
	err = contentDao.LimitRedis(dao.Dao.Redis, &conf.Conf.Customize.Moment.Limit)
	if err != nil {
		return nil, err
	}

	req.UserId = auth.Id
	db := dao.Dao.GetDB(ctxi.Logger)
	/*	var count int64
		db.Table(`mood`).Where(`name = ?`, req.MoodName).Count(&count)
		if count == 0 {
			return nil, errorcode.ParamInvalid.Message("心情不存在")
		}*/
	tags, err := contentDao.GetTagsDB(db, req.Tags)
	if err != nil {
		return nil, err
	}
	req.UserId = auth.Id
	err = db.Transaction(func(tx *gorm.DB) error {
		if req.Permission == 0 {
			req.Permission = content.ViewPermissionAll
		}
		err = tx.Table(model.MomentTableName).Create(req).Error
		if err != nil {
			return ctxi.ErrorLog(err, err, "tx.CreateReq")
		}
		err = tx.Table(model.ContentExtTableName).Create(&model.ContentExt{
			Type:  content.ContentMoment,
			RefId: req.Id,
		}).Error
		if err != nil {
			return ctxi.ErrorLog(err, err, "tx.CreateReq")
		}
		var contentTags []model.ContentTag
		var noExist []content.Tag
	Loop:
		for i := range req.Tags {
			// 性能可以优化
			for j := range tags {
				if req.Tags[i] == tags[j].Name {
					contentTags = append(contentTags, model.ContentTag{
						Type:  content.ContentMoment,
						RefId: req.Id,
						TagId: tags[j].Id,
					})
					continue Loop
				}
			}
			noExist = append(noExist, content.Tag{Name: req.Tags[i], UserId: auth.Id})
		}
		if len(noExist) == 1 {
			if err = tx.Create(&noExist[1]).Error; err != nil {
				return ctxi.ErrorLog(err, err, "db.CreateNoExist")
			}
		}
		if len(noExist) > 1 {
			if err = tx.Create(&noExist).Error; err != nil {
				return ctxi.ErrorLog(err, err, "db.CreateNoExist")
			}
		}
		for i := range noExist {
			contentTags = append(contentTags, model.ContentTag{
				Type:  content.ContentMoment,
				RefId: req.Id,
				TagId: noExist[i].Id,
			})
		}
		if err = tx.Create(&contentTags).Error; err != nil {
			return ctxi.ErrorLog(err, err, "db.CreateContentTags")
		}
		return nil
	})
	if err != nil {
		if err != errorcode.DBError {
			return nil, ctxi.ErrorLog(errorcode.DBError, err, "Transaction")
		}
		return nil, errorcode.DBError
	}
	return nil, nil
}
func (*MomentService) Edit(context.Context, *content.AddMomentReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}

func (*MomentService) List(ctx context.Context, req *content.MomentListReq) (*content.MomentListRep, error) {
	ctxi, span := user.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := ctxi.GetAuthInfo(AuthWithUpdate)
	if err != nil {
		return nil, err
	}
	contentDao := dao.GetDao(ctxi)
	err = contentDao.LimitRedis(dao.Dao.Redis, &conf.Conf.Customize.Moment.Limit)
	if err != nil {
		return nil, err
	}

	db := dao.Dao.GetDB(ctxi.Logger)

	total, moments, err := contentDao.GetMomentListDB(db, req)
	if err != nil {
		return nil, err
	}
	var m = make(map[uint64]*content.Moment)
	var ids []uint64
	for i := range moments {
		ids = append(ids, moments[i].Id)
		m[moments[i].Id] = moments[i]
	}
	// tag
	tags, err := contentDao.GetContentTagDB(db, content.ContentMoment, ids)
	if err != nil {
		return nil, err
	}

	for i := range tags {
		if moment, ok := m[tags[i].RefId]; ok {
			moment.Tags = append(moment.Tags, &tags[i].TinyTag)
		}
	}

	//like
	if auth.Id != 0 {
		likes, err := contentDao.GetContentActionDB(db, content.ActionLike, content.ContentMoment, ids, auth.Id)
		if err != nil {
			return nil, err
		}
		for i := range likes {
			if moment, ok := m[likes[i].RefId]; ok {
				if likes[i].Action == content.ActionLike {
					moment.LikeId = likes[i].LikeId
				}
				if likes[i].Action == content.ActionUnlike {
					moment.UnlikeId = likes[i].LikeId
				}
			}
		}
		collects, err := contentDao.GetContentActionDB(db, content.ActionCollect, content.ContentMoment, ids, auth.Id)
		if err != nil {
			return nil, err
		}
		for i := range collects {
			if moment, ok := m[likes[i].RefId]; ok {
				moment.Collect = true
			}
		}
	}

	return &content.MomentListRep{
		Total: total,
		List:  moments,
	}, nil
}

func (*MomentService) Delete(ctx context.Context, req *request.Object) (*empty.Empty, error) {
	ctxi, span := user.CtxFromContext(ctx).StartSpan("")
	defer span.End()
	auth, err := ctxi.GetAuthInfo(AuthWithUpdate)
	if err != nil {
		return nil, err
	}
	contentDao := dao.GetDao(ctxi)
	err = contentDao.LimitRedis(dao.Dao.Redis, &conf.Conf.Customize.Moment.Limit)
	if err != nil {
		return nil, err
	}
	db := dao.Dao.GetDB(ctxi.Logger)
	err = contentDao.DelByAuthDB(db, model.MomentTableName, req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
