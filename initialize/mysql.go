package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/shansec/go-vue-admin/global"
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
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}
