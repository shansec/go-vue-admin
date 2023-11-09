package system

type SysDept struct {
	DeptId   int    `json:"deptId" gorm:"primary_key;not null;unique;comment:部门ID;size:90"`
	ParentId int    `json:"parentId" gorm:""`
	DeptPath string `json:"deptPath" gorm:"size:255"`
	DeptName string `json:"deptName" gorm:"size:255"`
}
