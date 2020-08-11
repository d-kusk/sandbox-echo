package main

import (
	"github.com/d-kusk/sandbox-echo/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/", handler.Hello()).Name = "root"
	e.GET("/hello", handler.Hello()).Name = "hellp"

	e.Logger.Fatal(e.Start(":1323"))
}
