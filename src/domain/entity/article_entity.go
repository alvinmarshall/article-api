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

type Articles []Article

type Author struct {
	BaseEntity
	User User
	// @one to one belongs to ^
	UserID   uuid.UUID `sql:"type:uuid"`
	Articles []Article `gorm:"foreign_key:AuthorID"`
	//@one to many(Favorite.AuthorID)
	Favorites []Favorite `gorm:"foreign_key:FavoriteBy"`
}

type Authors []Author

type Comment struct {
	BaseEntity
	Body    string
	Article Article
	//@one to one belongs to ^
	ArticleID uuid.UUID `sql:"type:uuid"`
	Author    Author
	//@one to one belong to ^
	AuthorID uuid.UUID `sql:"type:uuid"`
}

type Comments []Comment

type Favorite struct {
	BaseEntity
	Favorite Article
	//@one to one belongs to ^
	FavoriteID uuid.UUID `sql:"type:uuid"`
	FavoriteBy Author
	//@one to one belongs to ^
	FavoriteByID uuid.UUID `sql:"type:uuid"`
}
type Favorites []Favorite

type Tag struct {
	BaseEntity
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags"`
}
type Tags []Tag
