package response

import "github.com/shansec/go-vue-admin/model/system"

type SysMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}
