package model

import "github.com/golang-jwt/jwt"

// SignedDetails
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
