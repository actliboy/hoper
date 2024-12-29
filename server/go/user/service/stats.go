package service

import (
	"context"
	"github.com/hopeio/context/httpctx"
	dbi "github.com/hopeio/utils/dao/database/sql"
	"github.com/liov/hoper/server/go/protobuf/user"
	"github.com/liov/hoper/server/go/user/confdao"
	"github.com/liov/hoper/server/go/user/data"
	"github.com/liov/hoper/server/go/user/model"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hopeio/protobuf/errcode"
)

// 关注
func (u *UserService) Follow(ctx context.Context, req *user.FollowReq) (*emptypb.Empty, error) {
	ctxi := httpctx.FromContextValue(ctx)
	defer ctxi.StartSpanEnd("")()

	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}

	userDao := data.GetDBDao(ctxi, confdao.Dao.GORMDB.DB)
	exists, err := userDao.FollowExistsDB(req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, nil
	}
	err = userDao.Table(model.TableNameFollow).Create(&user.Follow{
		UserId:   req.Id,
		FollowId: auth.Id,
	}).Error
	if err != nil {
		return nil, ctxi.RespErrorLog(errcode.DBError, err, "Create")
	}
	return new(emptypb.Empty), nil
}

// 取消关注
func (u *UserService) DelFollow(ctx context.Context, req *user.FollowReq) (*user.BaseListRep, error) {
	ctxi := httpctx.FromContextValue(ctx)
	defer ctxi.StartSpanEnd("")()

	auth, err := auth(ctxi, true)
	if err != nil {
		return nil, err
	}

	userDao := data.GetDBDao(ctxi, confdao.Dao.GORMDB.DB)
	exists, err := userDao.FollowExistsDB(req.Id, auth.Id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	err = userDao.Table(model.TableNameFollow).Where("user_id = ? AND follow_id = ?"+dbi.WithNotDeleted, req.Id, auth.Id).
		UpdateColumn("deleted_at", ctxi.RequestAt.TimeString).Error
	if err != nil {
		return nil, ctxi.RespErrorLog(errcode.DBError, err, "Create")
	}
	return nil, nil
}
