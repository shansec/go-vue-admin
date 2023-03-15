package utils

import (
	"github.com/natefinch/lumberjack"
	"github/May-cloud/go-vue-admin/global"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		// 日志文件的位置
		Filename: file,
		// 在切割之前，日志文件的最大大小
		MaxSize: 10,
		// 保留旧文件的最大个数
		MaxBackups: 200,
		// 保留旧文件的最大天数
		MaxAge: 30,
		// 是否压缩
		Compress: true,
	}

	if global.MAY_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
