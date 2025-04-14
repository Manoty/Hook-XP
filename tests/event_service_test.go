package tests 

import (
	"testing"
	"StreefySherehes/models"
	"StreefySherehes/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEventRepository is a mock implementation of the EventRepository interface
type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) Create(event *models.Event) error {
	args := m.Called(event)
	return args.Error(0)
}
func TestCreateEventService(t *testing.T) {
	// Create a new instance of the EventService with the mock repository
	mockRepo := new(MockEventRepository)
	service := services.NewEventsService()

	event := &models.Event {
		Title: "Test Event",
		Description: "This is a test event",
	}

	mockRepo.On("Create", event).Return(nil)

	err := service.CreateEvent(event)
	assert.Nil(t, err) // Assert that no error was returned

	mockRepo.AssertExpectations(t) // Assert that the mock method was called as expected
}