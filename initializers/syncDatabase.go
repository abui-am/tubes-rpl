package initializers

import "gihub.com/abui-am/tubes-rpl/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
