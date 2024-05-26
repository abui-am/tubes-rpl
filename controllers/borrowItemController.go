package controllers

import (
	"net/http"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"gihub.com/abui-am/tubes-rpl/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateBorrowItem(c *gin.Context) {
	var body struct {
		UserID     uint `json:"userId"`
		BorrowerID uint `json:"borrowerId"`
		Items      []struct {
			ItemID   uint `json:"itemId"`
			Quantity int  `json:"quantity"`
		} `json:"items"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Create a new BorrowItem
	borrowItem := models.BorrowItem{
		UserID:            body.UserID,
		BorrowerID:        body.BorrowerID,
		Status:            "borrowed",
		ReturnedCondition: "",
		IsReturnedLate:    false,
	}

	result := initializers.DB.Create(&borrowItem)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create borrowItem"})
		return
	}

	// Create ItemToBorrowItem for each item
	for _, item := range body.Items {
		itemToBorrowItem := models.ItemToBorrowItem{
			ItemID:       item.ItemID,
			BorrowItemID: borrowItem.ID,
			Quantity:     item.Quantity,
		}

		result := initializers.DB.Debug().Create(&itemToBorrowItem)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itemToBorrowItem"})

			// Rollback the transaction
			initializers.DB.Debug().Delete(&borrowItem)
			return
		}

		// Update the quantity of the item
		result = initializers.DB.Debug().Model(&models.Item{}).Where("id = ?", item.ItemID).Update("quantity", gorm.Expr("quantity - ?", item.Quantity))

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item quantity"})

			// Rollback the transaction
			initializers.DB.Debug().Delete(&borrowItem)
			initializers.DB.Debug().Delete(&itemToBorrowItem)
			return
		}

	}

	// Respond with the created borrowItem
	c.JSON(http.StatusCreated, borrowItem)

}

func GetBorrowItems(c *gin.Context) {
	var borrowItems []models.BorrowItem
	// Preload the items
	result := initializers.DB.Debug().Preload(clause.Associations).Preload("User.Role").Preload("Items.Item").Find(&borrowItems)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowItems"})
		return
	}

	c.JSON(http.StatusOK, borrowItems)
}

func GetBorrowItem(c *gin.Context) {
	var borrowItem models.BorrowItem
	result := initializers.DB.Debug().Preload(clause.Associations).Preload("User.Role").Preload("Items.Item").Find(&borrowItem, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BorrowItem not found"})
		return
	}

	c.JSON(http.StatusOK, borrowItem)
}
