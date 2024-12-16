package event_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/event"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) GetAll() ([]entities.Event, error) {
	args := m.Called()
	return args.Get(0).([]entities.Event), args.Error(1)
}

func (m *MockEventRepository) GetByID(id uint) (*entities.Event, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Event), args.Error(1)
}

func (m *MockEventRepository) GetbyBestPrice() ([]entities.Event, error) {
	args := m.Called()
	return args.Get(0).([]entities.Event), args.Error(1)
}

func (m *MockEventRepository) Create(event *entities.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventRepository) Update(event *entities.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	events := []entities.Event{
		{ID: 1, Name: "Event 1"},
		{ID: 2, Name: "Event 2"},
	}

	mockRepo.On("GetAll").Return(events, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, events, result)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	event := &entities.Event{ID: 1, Name: "Event 1"}

	mockRepo.On("GetByID", uint(1)).Return(event, nil)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, event, result)
	mockRepo.AssertExpectations(t)
}

func TestGetbyBestPrice(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	events := []entities.Event{
		{ID: 1, Name: "Event 1", Price: 100},
		{ID: 2, Name: "Event 2", Price: 50},
	}

	mockRepo.On("GetbyBestPrice").Return(events, nil)

	result, err := service.GetbyBestPrice()

	assert.NoError(t, err)
	assert.Equal(t, events, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	event := &entities.Event{Name: "New Event"}

	mockRepo.On("Create", event).Return(nil)

	err := service.Create(event)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	event := &entities.Event{ID: 1, Name: "Updated Event"}

	mockRepo.On("Update", event).Return(nil)

	err := service.Update(event)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Error(t *testing.T) {
	mockRepo := new(MockEventRepository)
	service := event.NewEventService(mockRepo)

	mockRepo.On("GetByID", uint(1)).Return((*entities.Event)(nil), errors.New("Event not found"))

	result, err := service.GetByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}