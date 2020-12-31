package gin_build

import (
	"github.com/gin-gonic/gin"
	httpi "github.com/liov/hoper/go/v2/utils/net/http"
	"github.com/liov/hoper/go/v2/utils/net/http/gin/handler"
)

func Http(confPath,apiPath string,ginHandle func(engine *gin.Engine)) *gin.Engine {
	//openapi
	r := gin.New()
	WithConfiguration(r, confPath)
	//r.Use(gin.Logger())
	/*logger := (&log.Config{Development: initialize.InitConfig.Env == initialize.PRODUCT}).NewLogger()
	middleware.SetLog(r, logger, false)*/
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	r.Any("/debug/*path", handler.FromStd(httpi.Debug()))
	if ginHandle != nil {
		ginHandle(r)
	}
	return r
}
