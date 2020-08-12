package main

import (
	"html/template"
	"io"

	"github.com/d-kusk/sandbox-echo/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("", handlers.Index()).Name = "top"
	e.GET("/hello", handlers.Hello()).Name = "hello"

	e.Logger.Fatal(e.Start(":1323"))
}
