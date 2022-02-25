package validation

import (
	"echo-mongodb-api/model"
	"echo-mongodb-api/util"

	"github.com/labstack/echo/v4"
)

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.LoginBody
		)

		// ValidateStruct
		c.Bind(&body)

		//if err
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
