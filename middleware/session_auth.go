package middleware

import (
	"github.com/ceyeong/curry/context"
	"github.com/labstack/echo/v4"
)

type (
	// SessionAuthConfig : SessionAuthConfig
	SessionAuthConfig struct {
		SkipPath []string
	}
)

// SessionAuth : SessionAuth Middleware with default config
func SessionAuth() echo.MiddlewareFunc {
	config := SessionAuthConfig{
		SkipPath: []string{"/api/v1/login", "/api/v1/register", "/api/v1/csrf"},
	}
	return SessionAuthWithConfig(config)
}

// SessionAuthWithConfig :
func SessionAuthWithConfig(config SessionAuthConfig) echo.MiddlewareFunc {
	if config.SkipPath == nil {
		config.SkipPath = []string{}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*context.CurryContext)
			for _, path := range config.SkipPath {
				if path == c.Path() {
					return next(c)
				}
			}
			user, err := cc.GetFromSession("userID")
			if err != nil {
				return err
			}
			if user == nil {
				return echo.ErrUnauthorized
			}
			c.Set("user", user.(string))
			return next(c)
		}
	}
}
