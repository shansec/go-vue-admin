package request

import (
	"fmt"
	"github/shansec/go-vue-admin/config"
)

type InitDB struct {
	DbType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		Path:         i.Host,
		Port:         i.Port,
		Dbname:       i.DbName,
		Username:     i.UserName,
		Password:     i.Password,
		MaxOpenConns: 10,
		MaxIdleConns: 100,
		LogMode:      "error",
		Config:       "charset=utf8mb4&parseTime=True&loc=Local",
	}
}

// MysqlEmptyDsn 数据库链接拼接
func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}
