package system

import (
	"github.com/gin-gonic/gin"
	v1 "github/May-cloud/go-vue-admin/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user")
	baseApi := v1.ApiGroupAlias.SystemApiGroup.BaseApi
	{
		useRouter.POST("modifyPassword", baseApi.ModifyPassword) // 修改密码
	}
}
