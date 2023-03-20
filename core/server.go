package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/May-cloud/go-vue-admin/global"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.MAY_CONFIG.System.UseMultipoint || global.MAY_CONFIG.System.UseRedis {
		fmt.Printf("此处启用 redis")
	}

	address := fmt.Sprintf("%d", global.MAY_CONFIG.System.Addr)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	time.Sleep(10 * time.Microsecond)
	global.MAY_LOGGER.Info("server run success on", zap.String("address", address))
	initServer(address, r)
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
