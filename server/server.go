package server

import (
	"github.com/labstack/echo"

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
