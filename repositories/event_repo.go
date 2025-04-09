package repositories

import (
	"StreefySherehes/models"
	"StreefySherehes/databases"

	"gorm.io/gorm"
)
type EventRepository struct{
	DB*gorm.DB
}
func NewRepository() EventRepository {
	return EventRepository{
		DB: databases.DB,
	}
}
func (repo *EventRepository) CreateEvent(event models.Event) (models.Event, error) {
	if result := repo.DB.Create(&event); result.Error != nil {
		return models.Event{}, result.Error
	}
	return event, nil
}

func (repo *EventRepository) GetEventByID(id uint) (models.Event, error) {
	var event models.Event
	if result := repo.DB.First(&event, id); result.Error != nil {
		return models.Event{}, result.Error
	}
	return event, nil
}
