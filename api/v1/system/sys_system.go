package system

import (
	"github.com/gin-gonic/gin"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"go.uber.org/zap"
)

type SystemConfigApi struct{}

// GetServerInfo
// @Tags System
// @Summary 获取服务器信息
// @Produce json
// @Success 200 {object} response.Response{msg=string}	"获取服务器信息"
// @Router /system/status [GET]
func (s *SystemConfigApi) GetServerInfo(c *gin.Context) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.MAY_LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}
