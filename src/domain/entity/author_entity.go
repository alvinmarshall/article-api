package entity

import "github.com/google/uuid"

type Author struct {
	BaseEntity
	User      User
	UserID    uuid.UUID  `gorm:"type:uuid"`
	Articles  []Article  `gorm:"ForeignKey:AuthorID"`
	Favorites []Favorite `gorm:"ForeignKey:FavoriteByID"`
}
