package middleware

import "github.com/labstack/echo/v4"

type (
	// AuthConfig :
	AuthConfig struct {
		UseJWT     bool
		UseSession bool
	}
)

// Auth :
func Auth() echo.MiddlewareFunc {
	config := AuthConfig{UseJWT: true, UseSession: true}
	return AuthWithConfig(config)
}

// AuthWithConfig :
func AuthWithConfig(config AuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.UseJWT {
				//use jwt if authorization header have token
				if c.Request().Header.Get(echo.HeaderAuthorization) != "" {
					//use jwt
					c.Echo().Use(JWT())
					return next(c)
				}
			}
			//use session by default
			//activate csrf
			c.Echo().Use(CSRF())
			//append session auth
			c.Echo().Use(SessionAuth())
			return next(c)
		}
	}
}
