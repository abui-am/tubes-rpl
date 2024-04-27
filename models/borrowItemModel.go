package models

import "gorm.io/gorm"

type BorrowItem struct {
	gorm.Model
	UserID            uint     `gorm:"not null"`
	BorrowerID        uint     `gorm:"not null"`
	Borrower          Borrower `gorm:"foreignKey:BorrowerID;onDelete:CASCADE"`
	User              User     `gorm:"foreignKey:UserID;onDelete:CASCADE"`
	Status            string   `gorm:"not null"`
	ReturnedCondition string   `gorm:"not null"`
	IsReturnedLate    bool     `gorm:"not null"`
	Item              []Item   `gorm:"many2many:item_to_borrow_item;"`
}
