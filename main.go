package main

import (
	"log"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/http/handler"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/store/person"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	app := echo.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	p := handler.Person{
		Store:  person.NewInMemory(logger.Named("store")),
		Logger: logger.Named("http"),
	}
	p.Register(app.Group(""))

	app.GET("/hello", handler.Hello)

	if err := app.Start("127.0.0.1:1234"); err != nil {
		log.Fatal(err)
	}
}
