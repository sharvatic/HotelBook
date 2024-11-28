package controllers

import (
	"net/http"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/sharvatic/BookMyHotel/database"
	"github.com/sharvatic/BookMyHotel/models"
)

// PlaceOrder allows a user to place an order
func PlaceOrder(c *gin.Context) {
	var order models.Order
	var orderItems []models.OrderItem

	// Assume `getUserID` retrieves the user ID of the current user
	userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

	var uid uint
    if idFloat, ok := userID.(float64); ok {
        uid = uint(idFloat)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

	if err := c.ShouldBindJSON(&orderItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create new order
	order = models.Order{
		UserID:    uid,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	// Add items to the order
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}

	if err := database.DB.Create(&orderItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add items to order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully", "order": order, "items": orderItems})
}

func ViewAllOrders(c *gin.Context) {
	
	var orders []models.Order
	var allOrderItems []models.OrderItem

	// Retrieve all orders
	if err := database.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve orders"})
		return
	}

	// Retrieve all order items
	if err := database.DB.Find(&allOrderItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve order items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders, "order_items": allOrderItems})
}

func ViewMyOrders(c *gin.Context) {
    // Get user ID from context
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

    var uid uint
    if idFloat, ok := userID.(float64); ok {
        uid = uint(idFloat)
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    var orders []models.Order

    // Find all orders that belong to the user
    if err := database.DB.Where("user_id = ?", uid).Find(&orders).Error; err != nil {
        fmt.Println("Error retrieving orders:", err) // Log the error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve orders"})
        return
    }

    // Prepare a map to hold order items by order ID
    orderItemsMap := make(map[uint][]models.OrderItem)

    // Get order IDs
    orderIDs := make([]uint, len(orders))
    for i, order := range orders {
        orderIDs[i] = order.ID
    }

    // Retrieve all items for the user's orders in a single query
    var orderItems []models.OrderItem
    if err := database.DB.Where("order_id IN ?", orderIDs).Find(&orderItems).Error; err != nil {
        fmt.Println("Error retrieving order items:", err) // Log the error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve order items"})
        return
    }

    // Group order items by order ID
    for _, item := range orderItems {
        orderItemsMap[item.OrderID] = append(orderItemsMap[item.OrderID], item)
    }

    // Create the response structure
    response := []gin.H{}
    for _, order := range orders {
        response = append(response, gin.H{
            "order":  order,
            "items":  orderItemsMap[order.ID],
        })
    }

    c.JSON(http.StatusOK, gin.H{"orders": response})
}



