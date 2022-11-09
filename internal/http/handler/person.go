package handler

import (
	"errors"
	"net/http"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/http/request"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/model"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/store/person"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Person struct {
	Store  person.Person
	Logger *zap.Logger
	Tracer trace.Tracer
}

func (p Person) List(c echo.Context) error {
	_, span := p.Tracer.Start(c.Request().Context(), "persons.list")
	defer span.End()

	span.SetAttributes(attribute.String("username", c.Get("username").(string)))

	ps, err := p.Store.Get()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
