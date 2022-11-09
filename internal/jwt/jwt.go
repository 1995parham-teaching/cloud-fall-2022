package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const Secret = "Hello"

func Create(username string) string {
	// nolint: exhaustivestruct
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:   "cloud",
		Subject:  username,
		IssuedAt: jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}
