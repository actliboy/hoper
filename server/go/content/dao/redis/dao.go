package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/hopeio/tailmon/context/http_context"
	"github.com/hopeio/tailmon/utils/log"
)

type ContentRedisDao struct {
	*http_context.Context
	conn redis.Cmdable
}

func GetDao(ctx *http_context.Context, redis redis.Cmdable) *ContentRedisDao {
	if ctx == nil {
		log.Fatal("ctx can't nil")
	}
	return &ContentRedisDao{ctx, redis}
}
