package restfulapi

import (
	"fmt"
	"github.com/Weclone-org/core/config"
	"github.com/Weclone-org/core/restfulapi/cdn"
	"github.com/Weclone-org/core/restfulapi/server"
	"github.com/gin-gonic/gin"
	"time"
)

func logFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] [RESTful-API] %s %s %s %d %s%s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		param.Method,
		param.Path,
		param.Latency,
		param.StatusCode,
		param.Request.Proto,
		param.ErrorMessage,
	)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logFormat))
	switch config.Conf.Mode {
	case "root":
		cdn.CDNRouterInit(r)
	case "server":
		server.ServerRouterInit(r)
	case "cdn":
		cdn.CDNRouterInit(r)
	}
	return r
}
