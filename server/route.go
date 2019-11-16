package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", "application/json")
		return c.String(http.StatusOK, "{\"message\": \"Hello world\"}")
	})
}
