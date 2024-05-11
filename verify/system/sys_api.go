package system

import "github/shansec/go-vue-admin/utils"

var (
	CreateApiVerify = utils.Rules{"Path": {utils.NotEmpty()}, "Description": {utils.NotEmpty()}, "ApiGroup": {utils.NotEmpty()}, "Method": {utils.NotEmpty()}}
	DeleteApiVerify = utils.Rules{"ID": {utils.NotEmpty()}}
)
