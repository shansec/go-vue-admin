package system

import (
	"errors"
	"fmt"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"gorm.io/gorm"
	"time"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/utils"

	"github.com/satori/uuid"
)

type UserService struct{}

const USER_STATUS = "2"

// Login
// @author: [Shansec](https://github.com/shansec)
// @function: Login
// @description: 用户登录
// @param: u *system.SysUser
// @return: userInfo *system.SysUser, err error
func (userService *UserService) Login(u *system.SysUser) (userInfo *system.SysUser, err error) {
	if nil == global.MAY_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.MAY_DB.Where("username = ?", u.Username).Preload("SysDept").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		return &user, nil
	}
	return &user, err
}

// ChangePassword
// @author: [Shansec](https://github.com/shansec)
// @function: ChangePassword
// @description: 修改密码
// @param: u *system.SysUser, newPassword string
// @return: userInfo *system.SysUser, err error
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.MAY_DB.Where("id = ?", u.ID).First(&user).Error
	if err == nil {
		if passIsRight := utils.BcryptCheck(u.Password, user.Password); !passIsRight {
			return nil, errors.New("原密码有误")
		}
		user.Password = utils.BcryptHash(newPassword)
		err = global.MAY_DB.Save(&user).Error
		return &user, err
	}
	return nil, errors.New("非法访问")
}

// Register
// @author: [Shansec](https://github.com/shansec)
// @function: Register
// @description: 用户注册
// @param: u system.SysUser
// @return: userInfo system.SysUser, err error
func (userService *UserService) Register(u system.SysUser) (userInfo system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.MAY_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInfo, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.MAY_DB.Create(&u).Error
	return u, err
}

// DelUserInformation
// @author: [Shansec](https://github.com/shansec)
// @function: DelUserInformation
// @description: 删除用户信息
// @param: uuid uuid.UUID
// @return: err error
func (userService *UserService) DelUserInformation(uuid uuid.UUID) error {
	var user system.SysUser
	err := global.MAY_DB.Where("uuid = ?", uuid).Delete(&user).Error
	if err != nil {
		return errors.New("删除用户信息失败")
	}
	return nil
}

// UpdateUserInformation
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateUserInformation
// @description: 更改用户信息
// @param: userInfo *system.SysUser
// @return: err error
func (userService *UserService) UpdateUserInformation(userInfo *system.SysUser) error {
	var user system.SysUser
	err := global.MAY_DB.Model(&user).
		Select("updated_at", "nick_name", "header_img", "phone", "email").
		Where("uuid = ?", userInfo.UUID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  userInfo.NickName,
			"header_img": userInfo.HeaderImg,
			"phone":      userInfo.Phone,
			"email":      userInfo.Email,
		}).Error
	if err != nil {
		return errors.New("更新用户信息失败")
	} else {
		return nil
	}
}

// GetUserInformation
// @author: [Shansec](https://github.com/shansec)
// @function: GetUserInformation
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: userInfo *system.SysUser, err error
func (userService *UserService) GetUserInformation(uuid uuid.UUID) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.MAY_DB.Preload("SysDept").Where("uuid = ?", uuid).First(&user).Error
	if err == nil {
		return &user, nil
	}
	return nil, errors.New("获取用户信息失败")
}

// GetUsersInformation
// @author: [Shansec](https://github.com/shansec)
// @function: GetUsersInformation
// @description: 获取用户列表
// @param: nil
// @return: usersInfo []system.SysUser, err error
func (userService *UserService) GetUsersInformation(info systemReq.GetUserList) (usersInfo []system.SysUser, total int64, err error) {
	var users []system.SysUser
	limit := info.PagSize
	offset := info.PagSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysUser{})
	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Limit(limit).Offset(offset).Preload("SysDept").Find(&users).Error
	if err != nil {
		return nil, 0, errors.New("获取用户列表失败")
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, errors.New("获取用户列表失败")
	}
	return users, total, nil
}

// UpdateStatus
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateStatus
// @description: 更改用户状态
// @param: uuid uuid.UUID
// @return: err error
func (userService *UserService) UpdateStatus(uuid uuid.UUID) (err error) {
	var user system.SysUser
	if err = global.MAY_DB.Where("uuid = ?", uuid).First(&user).Error; err == nil {
		if user.Status == USER_STATUS {
			global.MAY_DB.Model(&user).Where("uuid = ?", uuid).Update("status", "1")
		} else {
			global.MAY_DB.Model(&user).Where("uuid = ?", uuid).Update("status", "2")
		}
		return nil
	} else {
		return err
	}
}
