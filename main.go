package main

import (
	"log"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/config"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/http/handler"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/http/middleware"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/store/person"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/tracing"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	log.Printf("%+v", cfg)

	app := echo.New()

	_ = tracing.New(cfg.Tracing)

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	p := handler.Person{
		Store:  person.NewInMemory(logger.Named("store")),
		Logger: logger.Named("http"),
	}
	p.Register(app.Group("", middleware.JWT))

	app.GET("/hello", handler.Hello)

	if err := app.Start("127.0.0.1:1234"); err != nil {
		log.Fatal(err)
	}
}
