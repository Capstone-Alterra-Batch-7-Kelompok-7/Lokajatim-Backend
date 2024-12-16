package category_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/category"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository untuk CategoryRepositoryInterface
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetCategories() ([]entities.Category, error) {
	args := m.Called()
	return args.Get(0).([]entities.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetCategoryByID(id int) (entities.Category, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoryRepository) CreateCategory(category entities.Category) (entities.Category, error) {
	args := m.Called(category)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoryRepository) UpdateCategory(id int, category entities.Category) (entities.Category, error) {
	args := m.Called(id, category)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoryRepository) DeleteCategory(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test GetCategories sukses
func TestGetCategories_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	expectedCategories := []entities.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}

	mockRepo.On("GetCategories").Return(expectedCategories, nil)

	result, err := categoryService.GetCategories()

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCategories error
func TestGetCategories_Error(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	mockRepo.On("GetCategories").Return(nil, errors.New("failed to fetch categories"))

	result, err := categoryService.GetCategories()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCategoryByID sukses
func TestGetCategoryByID_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	expectedCategory := entities.Category{ID: 1, Name: "Category 1"}

	mockRepo.On("GetCategoryByID", 1).Return(expectedCategory, nil)

	result, err := categoryService.GetCategoryByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCategoryByID error
func TestGetCategoryByID_Error(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	mockRepo.On("GetCategoryByID", 1).Return(entities.Category{}, errors.New("category not found"))

	result, err := categoryService.GetCategoryByID(1)

	assert.Error(t, err)
	assert.Equal(t, entities.Category{}, result)
	mockRepo.AssertExpectations(t)
}

// Test CreateCategory sukses
func TestCreateCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	newCategory := entities.Category{Name: "New Category"}
	createdCategory := entities.Category{ID: 1, Name: "New Category"}

	mockRepo.On("CreateCategory", newCategory).Return(createdCategory, nil)

	result, err := categoryService.CreateCategory(newCategory)

	assert.NoError(t, err)
	assert.Equal(t, createdCategory, result)
	mockRepo.AssertExpectations(t)
}

// Test CreateCategory error
func TestCreateCategory_Error(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	newCategory := entities.Category{Name: "New Category"}

	mockRepo.On("CreateCategory", newCategory).Return(entities.Category{}, errors.New("failed to create category"))

	result, err := categoryService.CreateCategory(newCategory)

	assert.Error(t, err)
	assert.Equal(t, entities.Category{}, result)
	mockRepo.AssertExpectations(t)
}

// Test UpdateCategory sukses
func TestUpdateCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	updatedData := entities.Category{Name: "Updated Category"}
	updatedCategory := entities.Category{ID: 1, Name: "Updated Category"}

	mockRepo.On("UpdateCategory", 1, updatedData).Return(updatedCategory, nil)

	result, err := categoryService.UpdateCategory(1, updatedData)

	assert.NoError(t, err)
	assert.Equal(t, updatedCategory, result)
	mockRepo.AssertExpectations(t)
}

// Test UpdateCategory error
func TestUpdateCategory_Error(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	updatedData := entities.Category{Name: "Updated Category"}

	mockRepo.On("UpdateCategory", 1, updatedData).Return(entities.Category{}, errors.New("failed to update category"))

	result, err := categoryService.UpdateCategory(1, updatedData)

	assert.Error(t, err)
	assert.Equal(t, entities.Category{}, result)
	mockRepo.AssertExpectations(t)
}

// Test DeleteCategory sukses
func TestDeleteCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	mockRepo.On("DeleteCategory", 1).Return(nil)

	err := categoryService.DeleteCategory(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test DeleteCategory error
func TestDeleteCategory_Error(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryService := category.NewCategoryService(mockRepo)

	mockRepo.On("DeleteCategory", 1).Return(errors.New("failed to delete category"))

	err := categoryService.DeleteCategory(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to delete category")
	mockRepo.AssertExpectations(t)
}
