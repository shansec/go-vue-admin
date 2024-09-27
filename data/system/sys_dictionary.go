package system

import (
	"context"

	"github.com/pkg/errors"
	model "github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/service/system"
	"gorm.io/gorm"
)

const initOrderDictionary = initOrderUser + 1

type initDictionary struct{}

func init() {
	system.RegisterInit(initOrderDictionary, &initDictionary{})
}

func (d initDictionary) InitTableName() string {
	return model.SysDictionary{}.TableName()
}

func (d *initDictionary) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysDictionary{})

}

func (d *initDictionary) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	trueTmp := true
	entities := []model.SysDictionary{
		{Name: "性别", Type: "gender", Status: &trueTmp, Desc: "性别"},
		{Name: "数据库int类型", Type: "int", Status: &trueTmp, Desc: "int类型对应的数据库类型"},
		{Name: "数据库时间日期类型", Type: "time.Time", Status: &trueTmp, Desc: "数据库时间日期类型"},
		{Name: "数据库浮点型", Type: "float64", Status: &trueTmp, Desc: "数据库浮点型"},
		{Name: "数据库字符串", Type: "string", Status: &trueTmp, Desc: "数据库字符串"},
		{Name: "数据库bool类型", Type: "bool", Status: &trueTmp, Desc: "数据库bool类型"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysDictionary{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, d.InitTableName(), entities)
	return next, nil
}

func (d *initDictionary) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysDictionary{})
}

func (d *initDictionary) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var dict model.SysDictionary
	if errors.Is(db.Where("type = ?", "gender").First(&dict).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
