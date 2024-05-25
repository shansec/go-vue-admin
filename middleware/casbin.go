package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/service/system"
	"github/shansec/go-vue-admin/utils"
)

func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUse, err := utils.GetClaims(c)
		if err != nil {
			response.FailWithMessage("非法访问！", c)
			c.Abort()
		}
		// 获取请求的 path
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.MAY_CONFIG.System.RouterPrefix)
		// 获取请求的方法
		method := c.Request.Method
		// 获取用户的所有角色
		role := strconv.Itoa(int(requestUse.RoleId))

		// 判断策略是否存在
		enforce := system.CasbinServiceNew.Casbin()
		success, _ := enforce.Enforce(obj, method, role)
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
