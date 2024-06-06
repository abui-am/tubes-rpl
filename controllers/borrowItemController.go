package controllers

import (
	"net/http"
	"time"

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
		Description  string `json:"description"`
		ReturnBefore string `json:"returnBefore"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Convert the returnedDate string to time.Time
	time, err := time.Parse(time.RFC3339, body.ReturnBefore)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse returnedDate"})
		return
	}

	// Create a new BorrowItem
	borrowItem := models.BorrowItem{
		UserID:            body.UserID,
		BorrowerID:        body.BorrowerID,
		Status:            "borrowed",
		ReturnedCondition: "",
		IsReturnedLate:    false,
		Description:       body.Description,
		// body.returnedDate is a string, we need to convert it to time.Time
		ReturnBefore: &time,
		ReturnedDate: nil,
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

		// TODO: Handle the case when the item quantity is not enough
		// Check if the item quantity is enough

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
	result := initializers.DB.Debug().Preload(clause.Associations).Preload("User.Role").Preload("Items.Item").Order("created_at desc").Find(&borrowItems)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowItems"})
		return
	}

	// TODO : Implement search query
	// Handle search query
	// Filter the borrowItems based on the search query

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

func UpdateBorrowItem(c *gin.Context) {
	var body struct {
		Status            string `json:"status"`
		ReturnedCondition string `json:"returnedCondition"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var borrowItem models.BorrowItem
	result := initializers.DB.Find(&borrowItem, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BorrowItem not found"})
		return
	}

	// Update the borrowItem
	borrowItem.Status = body.Status
	borrowItem.ReturnedCondition = body.ReturnedCondition

	// Update the returnedDate if the status is returned

	time := time.Now()
	if body.Status == "returned" {
		borrowItem.ReturnedDate = &time
	}

	result = initializers.DB.Save(&borrowItem)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update borrowItem"})
		return
	}

	c.JSON(http.StatusOK, borrowItem)
}
