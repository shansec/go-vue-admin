package system

import (
	"github/May-cloud/go-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupAlias.SystemServiceGroup.UserService
)
