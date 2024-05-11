package system

import (
	"github/shansec/go-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
	DeptApi
	SystemConfigApi
	AutoCodeApi
	DBApi
	SysApi
}

var (
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
	deptService         = service.ServiceGroupApp.SystemServiceGroup.DeptService
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	autoCodeService     = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	initDbService       = service.ServiceGroupApp.SystemServiceGroup.InitService
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
)
