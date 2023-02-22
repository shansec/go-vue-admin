package system

import (
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

func (userService *UserService) Login(u *system.SysUser) (err error, userInfo *system.SysUser) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}
	//var user system.SysUser
	u.Password = utils.MD5([]byte(u.Password))
	//err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload()
}
