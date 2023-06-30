package utils

import (
	"errors"
	"time"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system/request"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.MAY_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.MAY_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.MAY_CONFIG.JWT.AExpiresTime,
			Issuer:    global.MAY_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error, string, error) {
	// 生成 access_token
	aToken, aError := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	// 生成 refresh_token
	rToken, rError := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + global.MAY_CONFIG.JWT.RExpiresTime,
		Issuer:    global.MAY_CONFIG.JWT.Issuer,
	}).SignedString(j.SigningKey)
	return aToken, aError, rToken, rError
}

func (j *JWT) RefreshToken(aTokenString, rTokenString string) (string, string, error) {
	// 判断 refresh_token 是否过期
	_, err := jwt.Parse(rTokenString, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", "", err
	}
	// 从旧的 access_token 中解析数据
	_, err = jwt.ParseWithClaims(aTokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", "", TokenMalformed
			} else if v.Errors*jwt.ValidationErrorClaimsInvalid != 0 {
				claims := j.CreateClaims(request.BaseClaims{
					ID:       request.CustomClaims{}.ID,
					UUID:     request.CustomClaims{}.UUID,
					NickName: request.CustomClaims{}.NickName,
					Username: request.CustomClaims{}.Username,
				})
				aToken, _, rToken, _ := j.CreateToken(claims)
				return aToken, rToken, nil
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", "", TokenNotValidYet
			} else {
				return "", "", TokenInvalid
			}
		}
	}
	return "", "", err
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors*jwt.ValidationErrorClaimsInvalid != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
