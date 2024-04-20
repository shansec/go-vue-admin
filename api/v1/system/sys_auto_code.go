package system

import (
	"github.com/gin-gonic/gin"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/model/system/request"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
	"go.uber.org/zap"
	"strings"
)

type AutoCodeApi struct{}

// CreatePackage
// @Summary 自动创建代码包
// @Tags SysAutoCode
// @Accept json
// @Produce json
// @Param autoCode body system.SysAutoCode true "自动创建代码包"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /autocode/createPackage [POST]
func (a *AutoCodeApi) CreatePackage(c *gin.Context) {
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&autoCode)

	// 验证输入数据
	if err := utils.Verify(autoCode, SystemVerify.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 检查包名合法性，防止路径穿越
	if strings.Contains(autoCode.PackageName, "\\") || strings.Contains(autoCode.PackageName, "/") {
		response.FailWithMessage("包名不合法", c)
		return
	}

	err := autoCodeService.CreateAutoCode(&autoCode)
	if err != nil {
		global.MAY_LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// GetPackageList
// @Summary 获取创建的包列表
// @Description 分页查询创建的包列表
// @Tags SysAutoCode
// @Accept json
// @Produce json
// @Param   packageInfo body request.GetPackageList true "创建的包列表查询参数"
// @Success 200 {object} response.PageResult{list=[]system.SysAutoCode, msg=string}	"获取创建的包列表成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "获取创建的包列表失败"
// @Router /autocode/getPackageList [POST]
func (a *AutoCodeApi) GetPackageList(c *gin.Context) {
	var packageInfo request.GetPackageList
	_ = c.ShouldBindJSON(&packageInfo)

	if packages, total, err := autoCodeService.GetPackages(packageInfo); err != nil {
		global.MAY_LOGGER.Error("获取包列表失败", zap.Error(err))
		response.FailWithMessage("获取包列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     packages,
			Total:    total,
			Page:     packageInfo.Page,
			PageSize: packageInfo.PagSize,
		}, "获取包列表成功", c)
	}
}

// DelPackage
// @Summary 删除创建的包信息
// @Description 删除指定的创建的包信息
// @Tags SysAutoCode
// @Accept json
// @Produce json
// @Param   autoCode body system.SysAutoCode true "待删除的创建的包信息"
// @Success 200 {object} response.Response{msg=string} "删除创建的包，返回操作结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "删除创建的包失败"
// @Router /autocode/delPackageInfo [DELETE]
func (a *AutoCodeApi) DelPackage(c *gin.Context) {
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&autoCode)
	if err := autoCodeService.DelPackageInfo(&autoCode); err != nil {
		global.MAY_LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
