package system

import (
	"errors"
	"github/shansec/go-vue-admin/dao/common/request"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"time"

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

// GetMenuListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetMenuListService
// @description: 分页获取菜单信息
// @param: pageInfo request.PageInfo
// @return: list interface{}, total int64, err error
func (menuService *MenuService) GetMenuListService(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.MAY_DB.Model(&system.SysBaseMenu{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var menuList []system.SysBaseMenu
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", "0").Find(&menuList).Error
	if err != nil {
		return nil, 0, err
	}
	for index := range menuList {
		err = menuService.findChildrenMenu(&menuList[index])
	}
	return menuList, total, nil
}

// findChildrenMenu
// @author: [Shansec](https://github.com/shansec)
// @function: findChildrenMenu
// @description: 分页获取菜单信息辅助方法，查找子菜单
// @param: menu *system.SysBaseMenu
// @return: err error
func (menuService *MenuService) findChildrenMenu(menu *system.SysBaseMenu) (err error) {
	err = global.MAY_DB.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error
	if len(menu.Children) > 0 {
		for index := range menu.Children {
			err = menuService.findChildrenMenu(&menu.Children[index])
		}
	}
	return err
}

// UpdateMenuService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateMenuService
// @description: 修改菜单信息
// @param: menu system.SysBaseMenu
// @return: error
func (menuService *MenuService) UpdateMenuService(menu system.SysBaseMenu) error {
	var oldMenu system.SysBaseMenu
	err := global.MAY_DB.Where("id = ?", menu.ID).First(&oldMenu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("菜单不存在")
	}
	err = global.MAY_DB.Model(&oldMenu).Updates(map[string]interface{}{
		"parent_id":  menu.ParentId,
		"name":       menu.Name,
		"path":       menu.Path,
		"hidden":     menu.Hidden,
		"component":  menu.Component,
		"sort":       menu.Sort,
		"keep_alive": menu.Meta.KeepAlive,
		"title":      menu.Meta.Title,
		"icon":       menu.Meta.Icon,
		"affix":      menu.Meta.Affix,
		"updated_at": time.Now(),
	}).Error
	return err
}

// GetMenuTreeService
// @author: [Shansec](https://github.com/shansec)
// @function: GetMenuTreeService
// @description: 获取树状菜单
// @param: nil
// @return: list interface{}, err error
func (menuService *MenuService) GetMenuTreeService() (list interface{}, err error) {
	var menuList []system.SysBaseMenu
	db := global.MAY_DB.Model(&system.SysBaseMenu{})
	err = db.Where("parent_id = ?", "0").Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	for index := range menuList {
		err = menuService.findChildrenMenu(&menuList[index])
	}
	return menuList, nil
}

// GetRoleMenuService
// @author: [Shansec](https://github.com/shansec)
// @function: GetRoleMenuService
// @description: 获取角色菜单
// @param: roleId uint
// @return: menus []system.SysBaseMenu, err error
func (menuService *MenuService) GetRoleMenuService(roleId uint) (menuList []system.SysBaseMenu, err error) {
	var roleMenus []system.SysRoleMenu
	var menus []system.SysBaseMenu
	var menuIds []uint

	err = global.MAY_DB.Where("sys_role_role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, err
	}

	for _, menu := range roleMenus {
		menuIds = append(menuIds, menu.MenuId)
	}

	err = global.MAY_DB.Where("id IN (?)", menuIds).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for _, menu := range menus {
		if menu.ParentId != 0 {
			continue
		}
		menuResult := menuService.GetMenuCall(menus, menu)
		menuList = append(menuList, menuResult)
	}
	return menuList, nil
}

func (menuService *MenuService) GetMenuCall(menuList []system.SysBaseMenu, menu system.SysBaseMenu) system.SysBaseMenu {
	var menuCalls []system.SysBaseMenu
	lists := menuList
	for _, list := range lists {
		if menu.ID != list.ParentId {
			continue
		}
		menuCall := system.SysBaseMenu{}
		menuCall.MenuLevel = list.MenuLevel
		menuCall.ID = list.ID
		menuCall.ParentId = list.ParentId
		menuCall.Name = list.Name
		menuCall.Path = list.Path
		menuCall.Meta = list.Meta
		menuCall.Children = list.Children
		menuCall.Sort = list.Sort
		menuCall.Component = list.Component
		menuCall.Hidden = list.Hidden
		menuCall.SysRoles = list.SysRoles
		menuCall.CreatedAt = list.CreatedAt
		menuCall.UpdatedAt = list.UpdatedAt
		menuCall.DeletedAt = list.DeletedAt
		mc := menuService.GetMenuCall(menuList, menuCall)
		menuCalls = append(menuCalls, mc)
	}
	menu.Children = menuCalls
	return menu
}
