package system

import (
	"github.com/gin-gonic/gin"
	"go-vue-admin/model/common/response"
	"go-vue-admin/model/system"
	systemReq "go-vue-admin/model/system/request"
	"go-vue-admin/utils"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	_ = c.ShouldBindJSON(&login)
	if err := utils.Verify(login, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.SysUser{Username: login.Username, Password: login.Password}

}
