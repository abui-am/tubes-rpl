package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserWithoutPassword struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
}