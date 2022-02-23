package controller

import (
	"echo-mongodb-api/model"
	"echo-mongodb-api/service"
	"echo-mongodb-api/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// PlayerCreate ...
func PlayerCreate(c echo.Context) error {
	var (
		body = c.Get("body").(model.PlayerCreateBody)
	)

	// Process data
	rawData, err := service.PlayerCreate(body)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

func MyProfile(c echo.Context) error {
	var (
		email = c.Get("email").(string)
	)

	// Process data
	doc, err := service.PlayerProfileByEmail(email)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, doc, "")
}
