package router

import (
	"github/shansec/go-vue-admin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupAlias = new(RouterGroup)
