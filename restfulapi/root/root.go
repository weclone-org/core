package root

import (
	"github.com/Weclone-org/core/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootRouterInit(r *gin.Engine) {
	r.GET("/api/mode", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "success", "mode": config.Conf.Mode})
	})
}
