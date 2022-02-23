package route

import (
	"echo-mongodb-api/controller"
	"echo-mongodb-api/validation"

	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
	players := e.Group("/auth")
	players.POST("/register", controller.Register, validation.PlayerCreate)
	players.POST("/login", controller.Login, validation.Login)
}
