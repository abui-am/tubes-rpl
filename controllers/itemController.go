package controllers

import (
	"net/http"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"gihub.com/abui-am/tubes-rpl/models"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	result := initializers.DB.Find(&items)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	var item models.Item
	result := initializers.DB.First(&item, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var body struct {
		Name     string `gorm:"not null"`
		Quantity int    `gorm:"not null"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	item := models.Item{Name: body.Name, Quantity: body.Quantity}
	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	var item models.Item
	result := initializers.DB.First(&item, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var body = models.Item{}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	result = initializers.DB.Model(&item).Updates(&body)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, item)
}
