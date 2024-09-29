package system

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	model "github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/service/system"
	"github.com/shansec/go-vue-admin/utils"
)

const initOrderUser = initOrderRole + 1

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
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, errors.New("missing db in context")
	}
	password := utils.BcryptHash("admin")
	data := model.SysUser{
		UUID:      uuid.Must(uuid.NewV4()),
		Username:  "admin",
		Password:  password,
		NickName:  "管理员",
		HeaderImg: "https://pic.netbian.com/uploads/allimg/240312/012812-17101780921c39.jpg",
		Sex:       1,
		Email:     "admin@qq.com",
		Phone:     "1234567890",
		RolesId:   888,
		Status:    1,
		DeptsId:   1,
	}
	if err := db.Create(&data).Error; err != nil {
		return ctx, errors.Wrap(err, model.SysUser{}.TableName()+"表初始化失败！")
	}
	next := context.WithValue(ctx, u.InitTableName(), data)
	return next, nil
}

func (u *initUser) TableInited(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.SysUser{})
}

func (u *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var user model.SysUser
	if errors.Is(db.Where("username = ?", "admin").First(&user).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
