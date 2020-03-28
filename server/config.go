package server

import (
	"github.com/ceyeong/curry/context"
	"github.com/labstack/echo/v4"
)

//Custom context
func curryContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CurryContext{Context: c}
		return next(cc)
	}
}
