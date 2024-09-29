package service

import (
	"github.com/shansec/go-vue-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
