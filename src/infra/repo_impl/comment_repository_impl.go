package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type commentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepositoryImpl(db *gorm.DB) *commentRepositoryImpl {
	return &commentRepositoryImpl{db}
}

var _ repository.CommentRepository = &commentRepositoryImpl{}

func (c commentRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	comment := &entity.Comment{}
	err := c.DB.Where(filter).Find(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c commentRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	comments := &entity.Comments{}
	err := c.DB.Find(comments, filter).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c commentRepositoryImpl) Remove(data interface{}) error {
	return c.DB.Delete(&entity.Comment{}, data).Error
}

func (c commentRepositoryImpl) Save(data interface{}) error {
	return c.DB.Create(data).Error
}

func (c commentRepositoryImpl) Update(data interface{}) error {
	return c.DB.Update(data).Error
}
