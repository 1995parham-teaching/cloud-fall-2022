package handler

import (
	"net/http"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/jwt"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, jwt.Create("parham.alvani"))
}
