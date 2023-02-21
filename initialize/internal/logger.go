package internal

import (
	"fmt"
	"go-vue-admin/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logzap bool
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logzap = global.GVA_CONFIG.Mysql.LogZap
	default:
		logzap = global.GVA_CONFIG.Mysql.LogZap
	}
	if logzap {
		global.GVA_LOGGER.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
