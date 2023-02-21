package initialize

import (
	"go-vue-admin/global"
	"go-vue-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// Gorm 初始化数据库并产生全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
	)
	if err != nil {
		global.GVA_LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOGGER.Info("register table success")
}
