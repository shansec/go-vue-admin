package global

import (
	"github.com/spf13/viper"
	"go-vue-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_LOGGER *zap.Logger
	GVA_DB     *gorm.DB
)
