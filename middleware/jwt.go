package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/utils"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 token
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.ResetWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		jwt := utils.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.ResetWithDetailed(gin.H{"reload": true}, "授权已过期", ctx)
				ctx.Abort()
				return
			}
			response.ResetWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.MAY_CONFIG.JWT.ExpiresTime
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := jwt.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
