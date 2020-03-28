package server

import (
	"net/http"
	"os"

	"github.com/ceyeong/curry/context"
	"github.com/ceyeong/curry/database"
	mid "github.com/ceyeong/curry/middleware"
	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Start : Starts server
func Start() {
	e := echo.New()

	//load environment variables
	if err := godotenv.Load(); err != nil {
		e.Logger.Fatalf("Failed to load environment variables. %s", err.Error())
	}

	//initialize database instance
	if err := database.InitDatabase(); err != nil {
		e.Logger.Fatalf("Failed to Initalize database.\n %s", err.Error())
	}
	//middlewares
	//custom context
	e.Use(curryContext)
	//cors enable
	corsConfig := middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost.com"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	corsConfig = middleware.DefaultCORSConfig
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	e.Use(middleware.CORSWithConfig(corsConfig))
	//logger
	e.Use(middleware.Logger())
	//recover
	e.Use(middleware.Recover())
	//session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	//csrf
	csrfConfig := middleware.CSRFConfig{
		ContextKey:   "csrf",
		CookiePath:   "/",
		CookieName:   "XSRF-TOKEN",
		TokenLookup:  "header:X-XSRF-TOKEN",
		CookieMaxAge: 86400,
		TokenLength:  32,
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))
	//sessionAuth
	e.Use(mid.SessionAuth())

	//append routes
	route(e)

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}

// initialize Jwt middleware and return
func jwt() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(err error) error {
			if err == middleware.ErrJWTMissing {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			//todo check by type
			if err.Error() == "Token is expired" {
				return echo.NewHTTPError(http.StatusUnauthorized, "token expired")
			}
			return err
		},
		SuccessHandler: func(c echo.Context) {
			//extract user ID from token
			token := c.Get("user").(*jwtlib.Token)
			claims := token.Claims.(jwtlib.MapClaims)
			userID := claims["user_id"].(string)
			//set it to access globally
			c.Set("user", userID)
		},
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/login" || c.Path() == "/api/v1/register" {
				return true
			}
			return false
		},
	}
	return middleware.JWTWithConfig(config)
}

//Custom context
func curryContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CurryContext{c}
		return next(cc)
	}
}
