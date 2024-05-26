package models

type BorrowItem struct {
	BaseModel
	UserID            uint      `gorm:"not null" json:"userId"`
	BorrowerID        uint      `gorm:"not null" json:"borrowerId"`
	Borrower          *Borrower `gorm:"foreignKey:BorrowerID;onDelete:CASCADE" json:"borrower"`
	User              *User     `gorm:"foreignKey:UserID;onDelete:CASCADE" json:"user"`
	Status            string    `gorm:"not null" json:"status"`
	ReturnedCondition string    `gorm:"not null"  json:"returnedCondition"`
	IsReturnedLate    bool      `gorm:"not null" json:"isReturnedLate" `
	Items             []Item    `gorm:"many2many:item_to_borrow_item;" json:"items"`
	ItemIDs           []uint    `gorm:"-"`
}
