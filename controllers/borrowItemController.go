package controllers

import (
	"net/http"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"gihub.com/abui-am/tubes-rpl/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateBorrowItem(c *gin.Context) {
	var body struct {
		UserID     uint   `json:"userId"`
		BorrowerID uint   `json:"borrowerId"`
		ItemIds    []uint `json:"itemIds"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Create a new BorrowItem
	borrowItem := models.BorrowItem{
		UserID:            body.UserID,
		BorrowerID:        body.BorrowerID,
		ItemIDs:           body.ItemIds,
		Status:            "borrowed",
		ReturnedCondition: "",
		IsReturnedLate:    false,
	}

	// get all items from itemIds
	var items []models.Item
	result :=
		initializers.DB.Find(&items, body.ItemIds)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	borrowItem.Items = items

	result = initializers.DB.Create(&borrowItem)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create borrowItem"})
		return
	}

	// Respond with the created borrowItem
	c.JSON(http.StatusCreated, borrowItem)

}

func GetBorrowItems(c *gin.Context) {
	var borrowItems []models.BorrowItem
	result := initializers.DB.Preload("Items").Find(&borrowItems)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowItems"})
		return
	}

	c.JSON(http.StatusOK, borrowItems)
}

func GetBorrowItem(c *gin.Context) {
	var borrowItem models.BorrowItem
	result := initializers.DB.Debug().Preload(clause.Associations).Preload("User.Role").First(&borrowItem, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BorrowItem not found"})
		return
	}

	c.JSON(http.StatusOK, borrowItem)
}
