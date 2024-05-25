package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github/shansec/go-vue-admin/global"
	model "github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/service/system"
)

const initOrderApi = InitOrder + 4

type initApi struct{}

func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (a *initApi) InitTableName() string {
	return model.SysApi{}.TableName()
}

func (a *initApi) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysApi{})

}

func (a *initApi) InitData(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	apis := []model.SysApi{
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/modifyPassword", Method: "POST", Description: "用户修改密码", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/register", Method: "POST", Description: "添加用户", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/delUserInfo", Method: "DELETE", Description: "删除用户", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/getUserInfo", Method: "GET", Description: "获取指定用户信息", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/getUsersInfo", Method: "GET", Description: "获取用户列表", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/updateUserInfo", Method: "PUT", Description: "修改用户信息", ApiGroup: "用户"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/user/updateUserStatus", Method: "PUT", Description: "修改用户状态", ApiGroup: "用户"},

		{Path: global.MAY_CONFIG.System.RouterPrefix + "/dept/createDept", Method: "POST", Description: "添加部门", ApiGroup: "部门"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/dept/delDeptInfo", Method: "DELETE", Description: "删除部门", ApiGroup: "部门"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/dept/getDeptList", Method: "GET", Description: "获取部门列表", ApiGroup: "部门"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/dept/updateDeptInfo", Method: "PUT", Description: "修改部门信息", ApiGroup: "部门"},

		{Path: global.MAY_CONFIG.System.RouterPrefix + "/autocode/createPackage", Method: "POST", Description: "自动化创建代码包", ApiGroup: "代码生成"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/autocode/delPackageInfo", Method: "DELETE", Description: "删除创建代码包", ApiGroup: "代码生成"},
		{Path: global.MAY_CONFIG.System.RouterPrefix + "/autocode/getPackageList", Method: "POST", Description: "获取创建的包列表", ApiGroup: "代码生成"},

		{Path: global.MAY_CONFIG.System.RouterPrefix + "/system/status", Method: "GET", Description: "获取服务器信息", ApiGroup: "服务器"},
	}
	if err := db.Create(&apis).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysApi{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, a.InitTableName(), apis)
	return next, nil
}

func (a *initApi) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysApi{})
}

func (a *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var api model.SysApi
	if errors.Is(db.Where("path = ? AND method = ?", global.MAY_CONFIG.System.RouterPrefix+"/user/modifyPassword", "POST").First(&api).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
