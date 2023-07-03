package response

import (
	"github/shansec/go-vue-admin/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type Login struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
