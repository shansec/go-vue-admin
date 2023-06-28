package service

import (
	"github/shansec/go-vue-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupAlias = new(ServiceGroup)
