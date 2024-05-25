package system

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"github/shansec/go-vue-admin/global"
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
	if err = global.MAY_DB.Where("role_id", role.RoleId).First(&system.SysUser{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用，无法删除该角色！")
	}
	if err = global.MAY_DB.Where("parent_id = ?", role.ParentId).First(&system.SysRole{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
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
