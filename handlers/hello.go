package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	Hello string
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
