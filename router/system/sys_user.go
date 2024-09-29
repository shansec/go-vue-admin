package system

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/shansec/go-vue-admin/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		useRouter.POST("register", userApi.Register)                // 注册用户
		useRouter.POST("modifyPassword", userApi.ModifyPassword)    // 修改密码
		useRouter.GET("getUserInfo", userApi.GetUserInfo)           // 获取用户信息
		useRouter.POST("getUsersInfo", userApi.GetUsersInfo)        // 获取用户列表
		useRouter.DELETE("delUserInfo", userApi.DelUserInfo)        // 删除用户信息
		useRouter.POST("updateUserInfo", userApi.UpdateUserInfo)    // 更改用户信息
		useRouter.PUT("updateUserStatus", userApi.UpdateUserStatus) // 更改用户状态
	}
}
