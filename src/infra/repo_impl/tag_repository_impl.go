package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type tagRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagRepositoryImpl(db *gorm.DB) *tagRepositoryImpl {
	return &tagRepositoryImpl{db}
}

var _ repository.TagRepository = &tagRepositoryImpl{}

func (t tagRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	tag := &entity.Tag{}
	err := t.DB.Where(filter).Find(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (t tagRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	tags := &entity.Tags{}
	err := t.DB.Preload("Articles").Find(filter).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t tagRepositoryImpl) Remove(data interface{}) error {
	return t.DB.Delete(&entity.Tag{}, data).Error
}

func (t tagRepositoryImpl) Save(data interface{}) error {
	return t.DB.Save(data).Error
}

func (t tagRepositoryImpl) Update(data interface{}) error {
	return t.DB.Update(data).Error
}
