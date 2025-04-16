package services

import (
	"StreefySherehes/models"
	"StreefySherehes/repositories"
	"fmt"
	"net/http"
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
func (es *EventService) CreateEvent(request EventRequest) (EventResponse, error) {
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

func (es *EventService) GetAllEvents() ([]EventResponse, error) {
	events, err := es.repo.GetAllEvents()
	if err != nil {
		return nil, err 

	}
	var eventResponse [] EventResponse
	for _, events := range events {
		eventResponse = append(eventResponse, EventResponse{
			ID: events.ID,
			Title: events.Title,
			Description: events.Description,
			Location: events.Location,
			Date: events.Date,
			TicketPrice: events.TicketPrice,

	})

	}
	return eventResponse, nil
}
func (es *EventService) GetEventByID(id uint) (EventResponse, error) {
	event, err := es.repo.GetEventByID(id)
	if err != nil {
		return EventResponse{}, err
	}
	return EventResponse{
		ID: event.ID,
		Title: event.Title,
		Description: event.Description,
		Location: event.Location,
		Date: event.Date,
		TicketPrice: event.TicketPrice,
	}, nil

}
// UpdateEvent updates an existing event
func (es *EventService) UpdateEvent(id uint, request EventRequest) (EventResponse, error) {
	// Check if the event exists
	event, err := es.repo.GetEventByID(id)
	if err != nil {
		return EventResponse{}, fmt.Errorf("event not found: %w", err)
	}
	event = models.Event{
		Title:       request.Title,
		Description: request.Description,
		Location:    request.Location,
		Date:        request.Date,
		TicketPrice: request.TicketPrice,
	}
	updatedEvent, err := es.repo.UpdateEvent(id, event)
	if err != nil {
		return EventResponse{}, err 
	}
	return EventResponse{
		ID:          updatedEvent.ID,
		Title:	   updatedEvent.Title,
		Description: updatedEvent.Description,
		Location:    updatedEvent.Location,
		Date:     updatedEvent.Date,
		TicketPrice: updatedEvent.TicketPrice,
	},nil
}

// DeleteEvent deletes an event by ID
func (es *EventService) DeleteEvent(id uint) error {
	event, err := es.repo.GetEventByID(id)
	if err != nil {
		return fmt.Errorf("event not found: %w", err)

	}
	err = es.repo.DeleteEvent(id)
	if err != nil {
		return fmt.Errorf("failed to delete event with id: %w",id, err)
	}

	return nil


}