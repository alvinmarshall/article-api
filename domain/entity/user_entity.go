package entity

import (
	"article_api/model/dto"
	"article_api/utils/encryption"
)

type User struct {
	BaseEntity
	Name     string
	Username string
	Email    string
	Password string
}

func (user *User) UserOutput() *dto.UserOutput {
	return &dto.UserOutput{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (user *User) BeforeSave() error {
	err := user.getHashPassword(user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) getHashPassword(password string) error {
	hashPassword, err := encryption.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return nil
}
