package system

import "github/shansec/go-vue-admin/utils"

var (
	CreateDeptVerify = utils.Rules{"DeptName": {utils.NotEmpty()}, "Leader": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}, "Email": {utils.NotEmpty()}}
	DeleteDeptVerify = utils.Rules{"DeptId": {utils.NotEmpty()}}
	UpdateDeptVerify = utils.Rules{"DeptName": {utils.NotEmpty()}, "Leader": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}, "Email": {utils.NotEmpty()}}
)
