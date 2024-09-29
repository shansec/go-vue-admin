package system

import (
	"github.com/shansec/go-vue-admin/dao/common/request"
	"github.com/shansec/go-vue-admin/dao/common/response"
	req "github.com/shansec/go-vue-admin/dao/request"
	res "github.com/shansec/go-vue-admin/dao/response"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/utils"
	commonVerify "github.com/shansec/go-vue-admin/verify/common"
	systemVerify "github.com/shansec/go-vue-admin/verify/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct{}

// CreateMenu
// @Summary 添加菜单
// @Description 添加菜单，返回添加结果
// @Tags SysBaseMenu
// @Produce json
// @Param   menuInfo body system.SysBaseMenu true "添加菜单"
// @Success 200 {object} response.Response{data=res.SysMenuResponse, msg=string}	"添加菜单,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "添加菜单失败"
// @Router /menu/createMenu [POST]
func (m *MenuApi) CreateMenu(c *gin.Context) {
	var menuInfo system.SysBaseMenu
	var err error
	if err = c.ShouldBindJSON(&menuInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = utils.Verify(menuInfo, systemVerify.CreateMenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := menuService.CreateMenuService(menuInfo)
	if err != nil {
		global.MAY_LOGGER.Error("创建菜单失败", zap.Error(err))
		response.FailWithMessage("创建菜单失败", c)
		return
	}
	response.OkWithDetailed(res.SysMenuResponse{Menu: menu}, "创建成功", c)
}

// DeleteMenu
// @Summary 删除菜单
// @Description 删除菜单，返回操作结果
// @Tags SysBaseMenu
// @Produce json
// @Param   menuInfo body system.SysBaseMenu true "删除菜单"
// @Success 200 {object} response.Response{msg=string}	"删除菜单,返回操作结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "删除菜单失败"
// @Router /menu/deleteMenu [DELETE]
func (m *MenuApi) DeleteMenu(c *gin.Context) {
	var menuInfo system.SysBaseMenu
	var err error
	_ = c.ShouldBindJSON(&menuInfo)

	err = utils.Verify(menuInfo, systemVerify.DeleteMenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = menuService.DeleteMenuService(&menuInfo); err != nil {
		global.MAY_LOGGER.Error("删除菜单失败", zap.Error(err))
		response.FailWithMessage("删除菜单失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除菜单成功", c)
}

// GetMenuList
// @Summary 分页获取菜单
// @Description 分页获取菜单
// @Tags SysBaseMenu
// @Produce json
// @Param   pageInfo body request.PageInfo true "分页获取菜单"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string}	"分页获取菜单"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "分页获取菜单失败"
// @Router /menu/getMenuList [POST]
func (m *MenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	var err error
	if err = c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(pageInfo, commonVerify.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := menuService.GetMenuListService(pageInfo)
	if err != nil {
		global.MAY_LOGGER.Error("分页获取菜单失败", zap.Error(err))
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

// UpdateMenu
// @Summary 修改菜单
// @Description 修改菜单，返回操作结果
// @Tags SysBaseMenu
// @Produce json
// @Param   menuInfo body system.SysBaseMenu true "修改菜单"
// @Success 200 {object} response.Response{msg=string}	"修改菜单,返回操作结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "修改菜单失败"
// @Router /menu/updateMenu [PUT]
func (m *MenuApi) UpdateMenu(c *gin.Context) {
	var menuInfo system.SysBaseMenu
	_ = c.ShouldBindJSON(&menuInfo)

	err := utils.Verify(menuInfo, systemVerify.UpdateMenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = menuService.UpdateMenuService(menuInfo)
	if err != nil {
		global.MAY_LOGGER.Error("修改菜单信息失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("修改成功", c)
}

// GetMenuTree
// @Summary 获取树状菜单
// @Description 获取树状菜单
// @Tags SysBaseMenu
// @Produce json
// @Success 200 {object} response.Response{data=response.NoPageResult, msg=string}	"获取树状菜单"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "获取树状菜单失败"
// @Router /menu/getMenuTree [POST]
func (m *MenuApi) GetMenuTree(c *gin.Context) {
	list, err := menuService.GetMenuTreeService()
	if err != nil {
		global.MAY_LOGGER.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取树状菜单失败", c)
		return
	}

	response.OkWithDetailed(response.NoPageResult{
		List: list,
	}, "获取成功", c)
}

// GetRoleMenu
// @Summary 获取当前登录角色菜单
// @Description 获取当前登录角色菜单
// @Tags SysBaseMenu
// @Produce json
// @Success 200 {object} response.Response{data=response.NoPageResult, msg=string}	"获取当前登录角色菜单"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "获取当前登录角色菜单失败"
// @Router /menu/getRoleMenu [POST]
func (m *MenuApi) GetRoleMenu(c *gin.Context) {
	roleId := utils.GetUserRoleId(c)
	if roleId == 0 {
		response.FailWithMessage("非法访问", c)
		return
	}

	menus, err := menuService.GetRoleMenuService(roleId)
	if err != nil {
		global.MAY_LOGGER.Error("获取角色菜单失败", zap.Error(err))
		response.FailWithMessage("获取角色菜单失败", c)
		return
	}
	response.OkWithDetailed(response.NoPageResult{List: menus}, "获取角色菜单成功", c)
}

// GetSpecialRoleMenu
// @Summary 获取指定角色菜单
// @Description 获取指定角色菜单
// @Tags SysBaseMenu
// @Produce json
// @Success 200 {object} response.Response{data=response.NoPageResult, msg=string}	"获取指定角色菜单"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "获取指定角色菜单失败"
// @Router /menu/getSpecialRoleMenu [POST]
func (m *MenuApi) GetSpecialRoleMenu(c *gin.Context) {
	var roleInfo req.GetSpecialRoleByID
	_ = c.ShouldBindJSON(&roleInfo)

	err := utils.Verify(roleInfo, systemVerify.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	menus, err := menuService.GetRoleMenuService(roleInfo.RoleId)
	if err != nil {
		global.MAY_LOGGER.Error("获取角色菜单失败", zap.Error(err))
		response.FailWithMessage("获取角色菜单失败", c)
		return
	}
	response.OkWithDetailed(response.NoPageResult{List: menus}, "获取角色菜单成功", c)
}
