package cart_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/cart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCartRepository untuk CartRepositoryInterface
type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) FindByUserID(userID int) (entities.Cart, error) {
	args := m.Called(userID)
	return args.Get(0).(entities.Cart), args.Error(1)
}

func (m *MockCartRepository) FindByID(cartID int) (entities.Cart, error) {
	args := m.Called(cartID)
	return args.Get(0).(entities.Cart), args.Error(1)
}

func (m *MockCartRepository) Create(cart entities.Cart) (entities.Cart, error) {
	args := m.Called(cart)
	return args.Get(0).(entities.Cart), args.Error(1)
}

func (m *MockCartRepository) AddItemToCart(userID int, cartItem entities.CartItem) (entities.CartItem, error) {
	args := m.Called(userID, cartItem)
	return args.Get(0).(entities.CartItem), args.Error(1)
}

func (m *MockCartRepository) UpdateItemQuantity(cartItemID int, quantity int) (entities.CartItem, error) {
	args := m.Called(cartItemID, quantity)
	return args.Get(0).(entities.CartItem), args.Error(1)
}

func (m *MockCartRepository) RemoveItemFromCart(cartItemID int) error {
	args := m.Called(cartItemID)
	return args.Error(0)
}

func (m *MockCartRepository) ClearCart(cartID int) error {
	args := m.Called(cartID)
	return args.Error(0)
}

// Test untuk FindByUserID
func TestFindByUserID_Success(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	expectedCart := entities.Cart{
		ID:     1,
		UserID: 1,
		Items:  []entities.CartItem{},
	}

	mockRepo.On("FindByUserID", 1).Return(expectedCart, nil)

	result, err := cartService.FindByUserID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCart, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByUserID_Error(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("FindByUserID", 2).Return(entities.Cart{}, errors.New("user not found"))

	result, err := cartService.FindByUserID(2)

	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
	assert.Equal(t, entities.Cart{}, result)
	mockRepo.AssertExpectations(t)
}

// Test untuk AddItemToCart
func TestAddItemToCart_Success(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	cartItem := entities.CartItem{
		ID:       1,
		ProductID: 1,
		Quantity:  2,
	}

	mockRepo.On("AddItemToCart", 1, cartItem).Return(cartItem, nil)

	result, err := cartService.AddItemToCart(1, cartItem)

	assert.NoError(t, err)
	assert.Equal(t, cartItem, result)
	mockRepo.AssertExpectations(t)
}

func TestAddItemToCart_Error(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	cartItem := entities.CartItem{
		ID:       1,
		ProductID: 1,
		Quantity:  2,
	}

	mockRepo.On("AddItemToCart", 1, cartItem).Return(entities.CartItem{}, errors.New("failed to add item"))

	result, err := cartService.AddItemToCart(1, cartItem)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to add item")
	assert.Equal(t, entities.CartItem{}, result)
	mockRepo.AssertExpectations(t)
}

// Test untuk UpdateItemQuantity
func TestUpdateItemQuantity_Success(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	updatedItem := entities.CartItem{
		ID:       1,
		ProductID: 1,
		Quantity:  3,
	}

	mockRepo.On("UpdateItemQuantity", 1, 3).Return(updatedItem, nil)

	result, err := cartService.UpdateItemQuantity(1, 3)

	assert.NoError(t, err)
	assert.Equal(t, updatedItem, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItemQuantity_Error(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("UpdateItemQuantity", 1, 3).Return(entities.CartItem{}, errors.New("failed to update quantity"))

	result, err := cartService.UpdateItemQuantity(1, 3)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to update quantity")
	assert.Equal(t, entities.CartItem{}, result)
	mockRepo.AssertExpectations(t)
}

// Test untuk RemoveItemFromCart
func TestRemoveItemFromCart_Success(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("RemoveItemFromCart", 1).Return(nil)

	err := cartService.RemoveItemFromCart(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemoveItemFromCart_Error(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("RemoveItemFromCart", 1).Return(errors.New("failed to remove item"))

	err := cartService.RemoveItemFromCart(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to remove item")
	mockRepo.AssertExpectations(t)
}

// Test untuk ClearCart
func TestClearCart_Success(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("ClearCart", 1).Return(nil)

	err := cartService.ClearCart(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestClearCart_Error(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartService := cart.NewCartService(mockRepo)

	mockRepo.On("ClearCart", 1).Return(errors.New("failed to clear cart"))

	err := cartService.ClearCart(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to clear cart")
	mockRepo.AssertExpectations(t)
}
