package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// PlayerCreateBody ...
	PlayerCreateBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	PlayerProfile struct {
		Name      string    `bson:"name"`
		Email     string    `bson:"email"`
		CreatedAt time.Time `bson:"createdAt"`
	}
)

// Validate PlayerCreatePayload
func (body PlayerCreateBody) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Name,
			validation.Required.Error("name is required"),
			validation.Length(3, 30).Error("name is length: 3 -> 30"),
		),
		validation.Field(
			&body.Email,
			validation.Required.Error("email is required"),
			is.Email,
		),
	)
}

// ConvertToBSON ....
func (body PlayerCreateBody) ConvertToBSON() PlayerBSON {
	result := PlayerBSON{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Email:     body.Email,
		Password:  body.Password,
		CreatedAt: time.Now(),
	}
	return result
}
