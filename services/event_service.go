package services


import (
	"StreefySherehes/models"
	"StreefySherehes/repositories"

)

type EventRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	TicketPrice float64 `json:"ticket_price"`
}

type EventResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	TicketPrice float64 `json:"ticket_price"`
}
type EventService struct {
	repo repositories.EventRepository
}

func NewEventsService() *EventService {
	return &EventService{
		repo: repositories.NewEventRepository(),

	}
}
func (es *EventRequest) CreateEvent(request EventRequest) (EventResponse, error) {
	event := models.Event{
		Title: request.Title,
		Description: request.Description,
		Location: request.Location,
		Date: request.Date,
		TicketPrice: request.TicketPrice,
	}
	createdEvent, err := es.repo.CreateEvent(event)
	if err != nil {
		return EventResponse{}, err
	}
	return EventResponse{
		ID: createdEvent.ID,
		Title: createdEvent.Title,
		Description: createdEvent.Description,
		Location: createdEvent.Location,
		Date: createdEvent.Date,
		TicketPrice: createdEvent.TicketPrice,
	}, nil
}