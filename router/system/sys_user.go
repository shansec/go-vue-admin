package system

import (
	v1 "github/shansec/go-vue-admin/api/v1"
	"github/shansec/go-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user").Use(middleware.JwtAuth())
	baseApi := v1.ApiGroupAlias.SystemApiGroup.BaseApi
	{
		useRouter.POST("modifyPassword", baseApi.ModifyPassword)    // 修改密码
		useRouter.GET("getUserInfo", baseApi.GetUserInfo)           // 获取用户信息
		useRouter.POST("getUsersInfo", baseApi.GetUsersInfo)        // 获取用户列表
		useRouter.DELETE("delUserInfo", baseApi.DelUserInfo)        // 删除用户信息
		useRouter.POST("updateUserInfo", baseApi.UpdateUserInfo)    // 更改用户信息
		useRouter.PUT("updateUserStatus", baseApi.UpdateUserStatus) // 更改用户状态
	}
}
