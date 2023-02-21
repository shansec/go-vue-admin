package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-admin/global"
	"go-vue-admin/initialize"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		fmt.Printf("此处启用 redis")
	}

	Router := initialize.Router()

	address := fmt.Sprintf("%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.GVA_LOGGER.Info("server run success on", zap.String("address", address))
	global.GVA_LOGGER.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}