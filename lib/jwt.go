package lib

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func JWTdecode(tokenString string) (*jwt.Token, error) {
	var token *jwt.Token
	var err error
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("[%v] Invalid token", err)
	}
	return token, nil
}
