package controller

import (
	"echo-mongodb-api/model"
	"echo-mongodb-api/service"
	"echo-mongodb-api/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Register ...
func Register(c echo.Context) error {
	var (
		body = c.Get("body").(model.PlayerCreateBody)
	)

	// Process data
	rawData, err := service.Register(body)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

func Login(c echo.Context) error {
	var (
		body = c.Get("body").(model.LoginBody)
	)

	// Process data
	token, err := service.Login(body)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, nil, token)
}
