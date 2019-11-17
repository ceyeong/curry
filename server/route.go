package server

import (
	"github.com/labstack/echo"

	"github.com/ceyeong/curry/handler"
)

func route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1", jwt())

	apiV1.POST("/register", handler.RegisterUser)
	apiV1.POST("/login", handler.LoginUser)

	apiV1.GET("/me", handler.Me)
}
