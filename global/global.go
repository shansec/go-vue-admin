package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/shansec/go-vue-admin/config"
)

var (
	MAY_VP      *viper.Viper
	MAY_CONFIG  config.Server
	MAY_LOGGER  *zap.Logger
	MAY_DB      *gorm.DB
	MAY_CONTROL = &singleflight.Group{}
)
