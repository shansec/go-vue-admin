package main

import (
	"go-vue-admin/core"
	"go-vue-admin/global"
	"go-vue-admin/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化 Viper
	global.GVA_VP = core.Viper()
	// 初始化日志
	global.GVA_LOGGER = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOGGER)
	// gorm 链接数据库
	global.GVA_DB = initialize.Gorm()

	core.RunWindowsServer()
}
