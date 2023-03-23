package system

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//useRouter := Router.Group("user")
	//baseApi := v1.ApiGroupAlias.SystemApiGroup.BaseApi

}
