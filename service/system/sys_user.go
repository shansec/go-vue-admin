package system

import (
	"errors"
	"fmt"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/utils"

	"github.com/satori/uuid"
	"gorm.io/gorm"
)

type UserService struct{}

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
	err = global.MAY_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		return &user, nil
	}
	return &user, err
}

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

// @author: [Shansec](https://github.com/shansec)
// @function: GetUserInformation
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: userInfo *system.SysUser, err error

func (userService *UserService) GetUserInformation(uuid uuid.UUID) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.MAY_DB.Where("uuid = ?", uuid).First(&user).Error
	if err == nil {
		return &user, nil
	}
	return nil, errors.New("获取用户信息失败")
}
