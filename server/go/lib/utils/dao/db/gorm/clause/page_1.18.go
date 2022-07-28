//go:build go1.18

package clausei

import (
	_type "github.com/actliboy/hoper/server/go/lib/utils/dao/db/gorm/type"
	"github.com/actliboy/hoper/server/go/lib/utils/def/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// unsupported data,完全不可用
// deprecated: 不可用
func List[T any, O _type.Ordered](db *gorm.DB, req *_type.ListReq[O]) ([]T, error) {
	var models []T

	clauses := append((*PageSortReq)(&req.PageSortReq).Clause(), (*RangeReq[O])(req.RangeReq).Clause())
	err := db.Clauses(clauses...).Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func ListClause[O _type.Ordered](req *_type.ListReq[O]) []clause.Expression {
	return append((*PageSortReq)(&req.PageSortReq).Clause(), (*RangeReq[O])(req.RangeReq).Clause())
}

type PageSortReq request.PageSortReq

func (req *PageSortReq) Clause() []clause.Expression {
	if req.PageNo == 0 && req.PageSize == 0 {
		return []clause.Expression{new(EmptyClause)}
	}
	if req.SortReq == nil || req.SortReq.SortField == "" {
		return []clause.Expression{Page(req.PageNo, req.PageSize)}
	}

	return []clause.Expression{Sort(req.SortField, req.SortType), Page(req.PageNo, req.PageSize)}
}

type ListReq[T _type.Ordered] _type.ListReq[T]

func (req *ListReq[O]) Clause() []clause.Expression {
	psqc := (*PageSortReq)(&req.PageSortReq).Clause()
	rqc := (*RangeReq[O])(req.RangeReq).Clause()
	if psqc != nil && rqc != nil {
		return append((*PageSortReq)(&req.PageSortReq).Clause(), (*RangeReq[O])(req.RangeReq).Clause())
	}
	if rqc == nil {
		return psqc
	}
	if rqc != nil {
		return []clause.Expression{rqc}
	}
	return []clause.Expression{new(EmptyClause)}
}
