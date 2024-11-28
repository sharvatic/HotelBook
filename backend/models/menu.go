package models

type Menu struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"not null" json:"name"`
	Active bool   `gorm:"default:true"`
}


type MenuItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	MenuID      uint    `gorm:"not null" json:"menu_id"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `gorm:"not null" json:"price"`
	Available   bool    `gorm:"default:true"`	
}
