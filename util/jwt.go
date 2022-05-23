package util

import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

var SecretKey = "2SOMEone.one"

type AuthToken struct {
	Token string
}

func (c AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c AuthToken) RequireTransportSecurity() bool {
	return false
}

func GenerateToken(user_id, phone string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		UserID: user_id,
		Phone:  phone,
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

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no token")
	}
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", errors.New("no token in metadata")
	}
	return token[0], nil
}

func CheckAuth(ctx context.Context) (string, error) {
	token, err := getTokenFromContext(ctx)
	if err != nil {
		return "", err
	}
	claims, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	return claims.UserID, nil
}

type LoginClaims struct {
	UserID string `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}
