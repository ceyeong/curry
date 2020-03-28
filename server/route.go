package server

import (
	"github.com/labstack/echo/v4"

	"github.com/ceyeong/curry/handler"
)

func route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	apiV1.POST("/register", handler.RegisterUser)
	apiV1.POST("/login", handler.LoginUser)
	apiV1.POST("/logout", handler.LogoutUser)

	apiV1.GET("/me", handler.Me)
	apiV1.GET("/csrf", handler.Csrf)
}
