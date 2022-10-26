package handler

import (
	"errors"
	"net/http"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/http/request"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/model"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/store/person"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Person struct {
	Store  person.Person
	Logger *zap.Logger
}

func (p Person) List(c echo.Context) error {
	ps, err := p.Store.Get()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, ps)
}

func (p Person) ByName(c echo.Context) error {
	name := c.Param("name")

	ps, err := p.Store.GetByName(name)
	if err != nil {
		if errors.Is(err, person.ErrPersonNotFound) {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, ps)
}

func (p Person) Create(c echo.Context) error {
	var req request.Person

	logger := p.Logger.With(
		zap.String("real-ip", c.RealIP()),
		zap.String("handler", "create"),
	)

	if err := c.Bind(&req); err != nil {
		logger.Error("request binding failed",
			zap.Error(err),
		)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		logger.Error("request validation failed",
			zap.Error(err),
		)

		return echo.ErrBadRequest
	}

	m := model.Person{
		Name:   req.Name,
		Family: req.Family,
		Age:    req.Age,
	}

	if err := p.Store.Save(m); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, m)
}

func (p Person) Register(g *echo.Group) {
	g.GET("/persons", p.List)
	g.GET("/persons/:name", p.ByName)
	g.POST("/persons", p.Create)
}
