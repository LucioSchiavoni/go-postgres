package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      uint
}
