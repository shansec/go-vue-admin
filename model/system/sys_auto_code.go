package system

import "github/shansec/go-vue-admin/global"

type SysAutoCode struct {
	global.MAY_MODEL
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:标签"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}

func (SysAutoCode) TableName() string {
	return "sys_auto_codes"
}
