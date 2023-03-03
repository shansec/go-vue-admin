package initialize

import (
	"github/May-cloud/go-vue-admin/global"
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

	// 暂时解决报错
	var test *gorm.DB
	return test
}
