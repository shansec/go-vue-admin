package utils

import (
	"github/May-cloud/go-vue-admin/global"
	systemReq "github/May-cloud/go-vue-admin/model/system/request"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.MAY_LOGGER.Error("从 jwt 中解析信息失败")
	}
	return claims, nil
}

// GetUID 获取用户的ID
func GetUID(c *gin.Context) uint {
	if claims, existence := c.Get("claims"); !existence {
		if claim, err := GetClaims(c); err != nil {
			return 0
		} else {
			return claim.ID
		}
	} else {
		nextUser := claims.(*systemReq.CustomClaims)
		return nextUser.ID
	}
}
