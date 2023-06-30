package middleware

import (
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 token
		aToken := ctx.Request.Header.Get("Authorization-aToken")
		rToken := ctx.Request.Header.Get("Authorization-rToken")
		if aToken == "" || rToken == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		jwt := utils.NewJWT()
		claims, err := jwt.ParseToken(aToken)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", ctx)
				ctx.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
