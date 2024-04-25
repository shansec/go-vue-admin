package system

import (
	"github.com/satori/uuid"
	"github/shansec/go-vue-admin/global"
)

type SysUser struct {
	global.MAY_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username  string    `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Sex       int       `json:"sex" gorm:"default:1;comment:用户性别"`                                                    // 用户性别 1为男，2为女
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:admin;comment:用户昵称"`                                           // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone     string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email     string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Status    int       `json:"status" gorm:"default:1;comment:用户状态 1为开启，2为禁用"`                                       // 用户状态 1为开启，2为禁用
	RolesId   int       `json:"rolesId" gorm:"comment:用户角色ID"`                                                        // 用户角色ID
	SysRole   SysRole   `json:"sysRole" gorm:"foreignKey:RolesId;references:RoleId"`
	DeptsId   int       `json:"deptsId" gorm:"用户部门ID"`
	SysDept   SysDept   `json:"sysDept" gorm:"foreignKey:DeptsId;references:DeptId"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
