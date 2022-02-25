package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the password in the DB.
func VerifyPassword(userPassword string, providedPassword string) bool {
	check := true
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword)); err != nil {
		check = false
	}
	return check
}
