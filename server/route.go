package server

import (
	"github.com/labstack/echo"

	"github.com/ceyeong/curry/handler"
)

func route(e *echo.Echo) {
	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)

	e.GET("/me", handler.Me)
}
