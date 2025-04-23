package models

type Ticket struct {
	EventID uint `json:"event_id"`
	UserID uint `json:"user_id"`
	Token string `json:"token"`
}