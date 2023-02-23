package system

import (
	"errors"
	"fmt"
	"go-vue-admin/global"
	"go-vue-admin/model/system"
	"go-vue-admin/utils"
)

type UserService struct{}

// @author: may
// @function: Login
// @description: 用户登录
// @param: u *model.SysUser
// @return: err error, userInfo *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (userInfo *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	//u.Password = utils.MD5([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
			return nil, errors.New("密码错误")
		}
		return &user, nil
	}
	return &user, err
}
