package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"not null"`
	Role     Role   `gorm:"foreignKey:RoleID;onDelete:CASCADE"`
}

type UserWithoutPassword struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"not null"`
}
