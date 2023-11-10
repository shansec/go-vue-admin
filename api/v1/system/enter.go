package system

import (
	"github/shansec/go-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
	DeptApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	deptService = service.ServiceGroupApp.SystemServiceGroup.DeptService
)
