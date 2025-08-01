package models

import "time"

type Post struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Slug       string `gorm:"unique"`
	Content    string `gorm:"type:text"`
	Cover      string
	CategoryID uint
	UserID     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   Category
	User       User
}
