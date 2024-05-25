package response

import (
	"github.com/satori/uuid"

	"github/shansec/go-vue-admin/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type SysUsersResponse struct {
	Users []system.SysUser `json:"user"`
}

type Login struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type UserResponse struct {
	UUID      uuid.UUID      `json:"uuid"`
	Username  string         `json:"userName"` // 用户登录名
	Password  string         `json:"-"`
	NickName  string         `json:"nickName"`  // 用户昵称
	HeaderImg string         `json:"headerImg"` // 用户头像
	Phone     string         `json:"phone"`     // 用户手机号
	Email     string         `json:"email"`     // 用户邮箱
	RolesId   int            `json:"rolesId"`   // 用户角色ID
	SysRole   system.SysRole `json:"sysRole"`
}
