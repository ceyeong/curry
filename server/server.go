package server

import (
	"os"

	"github.com/ceyeong/curry/database"
	mid "github.com/ceyeong/curry/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
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
	e.Use(mid.CORS())
	//logger
	e.Use(middleware.Logger())
	//recover
	e.Use(middleware.Recover())
	//session
	e.Use(mid.Session())
	//csrf
	e.Use(mid.CSRF())
	//sessionAuth
	e.Use(mid.SessionAuth())
	//append routes
	route(e)

	e.AutoTLSManager = autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("HOST")),
		Cache:      autocert.DirCache("certs"),
	}

	//start server
	e.Logger.Fatal(e.StartAutoTLS(getHost()))
}
