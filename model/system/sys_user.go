package system

import (
	"github/May-cloud/go-vue-admin/global"

	"github.com/satori/uuid"
)

type SysUser struct {
	global.MAY_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username  string    `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称 	// 用户侧边主题
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像 	// 用户角色ID
	Phone     string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email     string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
}
