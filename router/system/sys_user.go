package system

import (
	"github.com/gin-gonic/gin"
	v1 "github/shansec/go-vue-admin/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		useRouter.POST("register", baseApi.Register)                // 注册用户
		useRouter.POST("modifyPassword", baseApi.ModifyPassword)    // 修改密码
		useRouter.GET("getUserInfo", baseApi.GetUserInfo)           // 获取用户信息
		useRouter.POST("getUsersInfo", baseApi.GetUsersInfo)        // 获取用户列表
		useRouter.DELETE("delUserInfo", baseApi.DelUserInfo)        // 删除用户信息
		useRouter.POST("updateUserInfo", baseApi.UpdateUserInfo)    // 更改用户信息
		useRouter.PUT("updateUserStatus", baseApi.UpdateUserStatus) // 更改用户状态
	}
}
