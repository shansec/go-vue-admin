package v1

import (
	"github.com/shansec/go-vue-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
