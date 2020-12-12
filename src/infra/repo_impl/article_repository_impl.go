package repo_impl

import (
	"article_api/src/domain/entity"
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
	article := &entity.Article{}
	err := a.DB.Preload("Author").Where(filter).Find(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a articleRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	articles := &entity.Articles{}
	err := a.DB.Preload("Author").Find(articles, filter).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a articleRepositoryImpl) Remove(data interface{}) error {
	return a.DB.Delete(&entity.Article{}, data).Error
}

func (a articleRepositoryImpl) Save(data interface{}) error {
	return a.DB.Create(data).Error
}

func (a articleRepositoryImpl) Update(data interface{}) error {
	return a.DB.Update(data).Error
}
