package system

import (
	"errors"
	"fmt"
	"github.com/satori/uuid"
	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/model/system"
	"github/May-cloud/go-vue-admin/utils"
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

func (userService *UserService) Login(u *system.SysUser) (userInfo *system.SysUser, err error) {
	if nil == global.MAY_DB {
		return nil, fmt.Errorf("db not init")
	}
	return nil, err
}
