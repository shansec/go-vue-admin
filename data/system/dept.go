package system

import (
	"context"
	"github.com/pkg/errors"
	model "github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/service/system"
	"gorm.io/gorm"
)

const initOrderDept = InitOrder + 1

type initDept struct{}

func init() {
	system.RegisterInit(initOrderDept, &initDept{})
}

func (d *initDept) InitTableName() string {
	return model.SysDept{}.TableName()
}

func (d *initDept) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysDept{})

}

func (d *initDept) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	dept := model.SysDept{
		DeptId:   1,
		ParentId: 0,
		DeptPath: "/",
		DeptName: "顶级部门",
		Sort:     1,
		Status:   "1",
		Leader:   "admin",
		Phone:    "13412345678",
		Email:    "admin@163.com",
	}
	if err := db.Create(&dept).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysDept{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, d.InitTableName(), dept)
	return next, nil
}

func (d *initDept) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysDept{})
}

func (d *initDept) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var dept model.SysDept
	if errors.Is(db.Where("dept_id = ?", 1).First(&dept).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
