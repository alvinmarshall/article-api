package db

import (
	"article_api/src/domain/entity"
	"article_api/src/domain/repository"
	"article_api/src/infra/repo_impl"
	"article_api/src/utils/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

type entityManager struct {
	DB       *gorm.DB
	User     repository.UserRepository
	Article  repository.ArticleRepository
	Author   repository.AuthorRepository
	Comment  repository.CommentRepository
	Favorite repository.FavoriteRepository
	Tag      repository.TagRepository
	Follow   repository.FollowRepository
}

func GetDBsInstance() *gorm.DB {
	if conn != nil {
		return conn
	}
	cfg := config.GetDBConfig()
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable password=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Password)
	db, err := gorm.Open(cfg.Driver, connectionString)
	if err != nil {
		panic(err)
	}
	println("connection success")
	conn = db
	return conn
}

func NewEntityManager() *entityManager {
	instance := GetDBsInstance()
	return &entityManager{
		DB:       instance,
		User:     repo_impl.NewUserRepositoryImpl(instance),
		Article:  repo_impl.NewArticleRepositoryImpl(instance),
		Author:   repo_impl.NewAuthorRepositoryImpl(instance),
		Comment:  repo_impl.NewCommentRepositoryImpl(instance),
		Favorite: repo_impl.NewFavoriteRepositoryImpl(instance),
		Tag:      repo_impl.NewTagRepositoryImpl(instance),
		Follow:   repo_impl.NewFollowRepositoryImpl(instance),
	}
}

func (repos *entityManager) Close() error {
	return repos.DB.Close()
}

func (repos *entityManager) AutoMigrate() error {
	return repos.DB.AutoMigrate(
		&entity.User{},
		&entity.Profile{},
		&entity.Article{},
		&entity.Author{},
		&entity.Comment{},
		&entity.Follow{},
		&entity.Favorite{},
		&entity.Tag{},
	).Error
}
