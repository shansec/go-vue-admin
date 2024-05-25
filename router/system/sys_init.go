package system

import (
	"github.com/gin-gonic/gin"

	v1 "github/shansec/go-vue-admin/api/v1"
)

type InitRouter struct{}

func (i *InitRouter) InitDBRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initDB", dbApi.InitDB)   // 系统初始化
		initRouter.POST("checkDB", dbApi.CheckDB) // 检查初始化
	}
}
