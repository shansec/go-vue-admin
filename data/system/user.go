package system

import (
	"context"
	"errors"
	model "github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/service/system"
	"gorm.io/gorm"
)

const initOrderUser = InitOrder + 1

type initUser struct{}

func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (u *initUser) InitTableName() string {
	return model.SysUser{}.TableName()
}

func (u *initUser) MigrateTable(ctx context.Context) (cont context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	return ctx, db.AutoMigrate(&model.SysUser{})

}

func (u *initUser) InitData(ctx context.Context) (cont context.Context, err error) {
	//db, ok := ctx.Value("db").(*gorm.DB)
	//if !ok {
	//	return ctx, errors.New("missing db in context")
	//}
	//password := utils.BcryptHash("admin")
	//data := model.SysUser{
	//	Username:  "admin",
	//	Password:  password,
	//	NickName:  "管理员",
	//	HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg",
	//	Sex:       1,
	//	Email:     "admin@qq.com",
	//	Phone:     "1234567890",
	//	RolesId:   1,
	//	Status:    1,
	//	DeptsId:   1,
	//}
	return ctx, nil
}

func (u *initUser) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysUser{})
}

func (u *initUser) DataInserted(ctx context.Context) bool {
	return false
}
