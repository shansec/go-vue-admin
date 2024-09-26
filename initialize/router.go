package initialize

import (
	"github.com/gin-gonic/gin"

	docs "github.com/shansec/go-vue-admin/docs"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/middleware"
	"github.com/shansec/go-vue-admin/router"

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
		systemRouter.InitDBRouter(publicGroup)
	}

	PrivateGroup := Router.Group(global.MAY_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JwtAuth()).Use(middleware.CasbinAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitDeptRouter(PrivateGroup)
		systemRouter.InitSystemConfigRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitDictionaryRouter(PrivateGroup)
		systemRouter.InitDictionaryDetailRouter(PrivateGroup)
	}

	global.MAY_LOGGER.Info("router register success")
	return Router
}
