package models

type Borrower struct {
	BaseModel          `gorm:"embedded"`
	Name               string `gorm:"not null" json:"name"`
	Status             string `gorm:"not null" json:"status"`
	RegistrationNumber string `gorm:"not null" json:"registrationNumber"`
}
