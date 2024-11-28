package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Status    string    `gorm:"default:'pending'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint `gorm:"not null" json:"order_id"`
	MenuItemID uint `gorm:"not null" json:"menu_item_id"`
	Quantity   int  `gorm:"not null" json:"quantity"`
}

