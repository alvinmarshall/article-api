package dto

import (
	"gopkg.in/go-playground/validator.v9"
)

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4"`
}

func (input *UserInput) Validate() error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	return nil
}

type UserOutput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
