package main

import (
	"github.com/shansec/go-vue-admin/core"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/initialize"

	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Go-Vue-Admin Swagger API接口文档
// @version                     v1.0.0
// @description                 使用 gin + vue 进行开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /v1
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
