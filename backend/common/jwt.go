package common

import (
	"blog/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jewKey = []byte("a_secret_crect") // jwt加密的秘钥

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 设置过期时间，7天
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // token 发放时间
			Issuer:    "joey",            // token 发放者
			Subject:   "user token",          // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jewKey) // 使用jwt秘钥来生成token
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jewKey, nil
	})
	return token, claims, err
}
