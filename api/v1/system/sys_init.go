package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/dao/common/response"
	"github/shansec/go-vue-admin/dao/request"
	"github/shansec/go-vue-admin/global"
)

type DBApi struct{}

// InitDB
// @Summary 初始化数据
// @Description 初始化系统数据
// @Tags InitDB
// @Success 200 {object} response.Response{msg=string} "初始化系统数据"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response "初始化失败"
// @Router /init/initDB [POST]
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
	if err := initDbService.InitDBService(dbInfo); err != nil {
		global.MAY_LOGGER.Error("初始化失败", zap.Error(err))
		response.FailWithMessage("初始化失败", c)
		return
	}
	response.OkWithMessage("初始化成功", c)
}

// CheckDB
// @Summary 检查是否已初始化
// @Description 检查是否已初始化
// @Tags InitDB
// @Router /init/checkDB [POST]
func (d *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		isInited = false
	)

	if global.MAY_DB != nil {
		message = "数据库无需初始化"
		isInited = true
	}
	global.MAY_LOGGER.Info(message)
	response.OkWithDetailed(gin.H{"isInited": isInited}, message, c)
}
