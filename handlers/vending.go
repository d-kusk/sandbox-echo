package handlers

import (
	"net/http"

	"github.com/dsk52/sandbox-echo/db"
	"github.com/dsk52/sandbox-echo/models"
	"github.com/labstack/echo"
)

func List() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.DBConnection()
		defer conn.Close()

		var _drinks models.Drinks
		conn.Find(&_drinks)

		return c.JSON(http.StatusOK, _drinks)
	}
}
