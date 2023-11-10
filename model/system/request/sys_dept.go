package request

// Create structure
type Create struct {
	ParentId int    `json:"parentId"`
	DeptPath string `json:"deptPath"`
	DeptName string `json:"deptName"`
	Sort     int    `json:"sort"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

// GetDeptList structure
type GetDeptList struct {
	Page     int    `json:"page"`     // 页码
	PagSize  int    `json:"pageSize"` // 每页大小
	DeptName string `json:"deptName"` // 用户昵称 	// 用户手机号
	Status   string `json:"status"`   // 用户状态
}
