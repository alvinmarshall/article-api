package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type authorRepositoryImpl struct {
	DB *gorm.DB
}

var _ repository.AuthorRepository = &authorRepositoryImpl{}

func NewAuthorRepositoryImpl(db *gorm.DB) *authorRepositoryImpl {
	return &authorRepositoryImpl{db}
}

func (a authorRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	author := &entity.Author{}
	err := a.DB.Preload("User").Where(filter).Find(author).Error
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (a authorRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	authors := &entity.Authors{}
	err := a.DB.Preload("User").Find(authors, filter).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a authorRepositoryImpl) Remove(data interface{}) error {
	return a.DB.Delete(&entity.Author{}, data).Error
}

func (a authorRepositoryImpl) Save(data interface{}) error {
	return a.DB.Create(data).Error
}

func (a authorRepositoryImpl) Update(data interface{}) error {
	return a.DB.Update(data).Error
}
