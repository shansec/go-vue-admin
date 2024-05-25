package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemRes "github/shansec/go-vue-admin/model/system/response"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
)

type RoleApi struct{}

// CreateRole
// @Summary 添加角色
// @Description 添加角色，返回添加结果
// @Tags SysRole
// @Produce json
// @Param   createRoleInfo, roleInfo body system.SysRole true "添加角色"
// @Success 200 {object} response.Response{systemRes.SysRoleResponse, msg=string}	"添加角色,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "添加角色失败"
// @Router /role/createRole [POST]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var createRoleInfo, roleInfo system.SysRole
	_ = c.ShouldBindJSON(&createRoleInfo)
	var err error

	if err = utils.Verify(createRoleInfo, SystemVerify.CreateRoleVerify); err != nil {
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

	if err = utils.Verify(delRoleInfo, SystemVerify.DeleteRoleVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = roleService.DeleteRoleService(&delRoleInfo); err != nil {
		global.MAY_LOGGER.Error("删除角色失败", zap.Error(err))
		response.FailWithMessage("删除角色失败", c)
		return
	}
	response.OkWithMessage("删除角色成功", c)
}
