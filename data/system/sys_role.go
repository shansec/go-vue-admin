package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	model "github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/service/system"
	"github.com/shansec/go-vue-admin/utils"
)

const initOrderRole = initOrderMenu + 1

type initRole struct{}

func init() {
	system.RegisterInit(initOrderRole, &initRole{})
}

func (r initRole) InitTableName() string {
	return model.SysRole{}.TableName()
}

func (r *initRole) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysRole{})

}

func (r *initRole) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	role := []model.SysRole{
		{RoleId: 888, RoleName: "普通用户", ParentId: utils.Pointer[uint](0)},
		{RoleId: 9528, RoleName: "测试角色", ParentId: utils.Pointer[uint](0)},
		{RoleId: 8881, RoleName: "普通用户子角色", ParentId: utils.Pointer[uint](888)},
	}
	if err := db.Create(&role).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysRole{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, r.InitTableName(), role)
	return next, nil
}

func (r *initRole) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysRole{})
}

func (r *initRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var role model.SysRole
	if errors.Is(db.Where("role_id = ?", 1).First(&role).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
