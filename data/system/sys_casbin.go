package system

import (
	"context"

	"github.com/shansec/go-vue-admin/service/system"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitTableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: "/user/register", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/user/delUserInfo", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/user/getUsersInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/modifyPassword", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/updateUserInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/updateUserStatus", V2: "PUT"},

		{Ptype: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/deleteApi", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/updateApi", V2: "PUT"},

		{Ptype: "p", V0: "888", V1: "/autocode/createPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autocode/delPackageInfo", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/autocode/getPackageList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/dept/createDept", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dept/delDeptInfo", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/dept/getDeptList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dept/updateDeptInfo", V2: "PUT"},

		{Ptype: "p", V0: "888", V1: "/role/createRole", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/role/deleteRole", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/role/updateRole", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/role/getRoleList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/role/setChildRole", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/role/addRoleMenu", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/menu/createMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/deleteMenu", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/updateMenu", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuTree", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getRoleMenu", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/menu/getSpecialRoleMenu", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/dictionary/createDictionary", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dictionary/deleteDictionary", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/dictionary/updateDictionary", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/dictionary/getDictionary", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dictionary/getDictionaryInfoList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/dictionaryDetail/createDictionaryDetail", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dictionaryDetail/deleteDictionaryDetail", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/dictionaryDetail/updateDictionaryDetail", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/dictionaryDetail/getDictionaryDetail", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dictionaryDetail/getDictionaryDetailList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/system/getServerInfo", V2: "GET"},

		{Ptype: "p", V0: "9528", V1: "/user/register", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/system/getServerInfo", V2: "GET"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitTableName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitTableName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
