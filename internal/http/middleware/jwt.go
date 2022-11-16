package middleware

import (
	"errors"
	"fmt"
	"strings"

	cloudJWT "github.com/1995parham-teaching/cloud-fall-2022/internal/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var ErrUnexpectedSingingMethod = errors.New("unexpected signing method")

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		auths := strings.Fields(auth)

		if len(auths) != 2 {
			return echo.ErrUnauthorized
		}

		if auths[0] != "Bearer" {
			return echo.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(auths[1], new(jwt.RegisteredClaims), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrUnexpectedSingingMethod
			}

			return []byte(cloudJWT.Secret), nil
		})
		if err != nil {
			return echo.ErrUnauthorized
		}

		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			c.Set("username", claims.Subject)
			fmt.Println(c.Get("username").(string))
		} else {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
