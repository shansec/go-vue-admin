package request

// GetPackageList structure
type GetPackageList struct {
	Page        int    `json:"page"`         // 页码
	PagSize     int    `json:"pageSize"`     // 每页大小
	PackageName string `json:"package_name"` // 包名
}
