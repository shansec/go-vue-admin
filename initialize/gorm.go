package initialize

import (
	"os"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.MAY_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTable(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
	)
	if err != nil {
		global.MAY_LOGGER.Error("初始化数据库表失败", zap.Error(err))
		os.Exit(0)
	}
	global.MAY_LOGGER.Info("初始化数据库表成功")
}
