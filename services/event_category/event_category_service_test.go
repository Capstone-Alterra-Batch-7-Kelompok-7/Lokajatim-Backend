package event_category_test

import (
	"errors"
	"testing"

	"lokajatim/entities"
	"lokajatim/services/event_category"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEventCategoryRepository untuk EventCategoryRepository
type MockEventCategoryRepository struct {
	mock.Mock
}

func (m *MockEventCategoryRepository) GetAll() ([]entities.EventCategory, error) {
	args := m.Called()
	return args.Get(0).([]entities.EventCategory), args.Error(1)
}

func (m *MockEventCategoryRepository) GetByID(id int) (entities.EventCategory, error) {
	args := m.Called(id)
	return args.Get(0).(entities.EventCategory), args.Error(1)
}

func (m *MockEventCategoryRepository) Create(category entities.EventCategory) (entities.EventCategory, error) {
	args := m.Called(category)
	return args.Get(0).(entities.EventCategory), args.Error(1)
}

func (m *MockEventCategoryRepository) Update(category entities.EventCategory) (entities.EventCategory, error) {
	args := m.Called(category)
	return args.Get(0).(entities.EventCategory), args.Error(1)
}

func (m *MockEventCategoryRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllEventCategories(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockData := []entities.EventCategory{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}

	mockRepo.On("GetAll").Return(mockData, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetEventCategoryByID(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockData := entities.EventCategory{ID: 1, Name: "Category 1"}

	mockRepo.On("GetByID", 1).Return(mockData, nil)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateEventCategory(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockData := entities.EventCategory{ID: 1, Name: "Category 1"}

	mockRepo.On("Create", mock.AnythingOfType("entities.EventCategory")).Return(mockData, nil)

	result, err := service.Create(mockData)

	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateEventCategory(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockExistingData := entities.EventCategory{ID: 1, Name: "Old Name"}
	mockUpdatedData := entities.EventCategory{ID: 1, Name: "New Name"}

	mockRepo.On("GetByID", 1).Return(mockExistingData, nil)
	mockRepo.On("Update", mock.AnythingOfType("entities.EventCategory")).Return(mockUpdatedData, nil)

	result, err := service.Update(mockUpdatedData)

	assert.NoError(t, err)
	assert.Equal(t, mockUpdatedData, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteEventCategory(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockRepo.On("GetByID", 1).Return(entities.EventCategory{ID: 1, Name: "Category 1"}, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteEventCategory_NotFound(t *testing.T) {
	mockRepo := new(MockEventCategoryRepository)
	service := event_category.NewEventCategoryService(mockRepo)

	mockRepo.On("GetByID", 1).Return(entities.EventCategory{}, errors.New("not found"))

	err := service.Delete(1)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	mockRepo.AssertExpectations(t)
}
