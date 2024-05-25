package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
)

type DeptApi struct{}

// CreateDept
// @Summary 添加部门
// @Description 添加部门，返回添加结果
// @Tags SysDept
// @Produce json
// @Param   deptInfo body systemReq.Create true "添加部门"
// @Success 200 {object} response.Response{msg=string}	"添加部门,返回添加结果"
// @Failure 500 {object} response.Response   "添加部门失败"
// @Router /dept/createDept [POST]
func (d *DeptApi) CreateDept(c *gin.Context) {
	var deptInfo systemReq.Create
	_ = c.ShouldBindJSON(&deptInfo)

	// 验证输入数据
	if err := utils.Verify(deptInfo, SystemVerify.CreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 构建部门实体
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

	// 添加部门
	if err := deptService.EstablishDept(*dept); err != nil {
		global.MAY_LOGGER.Error("添加部门失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("添加部门成功", c)
	}
}

// GetDeptList
// @Summary 获取部门列表
// @Description 分页查询部门信息列表
// @Tags SysDept
// @Accept json
// @Produce json
// @Param   deptPageInfo body systemReq.GetDeptList true "部门列表查询参数"
// @Success 200 {object} response.PageResult{list=[]system.SysDept, msg=string}	"获取部门列表成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "获取部门列表失败"
// @Router /dept/getDeptList [POST]
func (d *DeptApi) GetDeptList(c *gin.Context) {
	var deptPageInfo systemReq.GetDeptList
	err := c.ShouldBindJSON(&deptPageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询部门列表
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
// @Summary 删除部门信息
// @Description 删除指定的部门及其所有下级部门信息
// @Tags SysDept
// @Accept json
// @Produce json
// @Param   deptInfo body system.SysDept true "待删除的部门信息"
// @Success 200 {object} response.Response{msg=string} "删除部门信息，返回操作结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "删除部门信息失败,请检查是否包含下级部门"
// @Router /dept/delDeptInfo [DELETE]
func (d *DeptApi) DelDeptInfo(c *gin.Context) {
	var deptInfo system.SysDept
	_ = c.ShouldBindJSON(&deptInfo)

	// 验证输入数据
	if err := utils.Verify(deptInfo, SystemVerify.DeleteVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 删除部门及其下级信息
	if err := deptService.DelDeptInformation(deptInfo); err != nil {
		global.MAY_LOGGER.Error("删除部门信息失败,请检查是否包含下级部门", zap.Error(err))
		response.FailWithMessage("删除部门信息失败,请检查是否包含下级部门", c)
	} else {
		response.OkWithMessage("删除部门信息成功", c)
	}
}

// UpdateDeptInfo
// @Summary 更新部门信息
// @Description 更新指定部门的详细信息
// @Tags SysDept
// @Accept json
// @Produce json
// @Param   dept body system.SysDept true "待更新的部门信息"
// @Success 200 {object} response.Response{msg=string} "更新部门信息,返回更新结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "更改部门信息失败"
// @Router /dept/updateDeptInfo [PUT]
func (d *DeptApi) UpdateDeptInfo(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)

	// 验证输入数据
	if err := utils.Verify(dept, SystemVerify.UpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 更新部门信息
	if err := deptService.UpdateDeptInformation(&dept); err != nil {
		global.MAY_LOGGER.Error("更改部门信息失败", zap.Error(err))
		response.FailWithMessage("更改部门信息失败", c)
	} else {
		response.OkWithMessage("更改部门信息成功", c)
	}
}
