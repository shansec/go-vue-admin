package main

import (
	"go-vue-admin/core"
	"go-vue-admin/global"
	"go-vue-admin/initialize"
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

	core.RunWindowsServer()
}
