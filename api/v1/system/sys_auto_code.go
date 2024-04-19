package system

import (
	"github.com/gin-gonic/gin"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
	"go.uber.org/zap"
	"strings"
)

type AutoCodeApi struct{}

func (a *AutoCodeApi) CreatePackage(c *gin.Context) {
	var autoCode system.SysAutoCode
	_ = c.ShouldBindJSON(&autoCode)
	if err := utils.Verify(autoCode, SystemVerify.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 防止路径穿越
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
