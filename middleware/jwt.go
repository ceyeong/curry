package middleware

import (
	"os"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

// JWT :
func JWT() echo.MiddlewareFunc {
	return m.JWTWithConfig(m.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: m.AlgorithmHS256,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		Claims:        jwtlib.MapClaims{},
		ContextKey:    "jwt-user",
		Skipper: func(ctx echo.Context) bool {
			print("skip check")
			if ctx.Path() == "api/v1/auth/token" || ctx.Path() == "api/v1/csrf" {
				return true
			}
			return false
		},
		SuccessHandler: func(c echo.Context) {
			//extract user ID from token
			token := c.Get("jwt-user").(*jwtlib.Token)
			claims := token.Claims.(jwtlib.MapClaims)
			userID := claims["user_id"].(string)
			//set it to access globally
			c.Set("user", userID)
		},
		ErrorHandler: func(err error) error {
			if err == m.ErrJWTMissing {
				return echo.ErrUnauthorized
			}
			return err
		},
	})
}
