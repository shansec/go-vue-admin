package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system/request"
	"gorm.io/gorm"
	"sort"
)

/* ———— Init ———— */
type orderInit struct {
	order uint
	SubInitializer
}

type inits []*orderInit

type SubInitializer interface {
	InitTableName() string
	MigrateTable(ctx context.Context) (cont context.Context, err error)
	InitData(ctx context.Context) (cont context.Context, err error)
	TableInited(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

type DBInitHandler interface {
	// CreateDB 创建数据库
	CreateDB(ctx context.Context, db *request.InitDB) (context.Context, error)
	// WConfig 写设置
	WConfig(ctx context.Context) error
	// TablesInit 表初始化
	TablesInit(ctx context.Context, init inits) error
	// DataInit 数据初始化
	DataInit(ctx context.Context, init inits) error
}

var (
	initalizer inits
	cache      map[string]*orderInit
)

// RegisterInit 数据初始化
func RegisterInit(order uint, s SubInitializer) {
	if initalizer == nil {
		initalizer = inits{}
	}
	if cache == nil {
		cache = map[string]*orderInit{}
	}
	name := s.InitTableName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("name existed %s", name))
	}
	init := orderInit{order, s}
	initalizer = append(initalizer, &init)
	cache[name] = &init
}

func (i inits) Len() int {
	return len(i)
}

func (i inits) Less(k, j int) bool {
	return i[k].order < i[j].order
}

func (i inits) Swap(k, j int) {
	i[k], i[j] = i[j], i[k]
}

/* ———— Service ———— */
type InitService struct{}

func (initService *InitService) InitDB(config request.InitDB) (err error) {
	ctx := context.TODO()
	if len(initalizer) == 0 {
		return errors.New("无可用初始化！")
	}
	sort.Sort(&initalizer)
	var dbInitHandler DBInitHandler = NewMysqlInit()
	ctx = context.WithValue(ctx, "dbtype", "mysql")

	next, err := dbInitHandler.CreateDB(ctx, &config)
	if err != nil {
		return err
	}
	db := next.Value("db").(*gorm.DB)
	global.MAY_DB = db

	if err := dbInitHandler.TablesInit(next, initalizer); err != nil {
		return err
	}
	if err := dbInitHandler.DataInit(next, initalizer); err != nil {
		return err
	}
	if err := dbInitHandler.WConfig(next); err != nil {
		return err
	}
	initalizer = inits{}
	cache = map[string]*orderInit{}
	return nil
}

// createDatabase 创建数据库
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// createTables 创建表
func createTables(ctx context.Context, inits inits) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableInited(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}
