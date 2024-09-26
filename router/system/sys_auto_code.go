package system

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/shansec/go-vue-admin/api/v1"
)

type AutoCodeRouter struct{}

func (a *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autocode")
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.POST("createPackage", autoCodeApi.CreatePackage)   // 创建包
		autoCodeRouter.POST("getPackageList", autoCodeApi.GetPackageList) // 查询包
		autoCodeRouter.DELETE("delPackageInfo", autoCodeApi.DelPackage)   // 删除包
	}
}
