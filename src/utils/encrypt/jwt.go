package encrypt

import (
	"project/src/config"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var privateKey = []byte(config.Config.Auth.JwtPrivateKey)

//GetPrivateKey 获取私钥的回调函数
var GetPrivateKey = func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return privateKey, nil
}

//Encrypt jwt加密
func Encrypt(mapClaims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("db Encrypt error:", err)
		return ""
	}
	return tokenString
}

//Decrypt jwt解密
func Decrypt(tokenString string) (jwt.MapClaims, error) {
	defer func() (jwt.MapClaims, error) {
		if err := recover(); err != nil {
			return nil, errors.New(fmt.Sprint(err))
		}

		return nil, errors.New("unknow error")
	}()

	if len(tokenString) < 2 {
		return nil, errors.New("token to short")
	}
	token, err := jwt.Parse(tokenString, GetPrivateKey)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
