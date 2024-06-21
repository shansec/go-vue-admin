package system

import (
	systemRes "github/shansec/go-vue-admin/dao/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/dao/common/request"
	"github/shansec/go-vue-admin/dao/common/response"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/utils"
	commonVerify "github/shansec/go-vue-admin/verify/common"
	systemVerify "github/shansec/go-vue-admin/verify/system"
)

type RoleApi struct{}

// CreateRole
// @Summary 添加角色
// @Description 添加角色，返回添加结果
// @Tags SysRole
// @Produce json
// @Param   createRoleInfo body system.SysRole true "添加角色"
// @Success 200 {object} response.Response{data=systemRes.SysRoleResponse, msg=string}	"添加角色,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "添加角色失败"
// @Router /role/createRole [POST]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var createRoleInfo, roleInfo system.SysRole
	_ = c.ShouldBindJSON(&createRoleInfo)
	var err error

	if err = utils.Verify(createRoleInfo, systemVerify.CreateRoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	roleInfo, err = roleService.CreateRoleService(createRoleInfo)
	if err != nil {
		global.MAY_LOGGER.Error("创建角色失败", zap.Error(err))
		response.FailWithMessage("创建角色失败", c)
		return
	}

	err = casbinService.FreshCasbin()
	if err != nil {
		global.MAY_LOGGER.Error("创建角色成功，权限刷新失败", zap.Error(err))
		response.FailWithMessage("创建角色成功，权限刷新失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysRoleResponse{Role: roleInfo}, "创建角色成功", c)
}

// DeleteRole
// @Summary 删除角色
// @Description 删除角色，返回添加结果
// @Tags SysRole
// @Produce json
// @Param   delRoleInfo body system.SysRole true "删除角色"
// @Success 200 {object} response.Response{msg=string}	"删除角色,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "删除角色失败"
// @Router /role/deleteRole [DELETE]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	var delRoleInfo system.SysRole
	var err error
	if err = c.ShouldBindJSON(&delRoleInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = utils.Verify(delRoleInfo, systemVerify.DeleteRoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = roleService.DeleteRoleService(&delRoleInfo); err != nil {
		global.MAY_LOGGER.Error("删除角色失败", zap.Error(err))
		response.FailWithMessage("删除角色失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("删除角色成功", c)
}

// UpdateRole
// @Summary 更新角色
// @Description 更新角色，返回添加结果
// @Tags SysRole
// @Produce json
// @Param   updateRole body system.SysRole true "更新角色"
// @Success 200 {object} response.Response{data=systemRes.SysRoleResponse, msg=string}	"更新角色,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "更新角色失败"
// @Router /role/updateRole [PUT]
func (r *RoleApi) UpdateRole(c *gin.Context) {
	var updateRole system.SysRole
	var err error
	if err = c.ShouldBindJSON(&updateRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.Verify(updateRole, systemVerify.UpdateRoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	role, err := roleService.UpdateRoleService(updateRole)
	if err != nil {
		global.MAY_LOGGER.Error("更新角色失败", zap.Error(err))
		response.FailWithMessage("更新角色失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysRoleResponse{Role: role}, "更新角色成功", c)
}

// GetRoleList
// @Summary 分页获取角色列表
// @Description 分页获取角色列表
// @Tags SysRole
// @Produce json
// @Param   pageInfo body request.PageInfo true "分页获取角色列表"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string}	"分页获取角色列表"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "分页获取角色列表失败"
// @Router /role/getRoleList [POST]
func (r *RoleApi) GetRoleList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, commonVerify.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := roleService.GetRoleListService(pageInfo)
	if err != nil {
		global.MAY_LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SetRole
// @Summary 设置子角色
// @Description 设置子角色
// @Tags SysRole
// @Produce json
// @Param   role body system.SysRole true "设置子角色"
// @Success 200 {object} response.Response{msg=string}	"设置子角色"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "设置子角色失败"
// @Router /role/setChildRole [POST]
func (r *RoleApi) SetRole(c *gin.Context) {
	var role system.SysRole
	err := c.ShouldBindJSON(&role)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(role, systemVerify.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = roleService.SetRoleService(role)
	if err != nil {
		global.MAY_LOGGER.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
