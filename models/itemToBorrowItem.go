package models

type ItemToBorrowItem struct {
	ItemID       uint        `gorm:"primaryKey"`
	BorrowItemID uint        `gorm:"primaryKey"`
	Item         Item        `gorm:"foreignKey:ItemID;references:ID" json:"item"`
	BorrowItem   *BorrowItem `gorm:"foreignKey:BorrowItemID;references:ID" json:"borrowItem"`
	Quantity     int         `gorm:"not null" json:"quantity"`
}
