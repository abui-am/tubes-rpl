package controllers

import (
	"net/http"
	"os"
	"time"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"gihub.com/abui-am/tubes-rpl/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get email and password from request
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   uint   `json:"roleId"`
		Name     string `json:"name"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return

	}

	// Create a new user
	user := models.User{Email: body.Email, Password: string(hash),
		RoleID: body.RoleID,
		Name:   body.Name,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, user)

}

func Login(c *gin.Context) {
	// Get email and password from request
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Find the user and preload the role
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).Preload("Role").First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	// Set the token as a cookie
	c.SetCookie("Authorization", tokenString, int(time.Hour*24*30), "/", "", false, true)

	// Respond with the user
	c.JSON(http.StatusOK, gin.H{
		"user": models.UserWithoutPassword{
			Email:     user.Email,
			RoleID:    user.RoleID,
			Role:      user.Role,
			BaseModel: user.BaseModel,
			Name:      user.Name,
		},
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Valid token"})
}

func GetUsers(c *gin.Context) {

	// Search for users
	var searchQuery = c.Query("search")

	var users []models.User
	var result = initializers.DB
	if searchQuery != "" {
		result = result.Debug().Where("name LIKE ?", "%"+searchQuery+"%").Preload("Role").Find(&users)
	} else {
		result = result.Preload("Role").Find(&users)
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Respond with the users
	var usersWithoutPassword []models.UserWithoutPassword

	for _, user := range users {
		usersWithoutPassword = append(usersWithoutPassword, models.UserWithoutPassword{
			BaseModel: user.BaseModel,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			Role:      user.Role,
		})
	}

	// If no users are found, return an empty array
	if len(usersWithoutPassword) == 0 {
		usersWithoutPassword = []models.UserWithoutPassword{}
	}

	c.JSON(http.StatusOK, usersWithoutPassword)
}

func GetUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.Preload("Role").First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, models.UserWithoutPassword{
		BaseModel: user.BaseModel,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		Role:      user.Role,
	})
}

func UpdateUser(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		RoleID   uint   `json:"roleId"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Email = body.Email
	user.RoleID = body.RoleID
	user.Name = body.Name
	user.Password = string(hash)

	result = initializers.DB.Save(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	initializers.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
