package system

import "github.com/shansec/go-vue-admin/global"

type SysDictionaryDetail struct {
	global.MAY_MODEL
	Label           string `json:"label" form:"label" gorm:"column:label;comment:展示值"`
	Value           string `json:"value" form:"value" gorm:"column:value;comment:字典值"`
	Extend          string `json:"extend" form:"extend" gorm:"column:extend;comment:扩展值"`
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id"`
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
