package system

import (
	v1 "github/shansec/go-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role")
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("createRole", roleApi.CreateRole)   // 创建角色
		roleRouter.DELETE("deleteRole", roleApi.DeleteRole) // 删除角色
		roleRouter.PUT("updateRole", roleApi.UpdateRole)    // 更新角色
		roleRouter.POST("getRoleList", roleApi.GetRoleList) // 获取角色列表
		roleRouter.POST("setChildRole", roleApi.SetRole)    // 设置子角色
	}
}
