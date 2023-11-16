package system

import "github/shansec/go-vue-admin/utils"

var (
	CreateVerify = utils.Rules{"DeptName": {utils.NotEmpty()}, "Leader": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}, "Email": {utils.NotEmpty()}}
	DeleteVerify = utils.Rules{"DeptId": {utils.NotEmpty()}}
	UpdateVerify = utils.Rules{"DeptName": {utils.NotEmpty()}, "Leader": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}, "Email": {utils.NotEmpty()}}
)
