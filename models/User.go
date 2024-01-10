package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"not null; unique_index"`
	Email    string `gorm:"not null; unique_index"`
	Image    string
	Password string
	Posts    []Post
}
