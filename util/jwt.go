package util

import (
	"mall/global"
	"time"

	"github.com/golang-jwt/jwt"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToke(username string) (string, error) {
	claims := Claims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer: username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}

//邮箱验证

type EmailClaims struct {
	UserId	uint `json:"userid"`
	UserEmail string `json:"useremail"`
	UserPasswd string `json:"userpasswd"`
	OperationType uint `json:"operationType"`
	jwt.StandardClaims
}
func EmailGenrateToken(useremail,userpasswd string,  Operationtype,userid uint)(string, error)  {
	emailclaims := EmailClaims{
		UserId: userid,
		UserEmail: useremail,
		UserPasswd: userpasswd,
		OperationType: Operationtype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60,
			Issuer: "Tesla",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, emailclaims)
	return token.SignedString(SigningKey)
}
func EmailVerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}