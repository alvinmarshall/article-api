package entity

import "github.com/google/uuid"

type Favorite struct {
	BaseEntity
	Favorite     Article
	FavoriteID   uuid.UUID `gorm:"type:uuid"`
	FavoriteBy   Author
	FavoriteByID uuid.UUID `gorm:"type:uuid"`
}
