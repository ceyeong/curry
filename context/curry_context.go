package context

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// CurryContext : Custom curry context
type CurryContext struct {
	echo.Context
}

//GetFromSession : Gets Value From Session
func (c *CurryContext) GetFromSession(key string) (interface{}, error) {
	s, err := session.Get("session", c)
	if err != nil {
		return s, err
	}
	return s.Values[key], nil
}

//SetToSession : Set value to Session
func (c *CurryContext) SetToSession(key string, value interface{}) error {
	s, err := session.Get("session", c)
	if err != nil {
		return err
	}
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
	}
	s.Values[key] = value
	s.Save(c.Request(), c.Response())
	return nil
}

// ClearSession : clears all current session values
func (c *CurryContext) ClearSession() error {
	s, err := session.Get("session", c)
	if err != nil {
		return err
	}
	s.Values = map[interface{}]interface{}{}
	s.Save(c.Request(), c.Response())
	return nil
}
