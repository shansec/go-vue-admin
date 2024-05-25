package common

import "github/shansec/go-vue-admin/utils"

var (
	PageInfoVerify = utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}}
)
