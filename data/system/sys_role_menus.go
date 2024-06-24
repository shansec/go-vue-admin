package system

import (
	"context"
	systemModel "github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/service/system"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

const initOrderMenuRole = initOrderMenu + initOrderRole

type initMenuRole struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuRole, &initMenuRole{})
}

func (i *initMenuRole) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuRole) TableInited(ctx context.Context) bool {
	return false // always replace
}

func (i initMenuRole) InitTableName() string {
	return "sys_menu_authorities"
}

func (i *initMenuRole) InitData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	authorities, ok := ctx.Value(initRole{}.InitTableName()).([]systemModel.SysRole)
	if !ok {
		return ctx, errors.Wrap(errors.New("missing dependent value in context"), "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}
	menus, ok := ctx.Value(initMenu{}.InitTableName()).([]systemModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx
	// 888
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(menus); err != nil {
		return next, err
	}

	// 8881
	menu8881 := menus[:2]
	menu8881 = append(menu8881, menus[7])
	if err = db.Model(&authorities[1]).Association("SysBaseMenus").Replace(menu8881); err != nil {
		return next, err
	}
	return next, nil
}

func (i *initMenuRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &systemModel.SysRole{}
	if ret := db.Model(auth).
		Where("role_id = ?", 9528).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}
