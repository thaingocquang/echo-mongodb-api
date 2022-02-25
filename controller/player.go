package controller

import (
	"echo-mongodb-api/service"
	"echo-mongodb-api/util"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func MyProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Process data
	doc, err := service.PlayerProfileFindByID(claims["ID"].(string))

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, doc, "")
}
