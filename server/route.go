package server

import (
	"github.com/labstack/echo/v4"

	"github.com/ceyeong/curry/handler"
)

func route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	apiV1.POST("/register", handler.RegisterUser).Name = "register"
	apiV1.POST("/login", handler.LoginUser).Name = "login"
	apiV1.POST("/logout", handler.LogoutUser).Name = "logout"

	apiV1.GET("/me", handler.Me).Name = "me"
	apiV1.GET("/csrf", handler.Csrf).Name = "csrf"
	apiV1.POST("/auth/token", handler.Token).Name = "token"
}
