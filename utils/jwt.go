package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type _claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var secret = []byte("夏天夏天悄悄过去")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	c := _claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    "thedoor",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) string {
	token, err := jwt.ParseWithClaims(tokenString, &_claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(*_claims); ok && token.Valid { // 校验token
		return claims.Username
	}
	return ""
}
