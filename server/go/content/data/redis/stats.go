package redis

import (
	"github.com/hopeio/lemon/protobuf/errorcode"
	"github.com/liov/hoper/server/go/content/model"
)

func (d *ContentRedisDao) UserContentEdit(field string, value interface{}) error {
	ctxi := d
	ctx := ctxi.Context
	key := model.UserContentCountKey + ctxi.ID

	err := d.conn.HSet(ctx, key, field, value).Err()
	if err != nil {
		return ctxi.ErrorLog(errorcode.RedisErr, err, "RedisUserInfoEdit")
	}
	return nil
}
