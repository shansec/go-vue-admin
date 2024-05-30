package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"

	"github/shansec/go-vue-admin/global"
	systemReq "github/shansec/go-vue-admin/model/system/request"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.MAY_LOGGER.Error("从 jwt 中解析信息失败")
	}
	return claims, nil
}

// GetUseID 根据 jwt 签证获取用户的ID
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

func GetUserRoleId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.RoleId
	}
}
