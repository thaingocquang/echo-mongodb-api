package validation

import (
	"echo-mongodb-api/model"
	"echo-mongodb-api/util"

	"github.com/labstack/echo/v4"
)

// PlayerCreate ...
func PlayerCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.PlayerCreateBody
		)

		// ValidateStruct
		c.Bind(&body)
		err := body.Validate()

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
