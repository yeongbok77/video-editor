// Package jwt 鉴权相关操作
package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySecret = []byte("videoeditor")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

//生成token
func GenToken(userID int64) (aToken, rToken string, err error) {
	c := MyClaims{
		UserID: userID, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(8760) * time.Hour).Unix(), // 过期时间
			Issuer:    "douyinapp",
		},
	}

	// 加密并获得完整的编码后的字符串Token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret) // 注意 这里的加密算法一定是 HS256

	// refresh token  不需要任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "videoeditor",
	}).SignedString(mySecret)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}
