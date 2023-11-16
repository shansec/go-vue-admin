package system

import (
	"github.com/gin-gonic/gin"
	v1 "github/shansec/go-vue-admin/api/v1"
)

type DeptRouter struct{}

func (d *DeptRouter) InitDeptRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("dept")
	deptApi := v1.ApiGroupApp.SystemApiGroup.DeptApi
	{
		deptRouter.POST("createDept", deptApi.CreateDept)        // 添加部门
		deptRouter.POST("getDeptList", deptApi.GetDeptList)      // 部门列表
		deptRouter.DELETE("delDeptInfo", deptApi.DelDeptInfo)    // 删除部门
		deptRouter.PUT("updateDeptInfo", deptApi.UpdateDeptInfo) // 修改部门
	}
}
