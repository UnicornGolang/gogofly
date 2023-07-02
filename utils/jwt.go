package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// =========================================
// 定义 jwt 生辰于解密的逻辑
// =========================================

// 自定义 jwt 中自定义区域的数据
type JwtCustomClaims struct {
	Id   int
	Name string
	jwt.RegisteredClaims
}

// 用于加密数据的 key
var signedKey = []byte(viper.GetString("jwt.signedKey"))

// 生成 jwt 令牌
func GenerateToken(id int, name string) (string, error) {
	jwtClaims := JwtCustomClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(signedKey)
}

// 解析 jwt 令牌
func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	//
	jwtClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return jwtClaims, err
}

// 校验 token 是否有效
func InValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
