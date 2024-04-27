package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}
