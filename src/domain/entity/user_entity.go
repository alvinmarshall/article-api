package entity

import (
	"article_api/src/utils/encryption"
	"article_api/src/vo"
)

type User struct {
	BaseEntity
	Name     string
	Username string
	Email    string
	Password string
}

func (user *User) UserOutput() *vo.UserResponse {
	return &vo.UserResponse{
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
