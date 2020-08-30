package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"name": "Dolly!",
		})
	}
}
