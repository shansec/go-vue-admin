package system

import "time"

type SysRole struct {
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time `sql:"index"`
	RoleId    uint       `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色编码
	RoleName  string     `json:"roleName" gorm:"size:128;"`                                      // 角色名称
	Status    string     `json:"status" gorm:"size:4;"`                                          // 状态 1禁用 2正常
	RoleKey   string     `json:"roleKey" gorm:"size:128;"`                                       //角色代码
	RoleSort  int        `json:"roleSort" gorm:""`                                               //角色排序
	Flag      string     `json:"flag" gorm:"size:128;"`                                          //
	Remark    string     `json:"remark" gorm:"size:255;"`                                        //备注
	Admin     bool       `json:"admin" gorm:"size:4;"`
	DataScope string     `json:"dataScope" gorm:"size:128;"`
}
