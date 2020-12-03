package entity

type Tag struct {
	BaseEntity
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags"`
}
