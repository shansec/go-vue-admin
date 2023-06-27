package global

import (
	"github/May-cloud/go-vue-admin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MAY_VP     *viper.Viper
	MAY_CONFIG config.Server
	MAY_LOGGER *zap.Logger
	MAY_DB     *gorm.DB
)
