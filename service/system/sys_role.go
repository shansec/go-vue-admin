package system

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/request"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
)

type RoleService struct{}

// CreateRoleService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateRoleService
// @description: 添加角色
// @param: role system.SysRole
// @return: system.SysRole, error
func (roleService *RoleService) CreateRoleService(role system.SysRole) (system.SysRole, error) {
	err := global.MAY_DB.Where("role_id = ?", role.RoleId).First(&system.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return role, errors.New("存在相同的角色")
	}

	e := global.MAY_DB.Transaction(func(ctx *gorm.DB) error {
		err = ctx.Create(&role).Error
		if err != nil {
			return err
		}

		role.SysBaseMenus = systemReq.DefaultMenu()
		if err = ctx.Model(&role).Association("SysBaseMenus").Replace(&role.SysBaseMenus); err != nil {
			return err
		}

		casbinInfos := systemReq.DefaultCasbin()
		roleId := strconv.Itoa(int(role.RoleId))
		rules := [][]string{}
		for _, casbinInfo := range casbinInfos {
			rules = append(rules, []string{roleId, casbinInfo.Path, casbinInfo.Method})
		}
		return CasbinServiceNew.AddPolicy(ctx, rules)
	})
	return role, e
}

// DeleteRoleService
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteRoleService
// @description: 删除角色
// @param: role *system.SysRole
// @return: error
func (roleService *RoleService) DeleteRoleService(role *system.SysRole) error {
	err := global.MAY_DB.Preload("Users").Where("role_id = ?", role.RoleId).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在！")
	}
	if len(role.Users) != 0 {
		return errors.New("此角色有用户正在使用，无法删除该角色！")
	}
	if err = global.MAY_DB.Where("roles_id", role.RoleId).First(&system.SysUser{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用，无法删除该角色！")
	}
	if err = global.MAY_DB.Where("parent_id = ?", role.RoleId).First(&system.SysRole{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色下还有子角色，无法删除该角色！")
	}
	err = global.MAY_DB.Transaction(func(ctx *gorm.DB) error {
		var err error
		// Unscoped 方法用于移除 GORM 默认的软删除约束
		if err = ctx.Preload("SysBaseMenus").Preload("DataRoleId").Where("role_id = ?", role.RoleId).First(role).Unscoped().Delete(role).Error; err != nil {
			return err
		}
		if len(role.SysBaseMenus) != 0 {
			if err = ctx.Model(role).Association("SysBaseMenus").Delete(role.SysBaseMenus); err != nil {
				return err
			}
		}
		if len(role.DataRoleId) != 0 {
			if err = ctx.Model(role).Association("DataRoleId").Delete(role.DataRoleId); err != nil {
				return err
			}
		}
		if err = ctx.Delete(&system.SysUserRole{}, "sys_role_role_id = ?", role.RoleId).Error; err != nil {
			return err
		}

		roleId := strconv.Itoa(int(role.RoleId))
		err = CasbinServiceNew.RemoveFilteredPolicy(ctx, roleId)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateRoleService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateRoleService
// @description: 更新角色
// @param: (role system.SysRole
// @return: system.SysRole, error
func (roleService *RoleService) UpdateRoleService(role system.SysRole) (system.SysRole, error) {
	var oldRole system.SysRole
	err := global.MAY_DB.Where("role_id = ?", role.RoleId).First(&oldRole).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return system.SysRole{}, errors.New("角色不存在！")
	}
	err = global.MAY_DB.Model(&oldRole).Updates(role).Error
	return role, err
}

// GetRoleListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetRoleListService
// @description: 分页获取角色列表
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (roleService *RoleService) GetRoleListService(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysRole{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return nil, 0, err
	}
	var roles []system.SysRole
	err = db.Limit(limit).Offset(offset).Preload("DataRoleId").Where("parent_id = ?", "0").Find(&roles).Error
	for k := range roles {
		err = roleService.findChildrenRole(&roles[k])
	}
	return roles, total, err
}

// findChildrenRole
// @author: [Shansec](https://github.com/shansec)
// @function: findChildrenRole
// @description: 查找子角色
// @param: role *system.SysRole
// @return: error
func (roleService *RoleService) findChildrenRole(role *system.SysRole) (err error) {
	err = global.MAY_DB.Preload("DataRoleId").Where("parent_id = ?", role.RoleId).Find(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			err = roleService.findChildrenRole(&role.Children[k])
		}
	}
	return err
}

// SetRoleService
// @author: [Shansec](https://github.com/shansec)
// @function: SetRoleService
// @description: 设置角色
// @param: role *system.SysRole
// @return: error
func (roleService *RoleService) SetRoleService(role system.SysRole) error {
	var r system.SysRole
	global.MAY_DB.Preload("DataRoleId").First(&r, "role_id = ?", role.RoleId)
	err := global.MAY_DB.Model(&r).Association("DataRoleId").Replace(&role.DataRoleId)
	return err
}
