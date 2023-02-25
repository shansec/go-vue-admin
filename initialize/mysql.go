package initialize

import (
	"go-vue-admin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.MAY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dns(),
		DefaultStringSize:         171,
		SkipInitializeWithVersion: false,
	}
	gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
}
