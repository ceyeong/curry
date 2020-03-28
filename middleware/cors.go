package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func corsConfig() m.CORSConfig {
	corsConfig := m.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	return corsConfig
}

// CORS : CORS Middleware
func CORS() echo.MiddlewareFunc {
	return m.CORSWithConfig(corsConfig())
}
