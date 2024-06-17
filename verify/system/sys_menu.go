package system

import "github/shansec/go-vue-admin/utils"

var (
	CreateMenuVerify = utils.Rules{"ParentId": {utils.NotEmpty()}, "Path": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}, "Component": {utils.NotEmpty()}, "Meta": {utils.NotEmpty()}}
	DeleteMenuVerify = utils.Rules{"ID": {utils.NotEmpty()}}
)
