package system

import "github/shansec/go-vue-admin/utils"

var (
	AutoPackageVerify = utils.Rules{"PackageName": {utils.NotEmpty()}}
	AutoCodeVerify    = utils.Rules{"Abbreviation": {utils.NotEmpty()}, "StructName": {utils.NotEmpty()}, "PackageName": {utils.NotEmpty()}, "Fields": {utils.NotEmpty()}}
)
