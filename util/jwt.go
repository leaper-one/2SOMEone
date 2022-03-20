package util

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SecretKey = "2SOMEone.one"

func GenerateToken(user_id string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		UserID: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "LEAPERone",
		},
	})

	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(SecretKey))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*LoginClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

type LoginClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
