package request

import (
	"github.com/satori/uuid"

	"github/shansec/go-vue-admin/model/system"
)

// Register structure
type Register struct {
	Username  string         `json:"username"`  // 用户登录名
	Sex       int            `json:"sex"`       // 用户性别
	Password  string         `json:"password"`  // 用户登录密码
	NickName  string         `json:"nickName"`  // 用户昵称
	HeaderImg string         `json:"headerImg"` // 用户头像
	Phone     string         `json:"phone"`     // 用户手机号
	Email     string         `json:"email"`     // 用户邮箱
	Status    int            `json:"status"`    // 用户状态
	DeptsId   int            `json:"deptsId"`   // 部门ID
	SysDept   system.SysDept `json:"sysDept"`
	RolesId   int            `json:"rolesId"` // 用户角色ID
	SysRole   system.SysRole `json:"sysRole"`
}

// Login structure
type Login struct {
	Username     string `json:"username"`     // 用户名
	Password     string `json:"password"`     // 密码
	Phone        string `json:"phone"`        // 手机号
	IsPhoneLogin bool   `json:"isPhoneLogin"` // 是否通过手机号登录
	Captcha      string `json:"captcha"`      // 验证码
	CaptchaId    string `json:"captchaId"`    // 验证码ID
}

// ChangePassword structure
type ChangePassword struct {
	ID          uint   `json:"uid"`         // user.id
	Password    string `json:"password"`    // 旧密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UUID struct {
	Uuid uuid.UUID `json:"uuid"`
}

// GetUserList structure
type GetUserList struct {
	Page     int    `json:"page"`     // 页码
	PagSize  int    `json:"pageSize"` // 每页大小
	NickName string `json:"nickName"` // 用户昵称
	Phone    string `json:"phone"`    // 用户手机号
	Status   string `json:"status"`   // 用户状态
}
