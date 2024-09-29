package system

import "github.com/shansec/go-vue-admin/global"

type SysDictionary struct {
	global.MAY_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名(中)"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名(英)"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	SysDictionaryDetails []SysDictionaryDetail `json:"sys-dictionary-details" form:"sysDictionaryDetails"`
}

func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}
