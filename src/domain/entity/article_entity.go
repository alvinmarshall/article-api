package entity

import "github.com/google/uuid"

type Article struct {
	BaseEntity
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      Author
	AuthorID    uuid.UUID `gorm:"type:uuid"`
	Tags        string
	Comments    string
}
