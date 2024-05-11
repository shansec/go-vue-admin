package system

import "github/shansec/go-vue-admin/utils"

var (
	IdVerify             = utils.Rules{"ID": []string{utils.NotEmpty()}}
	LoginVerify          = utils.Rules{"CaptchaId": {utils.NotEmpty()}, "Username": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}}
	RegisterVerify       = utils.Rules{"Username": {utils.NotEmpty()}, "NickName": {utils.NotEmpty()}, "Password": {utils.NotEmpty()}}
	UpdateUserVerify     = utils.Rules{"NickName": {utils.NotEmpty()}, "Phone": {utils.NotEmpty()}, "Email": {utils.NotEmpty()}}
	PageInfoVerify       = utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}}
	ChangePasswordVerify = utils.Rules{"Password": {utils.NotEmpty()}, "NewPassword": {utils.NotEmpty()}}
)
