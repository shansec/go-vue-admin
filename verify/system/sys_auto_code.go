package system

import "github/shansec/go-vue-admin/utils"

var (
	AutoPackageVerify = utils.Rules{"PackageName": {utils.NotEmpty()}}
)
