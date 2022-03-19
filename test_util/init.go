package testutil

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/route"

	"github.com/labstack/echo/v4"
)

// InitServer ...
func InitServer() *echo.Echo {
	config.Init()
	// database.Connect()
	ConnectTestDB()
	ClearDB()

	// new server
	e := echo.New()
	route.Route(e)

	return e
}
