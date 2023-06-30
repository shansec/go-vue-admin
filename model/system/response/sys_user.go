package response

import (
	"github/shansec/go-vue-admin/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type Login struct {
	User      system.SysUser `json:"user"`
	AToken    string         `json:"AToken"`
	RToken    string         `json:"RToken"`
	ExpiresAt int64          `json:"expiresAt"`
}
