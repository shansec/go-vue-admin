package system

import (
	uuid "github.com/satori/go.uuid"
	"go-vue-admin/global"
)

type SysUser struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment: 用户UUID"`
	Username  string    `json:"userName" gorm:"comment: 用户登录名"`
	Password  string    `json:"-" gorm:"comment: 用户登录密码"`
	NickName  string    `json:"nickName" gorm:"default: 系统用户; comment: 用户昵称"`
	HeaderImg string    `json:"headerImg" gorm:"default: 1111;comment: 用户头像"`
	AuthorityId
	Authorities
	AuthorityId string `json:"authorityId" gorm:"default: 888; comment: 用户角色ID"`
	Phone       string `json:"phone" gorm:"comment: 用户手机号"`
	Email       string `json:"email" gorm:"comment: 用户邮箱"`
}
