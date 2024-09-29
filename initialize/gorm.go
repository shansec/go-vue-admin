package initialize

import (
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
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
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysRole{},
		system.SysDept{},
		system.SysAutoCode{},
		system.SysDictionary{},
		system.SysDictionaryDetail{},
	)
	if err != nil {
		global.MAY_LOGGER.Error("初始化数据库表失败", zap.Error(err))
		os.Exit(0)
	}
	global.MAY_LOGGER.Info("初始化数据库表成功")
}
