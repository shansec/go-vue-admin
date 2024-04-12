package initialize

import (
	docs "github/shansec/go-vue-admin/docs"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/middleware"
	"github/shansec/go-vue-admin/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupAlias.System

	docs.SwaggerInfo.BasePath = global.MAY_CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	publicGroup := Router.Group(global.MAY_CONFIG.System.RouterPrefix)
	{
		// 健康检测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "OK")
		})
	}
	{
		systemRouter.InitBaseRouter(publicGroup)
	}

	PrivateGroup := Router.Group(global.MAY_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JwtAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitDeptRouter(PrivateGroup)
	}

	global.MAY_LOGGER.Info("router register sucess")
	return Router
}
