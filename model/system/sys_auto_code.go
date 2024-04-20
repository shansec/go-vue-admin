package system

import "github/shansec/go-vue-admin/global"

type SysAutoCode struct {
	global.MAY_MODEL
	PackageName string `json:"package_name" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:标签"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}
