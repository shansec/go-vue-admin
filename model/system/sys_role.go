package system

import "time"

type SysRole struct {
	CreatedAt     time.Time     // 创建时间
	UpdatedAt     time.Time     // 更新时间
	DeletedAt     *time.Time    `sql:"index"`                                                           // 删除时间
	RoleId        uint          `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleName      string        `json:"roleName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      *uint         `json:"parentId" gorm:"comment:父角色ID"`                                  // 父角色ID
	DataRoleId    []*SysRole    `json:"dataRoleId" gorm:"many2many:sys_data_role_id;"`
	Children      []SysRole     `json:"children" gorm:"-"`
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_role_menus;"`
	Users         []SysUser     `json:"-" gorm:"many2many:sys_user_role;"`
	DefaultRouter string        `json:"defaultRouter" gorm:"comment:默认菜单;default:Dashboard"` // 默认菜单(默认dashboard)
}

func (SysRole) TableName() string {
	return "sys_roles"
}
