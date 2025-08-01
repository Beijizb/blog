package models

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Slug      string `gorm:"unique"`
	CreatedAt time.Time
	Posts     []Post
}
