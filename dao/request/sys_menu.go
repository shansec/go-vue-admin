package request

import (
	"github/shansec/go-vue-admin/model/system"

	"github/shansec/go-vue-admin/global"
)

type MenuRoleInfo struct {
	RoleId uint                 `json:"roleId"`
	Menus  []system.SysBaseMenu `json:"menus"`
}

type GetMenuByName struct {
	Name string `json:"name"`
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{
		{
			MAY_MODEL: global.MAY_MODEL{ID: 1},
			ParentId:  0,
			Path:      "dashboard",
			Name:      "Dashboard",
			Hidden:    true,
			Component: "views/dashboard/index.vue",
			Sort:      1,
			Meta: system.Meta{
				KeepAlive: false,
				Title:     "首页",
				Icon:      "app-group-fill",
				Affix:     true,
			},
		},
	}
}
