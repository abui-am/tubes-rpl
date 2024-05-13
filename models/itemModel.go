package models

type Item struct {
	BaseModel `gorm:"embedded"`
	Name     string `gorm:"not null" json:"name"`
	Quantity int    `gorm:"not null" json:"quantity"`
}
