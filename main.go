package main

import (
	"github/May-cloud/go-vue-admin/core"
	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化 Viper
	global.MAY_VP = core.Viper()
	// 初始化日志
	global.MAY_LOGGER = core.Zap()
	zap.ReplaceGlobals(global.MAY_LOGGER)
	// gorm 链接数据库
	global.MAY_DB = initialize.Gorm()
	if global.MAY_DB != nil {
		initialize.RegisterTable(global.MAY_DB)

		db, _ := global.MAY_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
