package response

import "github/shansec/go-vue-admin/model/system"

type SysMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}
