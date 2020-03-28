package middleware

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Session : Session Middleware
func Session() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET"))))
}
