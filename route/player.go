package route

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

func player(e *echo.Echo) {
	players := e.Group("/players")
	players.GET("/my-profile", controller.MyProfile, middleware.JWT([]byte(envVars.JWT.SecretKey)))
}
