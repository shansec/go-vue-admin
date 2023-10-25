package request

import "github/shansec/go-vue-admin/model/system"

// Register structure
type Register struct {
	Username  string         `json:"username"`  // 用户登录名
	Password  string         `json:"password"`  // 用户登录密码
	NickName  string         `json:"nickName"`  // 用户昵称
	HeaderImg string         `json:"headerImg"` // 用户头像
	Phone     string         `json:"phone"`     // 用户手机号
	Email     string         `json:"email"`
	RolesId   int            `json:"rolesId"` // 用户角色ID
	SysRole   system.SysRole `json:"sysRole"`
}

// Login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// ChangePassword structure
type ChangePassword struct {
	ID          uint   `json:"uid"`         // user.id
	Password    string `json:"password"`    // 旧密码
	NewPassword string `json:"newPassword"` // 新密码
}
