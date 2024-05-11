package system

import (
	"github.com/gin-gonic/gin"
	v1 "github/shansec/go-vue-admin/api/v1"
)

type ApiRouter struct{}

func (a *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api")
	sysApi := v1.ApiGroupApp.SystemApiGroup.SysApi
	{
		apiRouter.POST("createApi", sysApi.CreateApi)   // 创建 api
		apiRouter.DELETE("deleteApi", sysApi.DeleteApi) // 删除 api
		apiRouter.PUT("updateApi", sysApi.UpdateApi)    // 更新 api
		apiRouter.GET("getApiList", sysApi.GetApiList)  // 获取 api 列表
	}
}
