package jwtx

import "github.com/golang-jwt/jwt/v4"

func GetToken(secretKey string, iat, second, uid int64) (string, error) {
	cliams := make(jwt.MapClaims)
	cliams["exp"] = iat + second
	cliams["iat"] = iat
	cliams["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = cliams
	return token.SignedString([]byte(secretKey))
}
