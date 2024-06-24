package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	model "github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/service/system"
)

const initOrderMenu = initOrderDept + 1

type initMenu struct{}

func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitTableName() string {
	return model.SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysBaseMenu{})

}

func (i *initMenu) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	menus := []model.SysBaseMenu{
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "dashboard", Name: "Dashboard", Component: "views/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "首页", Icon: "app-group-fill", Affix: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "supervisor", Name: "Supervisor", Component: "", Sort: 2, Meta: model.Meta{Title: "管理员", Icon: "admin", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 2, Path: "user", Name: "User", Component: "views/authorize/sys-user/index.vue", Sort: 6, Meta: model.Meta{Title: "用户管理", Icon: "user", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 2, Path: "menus", Name: "Menus", Component: "views/authorize/sys-menu/index.vue", Sort: 7, Meta: model.Meta{Title: "菜单管理", Icon: "menu", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 2, Path: "api", Name: "Api", Component: "views/authorize/sys-api/index.vue", Sort: 8, Meta: model.Meta{Title: "api 管理", Icon: "api", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 2, Path: "dept", Name: "Dept", Component: "views/authorize/sys-dept/index.vue", Sort: 9, Meta: model.Meta{Title: "部门管理", Icon: "tree", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "package", Name: "Package", Component: "views/authorize/sys-autocode/index.vue", Sort: 3, Meta: model.Meta{Title: "自动化包", Icon: "settings", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "profile", Name: "Profile", Component: "views/profile/index.vue", Sort: 4, Meta: model.Meta{Title: "个人设置", Icon: "user-setting", Affix: false}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "system", Name: "System", Component: "views/system/index.vue", Sort: 5, Meta: model.Meta{Title: "服务器信息", Icon: "service-side", Affix: false}},
	}
	if err := db.Create(&menus).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysBaseMenu{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, i.InitTableName(), menus)
	return next, nil
}

func (i *initMenu) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysBaseMenu{})
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var menu model.SysBaseMenu
	if errors.Is(db.Where("id = ?", 1).First(&menu).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
