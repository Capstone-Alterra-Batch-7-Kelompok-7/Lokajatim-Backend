package transaction_test

import (
	"errors"
	"testing"

	"lokajatim/entities"
	"lokajatim/services/transaction"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(t entities.Transaction) (entities.Transaction, error) {
	args := m.Called(t)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionByOrderID(orderID string) (entities.Transaction, error) {
	args := m.Called(orderID)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) UpdateTransactionStatusByOrderID(orderID string, status string) error {
	args := m.Called(orderID, status)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetTransactionByID(transactionID int) (entities.Transaction, error) {
	args := m.Called(transactionID)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetAllTransactions() ([]entities.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) UpdateTransaction(id int, updates map[string]interface{}) (entities.Transaction, error) {
	args := m.Called(id, updates)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) UpdateTransactionStatus(transactionID int, status string) error {
	args := m.Called(transactionID, status)
	return args.Error(0)
}

func (m *MockTransactionRepository) DeleteTransaction(transactionID int) error {
	args := m.Called(transactionID)
	return args.Error(0)
}

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

func (m *MockCartRepository) FindByCartItemID(cartItemID int) (entities.Cart, error) {
	args := m.Called(cartItemID)
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


func TestHandleMidtransNotification(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockCartRepo := new(MockCartRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockCartRepo)

	// Test successful update
	mockTransactionRepo.On("GetTransactionByOrderID", "ORDER-12345").Return(entities.Transaction{}, nil)
	mockTransactionRepo.On("UpdateTransactionStatusByOrderID", "ORDER-12345", "complete").Return(nil)

	err := transactionService.HandleMidtransNotification("ORDER-12345", "settlement")
	assert.Nil(t, err)

	// Test transaction not found
	mockTransactionRepo.On("GetTransactionByOrderID", "ORDER-99999").Return(entities.Transaction{}, errors.New("transaction not found"))

	err = transactionService.HandleMidtransNotification("ORDER-99999", "capture")
	assert.NotNil(t, err)
	assert.Equal(t, "transaction not found for order ID ORDER-99999: transaction not found", err.Error())
}

func TestUpdateTransaction(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockCartRepo := new(MockCartRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockCartRepo)

	// Test valid case
	updatedTransaction := entities.Transaction{
		UserID:     1,
		CartID:     1,
		TotalPrice: 2200,
		Status:     "Pending",
		PaymentURL: "http://example.com/payment2",
	}
	mockTransactionRepo.On("UpdateTransaction", 1, mock.Anything).Return(updatedTransaction, nil)

	transaction, err := transactionService.UpdateTransaction(1, updatedTransaction)
	assert.Nil(t, err)
	assert.Equal(t, 2200.0, transaction.TotalPrice)

	// Test invalid ID
	transaction, err = transactionService.UpdateTransaction(0, updatedTransaction)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid transaction ID", err.Error())
}

func TestUpdateTransactionStatus(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockCartRepo := new(MockCartRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockCartRepo)

	// Test valid case
	mockTransactionRepo.On("UpdateTransactionStatus", 1, "complete").Return(nil)

	err := transactionService.UpdateTransactionStatus(1, "complete")
	assert.Nil(t, err)

	// Test invalid status
	err = transactionService.UpdateTransactionStatus(1, "")
	assert.NotNil(t, err)
	assert.Equal(t, "status cannot be empty", err.Error())
}

func TestDeleteTransaction(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockCartRepo := new(MockCartRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockCartRepo)

	// Test successful deletion
	mockTransactionRepo.On("DeleteTransaction", 1).Return(nil)

	err := transactionService.DeleteTransaction(1)
	assert.Nil(t, err)

	// Test deletion failure
	mockTransactionRepo.On("DeleteTransaction", 999).Return(errors.New("failed to delete transaction"))

	err = transactionService.DeleteTransaction(999)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to delete transaction", err.Error())
}
