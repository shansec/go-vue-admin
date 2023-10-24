package system

import "time"

type SysRole struct {
	RoleId    int        `json:"roleId" gorm:"primary_key;not null;unique;comment:角色ID;size:90"` // 角色编码
	RoleName  string     `json:"roleName" gorm:"size:128;comment:角色名称;"`                       // 角色名称
	Status    string     `json:"status" gorm:"size:4;comment:状态;"`                               // 状态 1禁用 2正常
	RoleKey   string     `json:"roleKey" gorm:"size:128;comment:角色代码;"`                        // 角色代码
	RoleSort  int        `json:"roleSort" gorm:"size:16;comment:角色排序;"`                        // 角色排序
	Flag      string     `json:"flag" gorm:"size:128;comment:角色标记;"`                           // 角色标记
	Remark    string     `json:"remark" gorm:"size:255;comment:角色备注;"`                         // 备注
	Admin     bool       `json:"admin" gorm:"size:4;"`
	DataScope string     `json:"dataScope" gorm:"size:128;"`
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time // 删除时间
}
