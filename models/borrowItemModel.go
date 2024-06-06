package models

import "time"

type BorrowItem struct {
	BaseModel         `gorm:"embedded"`
	UserID            uint               `gorm:"not null" json:"userId"`
	BorrowerID        uint               `gorm:"not null" json:"borrowerId"`
	Borrower          *Borrower          `gorm:"foreignKey:BorrowerID;onDelete:CASCADE" json:"borrower"`
	User              *User              `gorm:"foreignKey:UserID;onDelete:CASCADE" json:"user"`
	Status            string             `gorm:"not null" json:"status"`
	ReturnedCondition string             `gorm:"not null"  json:"returnedCondition"`
	IsReturnedLate    bool               `gorm:"not null" json:"isReturnedLate" `
	Items             []ItemToBorrowItem `gorm:"foreignKey:BorrowItemID;references:ID" json:"items"`
	Description       string             `json:"description"`
	ReturnedDate      *time.Time         `json:"returnedDate"`
	ReturnBefore      *time.Time         `json:"returnBefore"`
}
