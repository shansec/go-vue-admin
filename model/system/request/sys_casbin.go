package request

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CasbinInReceive struct {
	RoleID      uint         `json:"roleId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/base/login", Method: "POST"},
	}
}
