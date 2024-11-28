package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharvatic/BookMyHotel/database"
	"github.com/sharvatic/BookMyHotel/models"
)

// CreateMenu allows staff to create a new menu
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create menu"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Menu created successfully", "menu": menu})
}

// ViewAllMenus retrieves all menus
func ViewAllMenus(c *gin.Context) {
	var menus []models.Menu
	if err := database.DB.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve menus"})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func AddMenuItem(c *gin.Context) {
	var menuItem models.MenuItem
	if err := c.ShouldBindJSON(&menuItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&menuItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add menu item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Menu item added successfully", "menu_item": menuItem})
}

func ViewAllMenuItems(c *gin.Context) {
	menuID := c.Param("menu_id")

	var menuItems []models.MenuItem
	if err := database.DB.Where("menu_id = ?", menuID).Find(&menuItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve menu items"})
		return
	}

	c.JSON(http.StatusOK, menuItems)
}

func UpdateMenuItems(c *gin.Context) {
	menuID := c.Param("menu_id")
	menuItemID := c.Param("menu_item_id")

	var menuItem models.MenuItem
	if err := c.ShouldBindJSON(&menuItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Model(&menuItem).Where("id = ? AND menu_id = ?", menuItemID, menuID).Updates(menuItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update menu item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu item updated successfully"})
}