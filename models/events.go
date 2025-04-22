package models

import (
	"gorm.io/gorm"
	
)
type Event struct {
	gorm.Model
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	TicketPrice float64 `json:"ticket_price"`

}
