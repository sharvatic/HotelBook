package models

type User struct {
    ID           uint   `gorm:"primaryKey"`
    Username     string `json:"username" gorm:"unique;not null"`
    Password     string `json:"password" gorm:"-"`     // Temporary field for input
    HashPassword string `gorm:"column:password_hash"`  // Stored in the database
    Role         string `gorm:"type:varchar(20);default:'user'"`
}

