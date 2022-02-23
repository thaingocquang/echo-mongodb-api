package validation

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/model"
	"echo-mongodb-api/util"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AuthorizeJWT ...
func AuthorizeJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientToken := c.Request().Header.Get("token")

		// parse with claims
		envVars := config.GetEnv()
		token, err := jwt.ParseWithClaims(clientToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(envVars.JWT.SecretKey), nil
		})

		// check token valid
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
			c.Set("email", claims.Email)
			return next(c)
		} else {
			return util.Response400(c, err, "token invalid")
		}

	}
}
