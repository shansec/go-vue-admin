package system

type SysRoleMenu struct {
	MenuId uint `json:"menuId" gorm:"column:sys_base_menu_id"`
	RoleId uint `json:"roleId" gorm:"column:sys_role_role_id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menus"
}
