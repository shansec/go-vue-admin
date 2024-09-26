package system

import "github.com/shansec/go-vue-admin/utils"

var (
	CreateMenuVerify  = utils.Rules{"Path": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}, "Component": {utils.NotEmpty()}, "Meta": {utils.NotEmpty()}}
	DeleteMenuVerify  = utils.Rules{"ID": {utils.NotEmpty()}}
	UpdateMenuVerify  = utils.Rules{"ID": {utils.NotEmpty()}, "Path": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}, "Component": {utils.NotEmpty()}, "Meta": {utils.NotEmpty()}}
	AddRoleMenuVerify = utils.Rules{"RoleId": {utils.NotEmpty()}}
)
