package system

import (
	"github.com/gin-gonic/gin"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
	"go.uber.org/zap"
)

type DeptApi struct{}

// CreateDept
// @Tags SysDept
// @Summary 添加部门
// @Produce json
// @Param   data body systemReq.Create true "添加部门"
// @Success 200 {object} response.Response{msg=string}	"添加部门,返回添加结果"
// @Router /dept/createDept [POST]
func (d *DeptApi) CreateDept(c *gin.Context) {
	var deptInfo systemReq.Create
	_ = c.ShouldBindJSON(&deptInfo)

	if err := utils.Verify(deptInfo, SystemVerify.CreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	dept := &system.SysDept{
		ParentId: deptInfo.ParentId,
		DeptPath: deptInfo.DeptPath,
		DeptName: deptInfo.DeptName,
		Sort:     deptInfo.Sort,
		Leader:   deptInfo.Leader,
		Phone:    deptInfo.Phone,
		Email:    deptInfo.Email,
		Status:   deptInfo.Status,
	}
	if err := deptService.EstablishDept(*dept); err != nil {
		global.MAY_LOGGER.Error("添加部门失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("添加部门成功", c)
	}
}

// GetDeptList
// @Tags SysDept
// @Summary 获取部门列表
// @Produce json
// @Param   data body systemReq.GetDeptList true "空"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string}	"获取部门列表,返回部门列表"
// @Router /dept/getDeptList [POST]
func (d *DeptApi) GetDeptList(c *gin.Context) {
	var deptPageInfo systemReq.GetDeptList
	err := c.ShouldBindJSON(&deptPageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if depts, total, err := deptService.GetDept(deptPageInfo); err != nil {
		global.MAY_LOGGER.Error("获取部门列表失败", zap.Error(err))
		response.FailWithMessage("获取部门列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     depts,
			Total:    total,
			Page:     deptPageInfo.Page,
			PageSize: deptPageInfo.PagSize,
		}, "获取部门列表成功", c)
	}
}

// DelDeptInfo
// @Tags SysDept
// @Summary 删除部门信息
// @Produce json
// @Param   data body system.SysDept true "删除部门信息"
// @Success 200 {object} response.Response{msg=string} "删除部门信息，返回操作结果"
// @Router /dept/delDeptInfo [DELETE]
func (d *DeptApi) DelDeptInfo(c *gin.Context) {
	var deptInfo system.SysDept
	_ = c.ShouldBindJSON(&deptInfo)

	if err := utils.Verify(deptInfo, SystemVerify.DeleteVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := deptService.DelDeptInformation(deptInfo); err != nil {
		global.MAY_LOGGER.Error("删除部门信息失败,请检查是否包含下级部门", zap.Error(err))
		response.FailWithMessage("删除部门信息失败,请检查是否包含下级部门", c)
	} else {
		response.OkWithMessage("删除部门信息成功", c)
	}
}

// UpdateDeptInfo
// @Tags SysDept
// @Summary 更新部门信息
// @Produce json
// @Param   data body system.SysDept true "更新部门信息"
// @Success 200 {object} response.Response{msg=string} "更新部门信息,返回更新结果"
// @Router /dept/updateDeptInfo [PUT]
func (d *DeptApi) UpdateDeptInfo(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)
	if err := utils.Verify(dept, SystemVerify.UpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deptService.UpdateDeptInformation(&dept); err != nil {
		global.MAY_LOGGER.Error("更改部门信息失败", zap.Error(err))
		response.FailWithMessage("更改部门信息失败", c)
	} else {
		response.OkWithMessage("更改部门信息成功", c)
	}
}
