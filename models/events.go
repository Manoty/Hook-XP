package models

import (
	"gorm.io/gorm"
	
)
type Event struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	TicketPrice float64 `json:"ticket_price"`

}
