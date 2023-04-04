package middleware

import (
	"mall/global"
	"mall/respond"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SigningKey =[]byte(global.Config.Jwt.SigningKey) 
type Claims struct {
	ID 		  uint 	`json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}
//jwt中间件
func JWT()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		code := 200
		token := ctx.GetHeader("authorization")
		if token == ""{
			code = 403
		}else{
			claims,err :=ParseToken(token)
			if err != nil {
				code = respond.ErrorAuthCheckTokenFail
			}else if time.Now().Unix() > claims.ExpiresAt{
				code = respond.ErrorAuthCheckTokenTimeout
			}
		}
		if code != 200{
			ctx.JSON(200,
			gin.H{
				"Code":code,
				"err":"token_err",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}



func GenerateToken(userid uint,username string,authority int) (string,error) {
	claims := Claims{
		ID:userid,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(6 * time.Hour).Unix(),
			Issuer:    "mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(SigningKey)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}