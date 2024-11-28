package models

import "time"

// Table struct represents the tables in the restaurant
type Table struct {
    ID          uint       `gorm:"primaryKey" json:"id"`
    TableNumber int        `gorm:"unique;not null" json:"table_number"`
    Seats       int        `gorm:"not null" json:"seats"`
    IsBooked    bool       `gorm:"default:false" json:"is_booked"`
    BookedBy    uint       `json:"booked_by,omitempty"`
    BookingTime *time.Time `json:"booking_time,omitempty"`
}

