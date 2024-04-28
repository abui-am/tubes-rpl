package seeder

import (
	"fmt"

	"gihub.com/abui-am/tubes-rpl/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{BaseModel: models.BaseModel{
			ID: 1,
		}, Name: "superadmin"},
		{BaseModel: models.BaseModel{
			ID: 2,
		}, Name: "admin"},
	}

	for _, role := range roles {
		err := db.Save(&role).Error
		if err != nil {
			fmt.Printf("Error when create roles: %s\n", role.Name)
		} else {
			fmt.Printf("Success create roles: %s\n", role.Name)
		}

	}
}
