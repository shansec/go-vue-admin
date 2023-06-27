package system

import (
	v1 "github/May-cloud/go-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
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