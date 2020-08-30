package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
