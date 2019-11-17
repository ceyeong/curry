package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ceyeong/curry/database"

	"github.com/joho/godotenv"
)

// Start : Starts server
func Start() {
	e := echo.New()

	//load environment variables
	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal("Failed to load environment variables")
	}

	//initialize database instance
	if err := database.InitDatabase(); err != nil {
		e.Logger.Fatal("Failed to Initalize database")
	}

	//append routes
	route(e)

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}

// initialize Jwt middleware and return
func jwt() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("secret")),
		ErrorHandler: func(err error) error {
			if err == middleware.ErrJWTMissing {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			return err
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
