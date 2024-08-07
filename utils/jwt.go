package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// SECRET 加解密因子
var SECRET string = viper.GetString("Settings.JWT_SECRET")

type MyClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
// GenerateToken 生成 token
func GenerateToken(id string) (string, error) {
	// 创建的 token 有效期是一周
	claims := &MyClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 校验签名是否被篡改
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		//返回密钥与上面签发时保持一致
		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Println("parse token failed ", err)
		//处理token解析后的各种错误
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("TokenMalformed")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("TokenExpired")
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("TokenNotValidYet")
		} else {
			return nil, errors.New("TokenInvalid")
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("TokenInvalid")
}
