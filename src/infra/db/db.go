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

type Repositories struct {
	User    repository.UserRepository
	Article repository.ArticleRepository
	DB      *gorm.DB
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

func NewRepositories() *Repositories {
	instance := GetDBsInstance()
	return &Repositories{
		User: repo_impl.NewUserRepositoryImpl(instance),
		DB:   instance,
		//Article: repo_impl.NewArticleRepositoryImpl(GetDBsInstance()),
	}
}

func (repos *Repositories) Close() error {
	return repos.DB.Close()
}

func (repos *Repositories) AutoMigrate() error {
	return repos.DB.AutoMigrate(
		&entity.User{},
		&entity.Profile{},
		//&entity.Article{},
		//&entity.Author{},
		//&entity.Comment{},
		//&entity.Follow{},
		//&entity.Favorite{},
		//&entity.Tag{},
	).Error
}
