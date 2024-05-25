package system

import (
	"github.com/gin-gonic/gin"

	v1 "github/shansec/go-vue-admin/api/v1"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)    // 登录
		baseRouter.GET("captcha", baseApi.Captcha) // 验证码获取
	}
	return baseRouter
}
