package main

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/route"

	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
	database.Connect()
}

func main() {
	envVars := config.GetEnv()
	e := echo.New()

	// Route ...
	route.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
