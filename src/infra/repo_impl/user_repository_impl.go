package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

var _ repository.UserRepository = &userRepositoryImpl{}

func NewUserRepositoryImpl(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (u userRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	var user entity.User
	err := u.DB.Preload("Profile").Where(filter).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepositoryImpl) Find(filter interface{}) (interface{}, error) {
	var users entity.Users
	err := u.DB.Preload("Profile").Find(&users, filter).Error
	if err != nil {
		return nil, err
	}
	return users.UsersOutput(), nil
}

func (u userRepositoryImpl) Remove(data interface{}) error {
	return u.DB.Delete(data).Error
}

func (u userRepositoryImpl) Save(data interface{}) error {
	return u.DB.Create(data).Error
}

func (u userRepositoryImpl) Update(data interface{}) error {
	return u.DB.Update(data).Error
}
