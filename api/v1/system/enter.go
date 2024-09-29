package system

import (
	"github.com/shansec/go-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
	DeptApi
	SystemConfigApi
	AutoCodeApi
	DBApi
	SysApi
	RoleApi
	MenuApi
	DictionaryApi
	DictionaryDetailApi
}

var (
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	deptService             = service.ServiceGroupApp.SystemServiceGroup.DeptService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	initDbService           = service.ServiceGroupApp.SystemServiceGroup.InitService
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	roleService             = service.ServiceGroupApp.SystemServiceGroup.RoleService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
)
