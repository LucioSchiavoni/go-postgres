package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"not null; unique_index"`
	Email    string `gorm:"not null; unique_index"`
	Image    string
	Address  string
	Password string `gorm:"not null; unique_index"`
	Rol      string `gorm:"default:usuario"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}
