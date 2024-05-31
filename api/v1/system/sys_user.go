package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	systemRes "github/shansec/go-vue-admin/model/system/response"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
)

type BaseApi struct{}

// Login
// @Summary 用户登录
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.Login true "用户登录"
// @Success 200 {object} response.Response{data=systemRes.Login, msg=string}	"用户登录"
// @Router /base/login [POST]
func (b *BaseApi) Login(c *gin.Context) {
	if global.MAY_DB == nil {
		global.MAY_LOGGER.Error("数据库未初始化，请先初始化")
		response.FailWithMessage("数据库未初始化，请先初始化", c)
		return
	}

	var login systemReq.Login
	_ = c.ShouldBindJSON(&login)
	if err := utils.Verify(login, SystemVerify.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(login.CaptchaId, login.Captcha, true) {
		u := &system.SysUser{Username: login.Username, Password: login.Password, Phone: login.Phone}
		if user, err := userService.Login(u, login.IsPhoneLogin); err != nil {
			var errStr string = "用户名不存在或者密码错误"
			if login.IsPhoneLogin {
				errStr = "手机号不存在或者密码错误"
			}
			global.MAY_LOGGER.Error(errStr, zap.Error(err))
			response.FailWithMessage(errStr, c)
		} else {
			b.TokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// TokenNext 登录成功签发 jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	jwt := &utils.JWT{SigningKey: []byte(global.MAY_CONFIG.JWT.SigningKey)}
	claims := jwt.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
		RoleId:   user.RolesId,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.MAY_LOGGER.Error("获取 token 失败", zap.Error(err))
		response.FailWithMessage("获取 token 失败", c)
		return
	}
	response.OkWithDetailed(systemRes.Login{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}

// ModifyPassword
// @Summary 修改密码
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.ChangePassword true "修改密码"
// @Success 200 {object} response.Response{msg=string}	"修改密码,返回修改结果"
// @Router /user/modifyPassword [POST]
func (b *BaseApi) ModifyPassword(c *gin.Context) {
	var modifyPassword systemReq.ChangePassword
	_ = c.ShouldBindJSON(&modifyPassword)

	if err := utils.Verify(modifyPassword, SystemVerify.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUseID(c)
	changePassword := &system.SysUser{MAY_MODEL: global.MAY_MODEL{ID: uid}, Password: modifyPassword.Password}
	if _, err := userService.ChangePassword(changePassword, modifyPassword.NewPassword); err != nil {
		global.MAY_LOGGER.Error("修改失败", zap.Error(err))
		response.FailWithMessage("修改失败，原密码不正确", c)
	} else {
		response.OkWithMessage("修改成功，请重新登录", c)
	}
}

// Register
// @Summary 用户注册账号
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.Register true "用户注册"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse, msg=string}	"用户注册"
// @Router /user/register [POST]
func (b *BaseApi) Register(c *gin.Context) {
	var register systemReq.Register
	_ = c.ShouldBindJSON(&register)
	if err := utils.Verify(register, SystemVerify.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{
		Username:  register.Username,
		Sex:       register.Sex,
		NickName:  register.NickName,
		Password:  register.Password,
		HeaderImg: register.HeaderImg,
		Phone:     register.Phone,
		Email:     register.Email,
		Status:    register.Status,
		DeptsId:   register.DeptsId,
		RolesId:   888,
	}
	userResgisterRes, err := userService.Register(*user)
	if err != nil {
		global.MAY_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userResgisterRes}, "注册成功", c)
	}
}

// DelUserInfo
// @Summary 删除用户信息
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.UUID true "删除用户信息"
// @Success 200 {object} response.Response{msg=string} "删除用户信息，返回操作结果"
// @Router /user/delUserInfo [Delete]
func (b *BaseApi) DelUserInfo(c *gin.Context) {
	var uuid systemReq.UUID
	_ = c.ShouldBindJSON(&uuid)
	if err := userService.DelUserInformation(uuid.Uuid); err != nil {
		global.MAY_LOGGER.Error("删除用户失败", zap.Error(err))
		response.FailWithMessage(" ", c)
	} else {
		response.OkWithMessage("删除用户成功", c)
	}
}

// UpdateUserInfo
// @Summary 更新用户信息
// @Tags SysUser
// @Produce json
// @Param   data body system.SysUser true "更新用户信息"
// @Success 200 {object} response.Response{msg=string} "更新用户信息，返回操作结果"
// @Router /user/updateUserInfo [POST]
func (b *BaseApi) UpdateUserInfo(c *gin.Context) {
	var user system.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, SystemVerify.UpdateUserVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.UpdateUserInformation(&user); err != nil {
		global.MAY_LOGGER.Error("更改用户信息失败", zap.Error(err))
		response.FailWithMessage("更改用户信息失败", c)
	} else {
		response.OkWithMessage("更改用户信息成功", c)
	}
}

// GetUserInfo
// @Summary 获取用户信息
// @Tags SysUser
// @Produce json
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse, msg=string} "获取用户信息"
// @Router /user/getUserInfo [GET]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUseUuid(c)
	if user, err := userService.GetUserInformation(uuid); err != nil {
		global.MAY_LOGGER.Error("获取用户信息失败", zap.Error(err))
		response.FailWithMessage("获取用户信息失败，用户信息不存在", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{
			User: *user,
		}, "获取用户信息成功", c)
	}
}

// GetUsersInfo
// @Summary 获取用户列表
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.GetUserList true "获取用户列表"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string} "获取用户列表"
// @Router /user/getUsersInfo [POST]
func (b *BaseApi) GetUsersInfo(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if users, total, err := userService.GetUsersInformation(pageInfo); err != nil {
		global.MAY_LOGGER.Error("获取用户列表失败", zap.Error(err))
		response.FailWithMessage("获取用户列表失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     users,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取用户列表成功", c)
	}
}

// UpdateUserStatus
// @Summary 更改用户状态
// @Tags SysUser
// @Produce json
// @Param   data body systemReq.UUID true "更改用户状态"
// @Success 200 {object} response.Response{msg=string} "更改用户状态，返回操作结果"
// @Router /user/updateUserStatus [PUT]
func (b *BaseApi) UpdateUserStatus(c *gin.Context) {
	var uuid systemReq.UUID
	_ = c.ShouldBindJSON(&uuid)
	if err := userService.UpdateStatus(uuid.Uuid); err != nil {
		global.MAY_LOGGER.Error("更改用户状态失败", zap.Error(err))
		response.FailWithMessage("更改用户状态失败", c)
	} else {
		response.OkWithMessage("更改成功", c)
	}
}
