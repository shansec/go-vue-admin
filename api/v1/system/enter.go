package system

import (
	"github/shansec/go-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupAlias.SystemServiceGroup.UserService
)
