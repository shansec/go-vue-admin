package common

import "github.com/shansec/go-vue-admin/utils"

var (
	PageInfoVerify = utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}}
)
