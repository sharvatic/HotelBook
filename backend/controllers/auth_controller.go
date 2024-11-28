package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sharvatic/BookMyHotel/database" // Import the database package
	"github.com/sharvatic/BookMyHotel/models" // Import the models package
	"github.com/sharvatic/BookMyHotel/utils"  // Import the utils package
	"github.com/sharvatic/BookMyHotel/middleware"
	"net/http"
)

// Signup handles user registration
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password using the utility function
	hashedPassword, err := utils.HashPassword(user.Password) // Assuming PasswordHash holds the plain password
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.HashPassword = hashedPassword

	// Create user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login handles user authentication

func Login(c *gin.Context) {
	var user models.User
	var input models.User // This will hold the login credentials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user by username
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Verify the password using the utility function
	if !utils.CheckPasswordHash(input.Password, user.HashPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenString, err := middleware.CreateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": tokenString})
}


