package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type followRepositoryImpl struct {
	DB *gorm.DB
}

func NewFollowRepositoryImpl(db *gorm.DB) *followRepositoryImpl {
	return &followRepositoryImpl{db}
}

var _ repository.FollowRepository = &followRepositoryImpl{}

func (f followRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	follow := &entity.Follow{}
	err := f.DB.Where(filter).Find(follow).Error
	if err != nil {
		return nil, err
	}
	return follow, nil
}

func (f followRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	follows := &entity.Follows{}
	err := f.DB.Find(follows, filter).Error
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (f followRepositoryImpl) Remove(data interface{}) error {
	return f.DB.Delete(&entity.Follow{}, data).Error
}

func (f followRepositoryImpl) Save(data interface{}) error {
	return f.DB.Save(data).Error
}

func (f followRepositoryImpl) Update(data interface{}) error {
	return f.DB.Update(data).Error
}
