package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func corsConfig() m.CORSConfig {
	corsConfig := m.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	return corsConfig
}

// CORS : CORS Middleware
func CORS() echo.MiddlewareFunc {
	return m.CORSWithConfig(corsConfig())
}

func corsHosts() []string {
	return strings.Split(os.Getenv("CORS_HOST"), ",")
}
