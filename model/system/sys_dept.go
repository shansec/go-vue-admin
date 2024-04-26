package system

import "time"

type SysDept struct {
	DeptId    int       `json:"deptId" gorm:"primary_key;not null;unique;comment:部门ID;size:90"`
	ParentId  int       `json:"parentId" gorm:"comment:上级部门"`
	DeptPath  string    `json:"deptPath" gorm:"size:255"`
	DeptName  string    `json:"deptName" gorm:"comment:部门名称;size:255"`
	Sort      int       `json:"sort" gorm:"comment:排序;size:4"`
	Leader    string    `json:"leader" gorm:"comment:负责人;size:128"`
	Phone     string    `json:"phone" gorm:"comment:手机号码;size:11"`
	Email     string    `json:"email" gorm:"comment:邮箱;size:64"`
	Status    string    `json:"status" gorm:"default:1;comment:状态;size:4"`
	Children  []SysDept `json:"children" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (SysDept) TableName() string {
	return "sys_depts"
}
