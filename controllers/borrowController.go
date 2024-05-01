package controllers

import (
	"net/http"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"gihub.com/abui-am/tubes-rpl/models"
	"github.com/gin-gonic/gin"
)

func CreateBorrower(c *gin.Context) {
	var body struct {
		Name               string `gorm:"not null" json:"name"`
		Status             string `gorm:"not null" json:"status"`
		RegistrationNumber string `gorm:"not null" json:"registrationNumber"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	borrower := models.Borrower{
		Name:               body.Name,
		Status:             body.Status,
		RegistrationNumber: body.RegistrationNumber,
	}

	result := initializers.DB.Create(&borrower)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, borrower)
}

func GetBorrowers(c *gin.Context) {
	var borrowers []models.Borrower
	result := initializers.DB.Find(&borrowers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowers"})
		return
	}

	c.JSON(http.StatusOK, borrowers)
}

func GetBorrower(c *gin.Context) {
	var borrower models.Borrower
	result := initializers.DB.First(&borrower, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrower not found"})
		return
	}

	c.JSON(http.StatusOK, borrower)
}

func UpdateBorrower(c *gin.Context) {
	var borrower models.Borrower
	result := initializers.DB.First(&borrower, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrower not found"})
		return
	}

	var body = models.Borrower{}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	result = initializers.DB.Model(&borrower).Updates(&body)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update borrower"})
		return
	}

	c.JSON(http.StatusOK, borrower)
}

func DeleteBorrower(c *gin.Context) {
	var borrower models.Borrower
	result := initializers.DB.First(&borrower, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrower not found"})
		return
	}

	result = initializers.DB.Delete(&borrower)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete borrower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Borrower deleted successfully",
	})
}
