package system

import (
	"github.com/gin-gonic/gin"
	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/model/common/response"
	"github/May-cloud/go-vue-admin/model/system"
	systemReq "github/May-cloud/go-vue-admin/model/system/request"
	systemRes "github/May-cloud/go-vue-admin/model/system/response"
	"github/May-cloud/go-vue-admin/utils"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Register(c *gin.Context) {
	var register systemReq.Register
	_ = c.ShouldBindJSON(&register)
	if err := utils.Verify(register, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: register.Username, NickName: register.NickName, Password: register.Password, HeaderImg: register.HeaderImg}
	userResgisterRes, err := userService.Register(*user)
	if err != nil {
		global.MAY_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册成功", c)
	}
}
