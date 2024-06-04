package system

import (
	"errors"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"

	"gorm.io/gorm"
)

type MenuService struct{}

// CreateMenuService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateMenuService
// @description: 添加菜单
// @param: menu system.SysBaseMenu
// @return: system.SysBaseMenu, error
func (menuService *MenuService) CreateMenuService(menu system.SysBaseMenu) (system.SysBaseMenu, error) {
	var menuInfo system.SysBaseMenu
	var err error
	if err = global.MAY_DB.Where("path = ? AND name = ?", menu.Path, menu.Name).First(&menuInfo).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return system.SysBaseMenu{}, errors.New("存在相同菜单")
	}
	err = global.MAY_DB.Create(&menu).Error
	if err != nil {
		return system.SysBaseMenu{}, err
	}
	return menu, nil
}
