package server

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/ceyeong/curry/handler"
)

func route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", "application/json")
		return c.String(http.StatusOK, "{\"message\": \"Hello world\"}")
	})

	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)
	e.GET("/me", handler.Me)
}
