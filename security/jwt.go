package security

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func SignJWTToken(message string) (string, error) { 
	// Token instantiate using jwt 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"message" : message, 
		"exp" : time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign it using the obtained secret key
	tokenString, err := token.SignedString([]byte(SECRET_KEY)) 
	if err != nil {
		return "", err
	} else {
		return tokenString, nil 
	}
}

func VerifyJWTToken(token string) (bool, error) {
	// Parse the token
	if _ , err := jwt.Parse(token, func(token * jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil 
	}); err != nil {
		return false, err 
	} else {
		return true, nil 
	}
}