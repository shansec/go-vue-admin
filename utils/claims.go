package utils

import (
	"github/shansec/go-vue-admin/global"
	systemReq "github/shansec/go-vue-admin/model/system/request"

	"github.com/gin-gonic/gin"
	"github.com/satori/uuid"
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

// GetUID 根据 jwt 签证获取用户的ID
func GetUseID(c *gin.Context) uint {
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
// GetUseUuid 根据 jwt 签证获取用户的uuid
func GetUseUuid(c *gin.Context) uuid.UUID {
	if claims, existence := c.Get("claims"); !existence {
		if claim, err := GetClaims(c); err != nil {
			return claim.UUID
		} else {
			return claim.UUID
		}
	} else {
		nextUser := claims.(*systemReq.CustomClaims)
		return nextUser.UUID
	}
}
