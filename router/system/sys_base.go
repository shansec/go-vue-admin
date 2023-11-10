package system

import (
	v1 "github/shansec/go-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("register", baseApi.Register) // 注册
		baseRouter.POST("login", baseApi.Login)       // 登录
		baseRouter.GET("captcha", baseApi.Captcha)    // 验证码获取
	}
	return baseRouter
}
