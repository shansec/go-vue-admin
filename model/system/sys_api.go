package system

import "github/shansec/go-vue-admin/global"

type SysApi struct {
	global.MAY_MODEL
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api所属组"`
	Method      string `json:"method" gorm:"comment:方法"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}
