package system

import (
	v1 "github/shansec/go-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user")
	baseApi := v1.ApiGroupAlias.SystemApiGroup.BaseApi
	{
		useRouter.POST("modifyPassword", baseApi.ModifyPassword) // 修改密码
	}
}
