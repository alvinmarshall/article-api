package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDtoWithValidInput(t *testing.T) {
	asserts := assert.New(t)
	input := &UserInput{
		Name:     "any name",
		Username: "any username",
		Password: "any password",
		Email:    "any@me.com"}
	err := input.Validate()
	asserts.Nil(err)
}

func TestUserDtoWithInvalidInput(t *testing.T) {
	asserts := assert.New(t)
	input := &UserInput{
		Name:     "",
		Username: "",
		Password: "",
		Email:    "",
	}
	err := input.Validate()
	asserts.NotNil(err)
}

func TestUserDtoWithPasswordLessThan4(t *testing.T) {
	asserts := assert.New(t)
	input := &UserInput{
		Name:     "any name",
		Username: "any username",
		Password: "any",
		Email:    "any@me.com"}
	err := input.Validate()
	asserts.NotNil(err)
}
