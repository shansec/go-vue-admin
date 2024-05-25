package system

import (
	"github.com/gin-gonic/gin"

	v1 "github/shansec/go-vue-admin/api/v1"
)

type SystemConfigRouter struct{}

func (s *SystemConfigRouter) InitSystemConfigRouter(Router *gin.RouterGroup) {
	systemRouter := Router.Group("system")
	systemConfigApi := v1.ApiGroupApp.SystemApiGroup.SystemConfigApi
	{
		systemRouter.GET("getServerInfo", systemConfigApi.GetServerInfo) // 获取服务器状态
	}
}
