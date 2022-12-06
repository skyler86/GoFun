package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTsecret = []byte("ABAB")		// jwt的密钥

type  Claims struct {
	// 加密的数据
	Id uint `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 创建完claims后再对token进行签发
// 1.先签发(生成)一个token
func GenerateToken(id uint,username,password string)(string,error)  {
	notTime := time.Now()		// 先设置时间戳，因为token会有过期
	expireTime := notTime.Add(24*time.Hour)		// 过期时间为24个小时
	claims := Claims{
		Id: id,
		UserName: username,
		Password: password,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),		// 将时间给放进去
			Issuer: "todoList",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err := tokenClaims.SignedString(JWTsecret)
	return token,err

}

// 2.验证token
func ParseToken(token string) (*Claims,error)  {		// 把token传进来进行验证
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}