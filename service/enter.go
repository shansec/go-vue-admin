package service

import "github/May-cloud/go-vue-admin/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupAlias = new(ServiceGroup)
