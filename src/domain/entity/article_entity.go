package entity

import uuid "github.com/satori/go.uuid"

type Article struct {
	BaseEntity
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string
	Body        string
	Author      Author
	// @one to one belongs to ^
	AuthorID uuid.UUID `sql:"type:uuid"`
	//@many to many
	Tags []Tag `gorm:"many2many:article_tags"`
	//@one to many (Comment.Article)
	Comments []Comment
}

type Author struct {
	BaseEntity
	User User
	// @one to one belongs to ^
	UserID   string    `sql:"type:uuid"`
	Articles []Article `gorm:"foreign_key:AuthorID"`
	//@one to many(Favorite.AuthorID)
	Favorites []Favorite `gorm:"foreign_key:FavoriteBy"`
}

type Comment struct {
	BaseEntity
	Body    string
	Article Article
	//@one to one belongs to ^
	ArticleID string `sql:"type:uuid"`
	Author    Author
	//@one to one belong to ^
	AuthorID uuid.UUID `sql:"type:uuid"`
}

type Favorite struct {
	BaseEntity
	Favorite Article
	//@one to one belongs to ^
	FavoriteID string `sql:"type:uuid"`
	FavoriteBy Author
	//@one to one belongs to ^
	FavoriteByID uuid.UUID `sql:"type:uuid"`
}

type Tag struct {
	BaseEntity
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags"`
}
