package system

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	model "github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/service/system"
	"gorm.io/gorm"
)

const initOrderDictionaryDetail = initOrderDictionary + 1

type initDictionaryDetail struct{}

func init() {
	system.RegisterInit(initOrderDictionaryDetail, &initDictionaryDetail{})
}

func (d *initDictionaryDetail) InitTableName() string {
	return model.SysDictionaryDetail{}.TableName()
}

func (d *initDictionaryDetail) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysDictionaryDetail{})

}

func (d *initDictionaryDetail) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	dictionaries, ok := ctx.Value(initDictionary{}.InitTableName()).([]model.SysDictionary)
	if !ok {
		return ctx, errors.Wrap(errors.New("missing dependent value in context"), fmt.Sprintf("未找到 %s 表初始化数据", model.SysDictionary{}.TableName()))
	}
	trueTmp := true
	dictionaries[0].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "男", Value: "1", Status: &trueTmp, Sort: 1},
		{Label: "女", Value: "2", Status: &trueTmp, Sort: 2},
	}

	dictionaries[1].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "smallint", Value: "1", Status: &trueTmp, Extend: "mysql", Sort: 1},
		{Label: "mediumint", Value: "2", Status: &trueTmp, Extend: "mysql", Sort: 2},
		{Label: "int", Value: "3", Status: &trueTmp, Extend: "mysql", Sort: 3},
		{Label: "bigint", Value: "4", Status: &trueTmp, Extend: "mysql", Sort: 4},
		{Label: "int2", Value: "5", Status: &trueTmp, Extend: "pgsql", Sort: 5},
		{Label: "int4", Value: "6", Status: &trueTmp, Extend: "pgsql", Sort: 6},
		{Label: "int6", Value: "7", Status: &trueTmp, Extend: "pgsql", Sort: 7},
		{Label: "int8", Value: "8", Status: &trueTmp, Extend: "pgsql", Sort: 8},
	}

	dictionaries[2].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "date", Status: &trueTmp},
		{Label: "time", Value: "1", Status: &trueTmp, Extend: "mysql", Sort: 1},
		{Label: "year", Value: "2", Status: &trueTmp, Extend: "mysql", Sort: 2},
		{Label: "datetime", Value: "3", Status: &trueTmp, Extend: "mysql", Sort: 3},
		{Label: "timestamp", Value: "5", Status: &trueTmp, Extend: "mysql", Sort: 5},
		{Label: "timestamptz", Value: "6", Status: &trueTmp, Extend: "pgsql", Sort: 5},
	}
	dictionaries[3].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "float", Status: &trueTmp},
		{Label: "double", Value: "1", Status: &trueTmp, Extend: "mysql", Sort: 1},
		{Label: "decimal", Value: "2", Status: &trueTmp, Extend: "mysql", Sort: 2},
		{Label: "numeric", Value: "3", Status: &trueTmp, Extend: "pgsql", Sort: 3},
		{Label: "smallserial", Value: "4", Status: &trueTmp, Extend: "pgsql", Sort: 4},
	}

	dictionaries[4].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "char", Status: &trueTmp},
		{Label: "varchar", Value: "1", Status: &trueTmp, Extend: "mysql", Sort: 1},
		{Label: "tinyblob", Value: "2", Status: &trueTmp, Extend: "mysql", Sort: 2},
		{Label: "tinytext", Value: "3", Status: &trueTmp, Extend: "mysql", Sort: 3},
		{Label: "text", Value: "4", Status: &trueTmp, Extend: "mysql", Sort: 4},
		{Label: "blob", Value: "5", Status: &trueTmp, Extend: "mysql", Sort: 5},
		{Label: "mediumblob", Value: "6", Status: &trueTmp, Extend: "mysql", Sort: 6},
		{Label: "mediumtext", Value: "7", Status: &trueTmp, Extend: "mysql", Sort: 7},
		{Label: "longblob", Value: "8", Status: &trueTmp, Extend: "mysql", Sort: 8},
		{Label: "longtext", Value: "9", Status: &trueTmp, Extend: "mysql", Sort: 9},
	}

	dictionaries[5].SysDictionaryDetails = []model.SysDictionaryDetail{
		{Label: "tinyint", Value: "1", Extend: "mysql", Status: &trueTmp},
		{Label: "bool", Value: "2", Extend: "pgsql", Status: &trueTmp},
	}
	for _, dict := range dictionaries {
		if err := db.Model(&dict).Association("SysDictionaryDetails").Replace(dict.SysDictionaryDetails); err != nil {
			return ctx, errors.Wrap(err, model.SysDictionaryDetail{}.TableName()+"表初始化失败！")
		}
	}
	return ctx, nil
}

func (d *initDictionaryDetail) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysDictionaryDetail{})
}

func (d *initDictionaryDetail) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var dict model.SysDictionary
	if err := db.Preload("SysDictionaryDetails").First(&dict, &model.SysDictionary{Name: "数据库bool类型"}).Error; err != nil {
		return false
	}
	return len(dict.SysDictionaryDetails) > 0 && dict.SysDictionaryDetails[0].Label == "tinyint"
}
