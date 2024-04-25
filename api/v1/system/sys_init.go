package system

import (
	"github.com/gin-gonic/gin"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system/request"
	"go.uber.org/zap"
)

type DBApi struct{}

func (d *DBApi) InitDB(c *gin.Context) {
	if global.MAY_DB != nil {
		global.MAY_LOGGER.Error("已存在数据库配置")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}

	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.MAY_LOGGER.Error("参数错误", zap.Error(err))
		response.FailWithMessage("参数错误", c)
		return
	}

}
