package initializers

import "gihub.com/abui-am/tubes-rpl/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Borrower{})
	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.BorrowItem{})
}
