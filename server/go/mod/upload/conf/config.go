package conf

import (
	"github.com/actliboy/hoper/server/go/lib/initialize/cache/ristretto"
	"github.com/actliboy/hoper/server/go/lib/initialize/db"
	"github.com/actliboy/hoper/server/go/lib/initialize/log"
	"github.com/actliboy/hoper/server/go/lib/initialize/redis"
	"github.com/actliboy/hoper/server/go/lib/initialize/server"
	"github.com/actliboy/hoper/server/go/lib/utils/fs"
	"runtime"
)

var Conf = &config{}

type config struct {
	//自定义的配置
	Customize serverConfig
	Server    server.ServerConfig
	GORMDB    db.DatabaseConfig
	Redis     redis.Config
	Cache     ristretto.CacheConfig
	Log       log.LogConfig
}

func (c *config) Init() {
	if runtime.GOOS == "windows" {
	}

	c.Customize.UploadMaxSize = c.Customize.UploadMaxSize * 1024 * 1024
}

type serverConfig struct {
	Volume fs.Dir

	UploadDir      fs.Dir
	UploadMaxSize  int64
	UploadAllowExt []string

	LuosimaoVerifyURL string
	LuosimaoAPIKey    string

	QrCodeSaveDir fs.Dir //二维码保存路径
	FontSaveDir   fs.Dir //字体保存路径
}
