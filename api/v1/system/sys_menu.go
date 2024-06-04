package system

import (
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemRes "github/shansec/go-vue-admin/model/system/response"
	"github/shansec/go-vue-admin/utils"
	systemVerify "github/shansec/go-vue-admin/verify/system"

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
// @Success 200 {object} response.Response{data=systemRes.SysMenuResponse, msg=string}	"添加菜单,返回添加结果"
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
	response.OkWithDetailed(systemRes.SysMenuResponse{Menu: menu}, "创建成功", c)
}
