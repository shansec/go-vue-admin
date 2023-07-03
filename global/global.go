package global

import (
	"github/shansec/go-vue-admin/config"
	"golang.org/x/sync/singleflight"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MAY_VP      *viper.Viper
	MAY_CONFIG  config.Server
	MAY_LOGGER  *zap.Logger
	MAY_DB      *gorm.DB
	MAY_CONTROL = &singleflight.Group{}
)
