package models

import "gorm.io/gorm"

type Borrower struct {
	gorm.Model
	Name               string `gorm:"not null"`
	Status             string `gorm:"not null"`
	RegistrationNumber string `gorm:"not null"`
	PhoneNumber        string `gorm:"not null"`
}
