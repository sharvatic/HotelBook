package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/sharvatic/BookMyHotel/database"
    "github.com/sharvatic/BookMyHotel/models"
)

// CreateTable allows staff to add new tables to the restaurant
func CreateTable(c *gin.Context) {
    var table models.Table
    if err := c.ShouldBindJSON(&table); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }



    // Create the table in the database
    if err := database.DB.Create(&table).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create table"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Table created successfully", "table": table})
}

// BookTable allows a user to book an available table
func BookTable(c *gin.Context) {
    tableID := c.Param("id")
    var table models.Table

    // Find the table by its ID
    if err := database.DB.First(&table, tableID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
        return
    }

    // Check if the table is already booked
    if table.IsBooked {
        c.JSON(http.StatusConflict, gin.H{"error": "Table is already booked"})
        return
    }

    // Retrieve the userID from the context set by the auth middleware
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

    // Safely cast userID to uint after confirming its type
    var uid uint
    if idFloat, ok := userID.(float64); ok {
        uid = uint(idFloat)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    // Update table booking details
    table.IsBooked = true
    table.BookedBy = uid
	currentTime := time.Now()
	table.BookingTime = &currentTime


    // Save the updated table status
    if err := database.DB.Save(&table).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not book table"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Table booked successfully", "table": table})
}

// CancelTable allows a user to cancel their booking
func CancelTable(c *gin.Context) {
    tableID := c.Param("id")
    var table models.Table

    // Find the table
    if err := database.DB.First(&table, tableID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
        return
    }

    // Ensure the user is the one who booked the table
    // Retrieve the userID from the context set by the auth middleware
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

    // Safely cast userID to uint after confirming its type
    var uid uint
    if idFloat, ok := userID.(float64); ok {
        uid = uint(idFloat)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    if table.BookedBy != uid {
        c.JSON(http.StatusForbidden, gin.H{"error": "You did not book this table"})
        return
    }

    // Cancel the booking
    table.IsBooked = false
    table.BookedBy = 0
    table.BookingTime = nil

    if err := database.DB.Save(&table).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel booking"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Booking canceled successfully", "table": table})
}

// GetAllTables retrieves all tables from the database.
func GetTables(c *gin.Context) {
    var tables []models.Table

    // Fetch all tables from the database
    if err := database.DB.Find(&tables).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tables"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tables": tables})
}
