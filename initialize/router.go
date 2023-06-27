package initialize

import (
	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/middleware"
	"github/May-cloud/go-vue-admin/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupAlias.System

	publicGroup := Router.Group("")
	{
		// 健康检测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "OK")
		})
	}
	{
		systemRouter.InitBaseRouter(publicGroup)
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JwtAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}

	global.MAY_LOGGER.Info("router register sucess")
	return Router
}
