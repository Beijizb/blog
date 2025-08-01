package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	PasswordHash string
	CreatedAt    time.Time
	Posts        []Post
}
