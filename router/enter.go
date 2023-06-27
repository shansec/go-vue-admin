package router

import (
	"github/May-cloud/go-vue-admin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupAlias = new(RouterGroup)
