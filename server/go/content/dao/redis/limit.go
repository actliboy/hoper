package redis

import (
	"time"

	"github.com/actliboy/hoper/server/go/content/confdao"
	"github.com/go-redis/redis/v8"
	"github.com/hopeio/dora/protobuf/errorcode"
	timei "github.com/hopeio/dora/utils/time"
)

var limitErr = errorcode.TimeTooMuch.Message("您的操作过于频繁，请先休息一会儿。")

func (d *ContentRedisDao) Limit(l *confdao.Limit) error {
	ctxi := d
	ctx := ctxi.Context
	minuteKey := l.MinuteLimitKey + ctxi.ID
	dayKey := l.DayLimitKey + ctxi.ID

	var minuteIntCmd, dayIntCmd *redis.IntCmd
	_, err := d.conn.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		minuteIntCmd = pipe.Incr(ctx, minuteKey)
		dayIntCmd = pipe.Incr(ctx, dayKey)
		return nil
	})
	if err != nil {
		return ctxi.ErrorLog(errorcode.RedisErr, err, "Incr")
	}

	if minuteIntCmd.Val() > l.MinuteLimitCount || dayIntCmd.Val() > l.DayLimitCount {
		return limitErr
	}
	var minuteDurationCmd, dayDurationCmd *redis.DurationCmd
	_, err = d.conn.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		minuteDurationCmd = pipe.PTTL(ctx, minuteKey)
		dayDurationCmd = pipe.PTTL(ctx, dayKey)
		return nil
	})
	if err != nil {
		return ctxi.ErrorLog(errorcode.RedisErr, err, "PTTL")
	}

	_, err = d.conn.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		if minuteDurationCmd.Val() < 0 {
			pipe.Expire(ctx, minuteKey, time.Minute)
		}
		if dayDurationCmd.Val() < 0 {
			pipe.Expire(ctx, dayKey, timei.TimeDay)
		}
		return nil
	})
	if err != nil {
		return ctxi.ErrorLog(errorcode.RedisErr, err, "Expire")
	}
	return nil
}
