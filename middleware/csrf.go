package middleware

import (
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func csrfConfig() m.CSRFConfig {
	csrfConfig := m.CSRFConfig{
		ContextKey:   "csrf",
		CookiePath:   "/",
		CookieName:   "XSRF-TOKEN",
		TokenLookup:  "header:X-XSRF-TOKEN",
		CookieMaxAge: 86400,
		TokenLength:  32,
	}
	return csrfConfig
}

// CSRF : csrf middlewared
func CSRF() echo.MiddlewareFunc {
	return m.CSRFWithConfig(csrfConfig())
}
