package system

import (
	v1 "github/shansec/go-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.POST("createMenu", menuApi.CreateMenu)   // 添加菜单
		menuRouter.DELETE("deleteMenu", menuApi.DeleteMenu) // 删除菜单
		menuRouter.POST("getMenuList", menuApi.GetMenuList) // 分页获取菜单
		menuRouter.PUT("updateMenu", menuApi.UpdateMenu)    // 修改菜单
	}
}
