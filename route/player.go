package route

import (
	"echo-mongodb-api/controller"
	"echo-mongodb-api/validation"

	"github.com/labstack/echo/v4"
)

func player(e *echo.Echo) {
	players := e.Group("/players")
	players.POST("", controller.PlayerCreate, validation.PlayerCreate)
	players.GET("/my-profile", controller.MyProfile, validation.AuthorizeJWT)
}
