package server

import (
	"net/http"
	"os"

	"github.com/ceyeong/curry/database"
	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	//logger
	e.Use(middleware.Logger())
	//recover
	e.Use(middleware.Recover())

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
