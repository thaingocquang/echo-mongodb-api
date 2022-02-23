package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body LoginBody) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Email,
			validation.Required.Error("Email is required"),
			validation.Length(10, 50),
		),
	)
}
