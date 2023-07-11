package confdao

import (
	"github.com/hopeio/dora/initialize/gormdb/postgres"
	"github.com/hopeio/dora/initialize/pebble"
	initredis "github.com/hopeio/dora/initialize/redis"
	"github.com/hopeio/dora/initialize/ristretto"
)

var Dao *dao = &dao{}

// dao dao.
type dao struct {
	GORMDB   postgres.DB
	PebbleDB pebble.DB
	// RedisPool Redis连接池
	Redis initredis.Redis
	Cache ristretto.Cache
}

func (d *dao) Init() {
	db := d.GORMDB
	db.Callback().Create().Remove("gorm:save_before_associations")
	db.Callback().Create().Remove("gorm:save_after_associations")
	db.Callback().Update().Remove("gorm:save_before_associations")
	db.Callback().Update().Remove("gorm:save_after_associations")

}
