package system

import "github/shansec/go-vue-admin/global"

type SysBaseMenu struct {
	global.MAY_MODEL
	MenuLevel uint          `json:"-`
	ParentId  uint          `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string        `json:"path" gorm:"comment:路由path"`
	Name      string        `json:"name" gorm:"comment:路由name"`
	Hidden    bool          `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string        `json:"component" gorm:"comment:对应的文件路径"`
	Sort      int           `json:"sort" gorm:"comment:排序"`
	Meta      Meta          `json:"meta" gorm:"embedded;comment:附加属性"`
	SysRoles  []SysRole     `json:"sysRoles" gorm:"many2many:sys_role_menus;"`
	Children  []SysBaseMenu `json:"children" gorm:"-"`
}

type Meta struct {
	KeepAlive bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	Title     string `json:"title" gorm:"comment:菜单名"`
	Icon      string `json:"icon" gorm:"comment:菜单图标"`
	Affix     bool   `json:"affix" gorm:"comment:自动关闭tab"`
}

func (SysBaseMenu) TableName() string {
	return "sys_menus"
}
