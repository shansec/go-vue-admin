package request

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/satori/uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
}
