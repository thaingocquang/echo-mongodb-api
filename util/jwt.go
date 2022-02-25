package util

import (
	"echo-mongodb-api/config"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type jwtCustomClaims struct {
	ID string
	jwt.StandardClaims
}

var envVars = config.GetEnv()

func GenerateUserToken(data map[string]interface{}) string {
	// claims ...
	claims := &jwtCustomClaims{
		data["id"].(primitive.ObjectID).Hex(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	st, err := token.SignedString([]byte(envVars.JWT.SecretKey))

	// if err
	if err != nil {
		return ""
	}

	return st
}
