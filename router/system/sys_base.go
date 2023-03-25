package system

import (
	"github.com/gin-gonic/gin"
	v1 "github/May-cloud/go-vue-admin/api/v1"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupAlias.SystemApiGroup.BaseApi
	{

		baseRouter.POST("register", baseApi.Register) // 注册
		baseRouter.POST("login", baseApi.Login)       // 登录
	}
	return baseRouter
}
