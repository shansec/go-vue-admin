package initialize

import (
	"github.com/gin-gonic/gin"
	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/middleware"
	"github/May-cloud/go-vue-admin/router"
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
