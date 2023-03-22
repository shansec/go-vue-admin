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

// @Tags SysUser
// @Summary 用户注册账号
// @Produce json
// @Param data body systemReq.Register {用户名、密码、昵称、手机号}
// @Success 200
// @Router /user/register POST

func (b *BaseApi) Register(c *gin.Context) {
	var register systemReq.Register
	_ = c.ShouldBindJSON(&register)
	if err := utils.Verify(register, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{
		Username:  register.Username,
		NickName:  register.NickName,
		Password:  register.Password,
		HeaderImg: register.HeaderImg,
		Phone:     register.Phone,
	}
	userResgisterRes, err := userService.Register(*user)
	if err != nil {
		global.MAY_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册成功", c)
	}
}

func (b *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	_ = c.ShouldBindJSON(&login)
	if err := utils.Verify(login, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &system.SysUser{Username: login.Username, Password: login.Password}
	if user, err := userService.Login(u); err != nil {
		global.MAY_LOGGER.Error("登陆失败，用户名不存在或者密码错误", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.TokenNext(c, *user)
	}
}

// 登录成功签发 jwt

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	jwt := &utils.JWT{SigningKey: []byte(global.MAY_CONFIG.JWT.SigningKey)}
	claims := jwt.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.MAY_LOGGER.Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(systemRes.Login{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
