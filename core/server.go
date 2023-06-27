package core

import (
	"fmt"
	"time"

	"github/May-cloud/go-vue-admin/global"
	"github/May-cloud/go-vue-admin/initialize"

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
	global.MAY_LOGGER.Error(s.ListenAndServe().Error())
}
