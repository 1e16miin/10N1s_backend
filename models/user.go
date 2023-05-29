package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"unique"`
	Username   string `gorm:"not null"`
	Password   string `gorm:"not null"`
	University string
}
