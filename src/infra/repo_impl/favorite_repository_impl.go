package repo_impl

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type favoriteRepositoryImpl struct {
	DB *gorm.DB
}

func NewFavoriteRepositoryImpl(db *gorm.DB) *favoriteRepositoryImpl {
	return &favoriteRepositoryImpl{db}
}

var _ repository.TagRepository = &favoriteRepositoryImpl{}

func (f favoriteRepositoryImpl) FindOne(filter interface{}) (interface{}, error) {
	favorite := &entity.Favorite{}
	err := f.DB.Where(filter).Find(favorite).Error
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func (f favoriteRepositoryImpl) Find(filter ...interface{}) (interface{}, error) {
	favorites := &entity.Favorites{}
	err := f.DB.Find(favorites, filter).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (f favoriteRepositoryImpl) Remove(data interface{}) error {
	return f.DB.Delete(&entity.Favorite{}, data).Error
}

func (f favoriteRepositoryImpl) Save(data interface{}) error {
	return f.DB.Save(&data).Error
}

func (f favoriteRepositoryImpl) Update(data interface{}) error {
	return f.DB.Update(data).Error
}
