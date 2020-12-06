package repo_impl

import (
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type articleRepositoryImpl struct {
	DB *gorm.DB
}

var _ repository.ArticleRepository = &articleRepositoryImpl{}

func NewArticleRepositoryImpl(db *gorm.DB) *articleRepositoryImpl {
	return &articleRepositoryImpl{db}
}

func (a articleRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	panic("implement me")
}

func (a articleRepositoryImpl) Find(filter interface{}) (interface{}, error) {
	panic("implement me")
}

func (a articleRepositoryImpl) Remove(data interface{}) error {
	panic("implement me")
}

func (a articleRepositoryImpl) Save(data interface{}) error {
	panic("implement me")
}

func (a articleRepositoryImpl) Update(data interface{}) error {
	panic("implement me")
}
