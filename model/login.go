package model

import (
	"regexp"

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
			validation.Required.Error("email is required"),
			validation.Length(3, 254).Error("email must contain 3-254 characters"),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")).Error("email not valid"),
		),
		validation.Field(
			&body.Password,
			validation.Required.Error("password is required"),
			validation.Length(3, 254).Error("password must contain 3-254 characters"),
		),
	)
}
