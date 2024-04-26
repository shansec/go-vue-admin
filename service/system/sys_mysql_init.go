package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github/shansec/go-vue-admin/config"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system/request"
	"github/shansec/go-vue-admin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/filepath"
)

type MysqlInit struct{}

func NewMysqlInit() *MysqlInit {
	return &MysqlInit{}
}

// WConfig mysql回写配置
func (m MysqlInit) WConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.MAY_CONFIG.System.DbType = "mysql"
	global.MAY_CONFIG.Mysql = c
	global.MAY_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	resultMap := utils.StructToMap(global.MAY_CONFIG)
	for k, v := range resultMap {
		global.MAY_VP.Set(k, v)
	}
	return global.MAY_VP.WriteConfig()
}

// CreateDB 创建数据库并初始化 mysql
func (m MysqlInit) CreateDB(ctx context.Context, config *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, errors.New("database type mismatch")
	}

	c := config.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)
	// 如果没有数据库名, 结束初始化数据
	if c.Dbname == "" {
		return ctx, nil
	}

	dsn := config.MysqlEmptyDsn()
	createDataSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if err = createDatabase(dsn, "mysql", createDataSql); err != nil {
		return nil, err
	} // 创建数据库

	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dns(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: true,    // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.MAY_CONFIG.AutoCode.Root, _ = filepath.Abs(".")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (m MysqlInit) TablesInit(ctx context.Context, inits inits) error {
	return createTables(ctx, inits)
}

func (m MysqlInit) DataInit(ctx context.Context, inits inits) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf("\n[%v] --> %v 的初始数据已存在!\n", "mysql", init.InitTableName())
			continue
		}
		if n, err := init.InitData(next); err != nil {
			color.Info.Printf("\n[%v] --> %v 的初始数据失败!\n err %v\n", "mysql", init.InitTableName(), err)
			return err
		} else {
			next = n
			color.Info.Printf("\n[%v] --> %v 的初始数据成功!\n", "mysql", init.InitTableName())
		}
	}
	color.Info.Printf("\n[%v] --> 的初始数据成功!\n", "mysql")
	return nil
}
