package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySecret = []byte("海韵")

const TokenExpireDuration = time.Hour * 24

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenRegisteredClaims 使用默认声明创建jwt
func GenToken(userId int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		userId,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "Mark", // 签发人
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
