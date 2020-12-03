package entity

import "github.com/google/uuid"

type Comment struct {
	BaseEntity
	Article   Article
	ArticleID uuid.UUID `gorm:"type:uuid"`
	Author    Author
	AuthorID  uuid.UUID `gorm:"type:uuid"`
	Body      string    `gorm:"size:2048"`
}
