package middleware

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID 		  uint 	`json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}
func Jwt()  {
	
}