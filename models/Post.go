package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       float64
	ImagePost   string
	UserID      uint
	Usuario     User `gorm:"foreignKey:UserID"`
}
