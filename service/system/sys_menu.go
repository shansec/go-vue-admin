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

// DeleteMenuService
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteMenuService
// @description: 删除菜单
// @param: menu system.SysBaseMenu
// @return: error
func (menuService *MenuService) DeleteMenuService(menu *system.SysBaseMenu) error {
	err := global.MAY_DB.Preload("SysRoles").Where("id = ?", menu.ID).First(&menu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("菜单不存在！")
	}
	err = global.MAY_DB.Transaction(func(ctx *gorm.DB) error {
		var err error

		if err = ctx.Preload("SysRoles").Where("id = ?", menu.ID).First(menu).Error; err != nil {
			return err
		}

		if len(menu.SysRoles) != 0 {
			if err = ctx.Model(menu).Association("SysRoles").Delete(menu.SysRoles); err != nil {
				return err
			}
		}

		if err = ctx.Where("id = ?", menu.ID).Unscoped().Delete(menu).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}
