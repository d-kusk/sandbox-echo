package handlers

import (
	"net/http"
	"sandbox-echo/db"
	"sandbox-echo/models"

	"github.com/labstack/echo/v4"
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
