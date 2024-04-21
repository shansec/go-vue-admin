package core

import (
	"fmt"
	"time"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.MAY_CONFIG.System.UseMultipoint || global.MAY_CONFIG.System.UseRedis {
		fmt.Printf("此处启用 redis")
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.MAY_CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.MAY_LOGGER.Info("server run success on", zap.String("address", address))

	fmt.Printf(`go-vue-admin 启动成功`)

	fmt.Printf(`
	欢迎使用 go-vue-admin
	当前版本:v1.0.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.MAY_LOGGER.Error(s.ListenAndServe().Error())
}
