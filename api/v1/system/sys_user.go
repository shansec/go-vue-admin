package system

import (
	"github.com/gin-gonic/gin"
	"go-vue-admin/global"
	"go-vue-admin/model/common/response"
	"go-vue-admin/model/system"
	systemReq "go-vue-admin/model/system/request"
	"go-vue-admin/utils"
	"go.uber.org/zap"
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
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOGGER.Error("登陆失败，用户名或者密码错误！", zap.Error(err))
		return
	}
	b.TokenNext(c, *user)
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {

}
