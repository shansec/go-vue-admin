package initialize

import "github.com/gin-gonic/gin"

// Router 初始化总路由
func Router() *gin.Engine {
	Router := gin.Default()

	return Router
}