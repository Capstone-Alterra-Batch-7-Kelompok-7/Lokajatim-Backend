package transaction

import (
	"errors"
	"fmt"
	"lokajatim/entities"
	"lokajatim/repositories/cart"
	"lokajatim/repositories/transaction"
	"lokajatim/utils"
)

type TransactionService struct {
	transactionRepo transaction.TransactionRepositoryInterface
	cartRepo        cart.CartRepositoryInterface
}

func NewTransactionService(transactionRepo transaction.TransactionRepositoryInterface, cartRepo cart.CartRepositoryInterface) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		cartRepo:        cartRepo,
	}
}

func (s *TransactionService) CreateTransaction(userID, cartID int) (entities.Transaction, error) {
	if userID == 0 || cartID == 0 {
		return entities.Transaction{}, errors.New("user_id or cart_id is missing")
	}

	cart, err := s.cartRepo.FindByID(cartID)
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("failed to fetch cart: %w", err)
	}

	totalGrossAmount := 0.0
	for _, item := range cart.Items {
		totalGrossAmount += float64(item.Product.Price) * float64(item.Quantity)
	}

	transaction := entities.Transaction{
		UserID:        userID,
		CartID:        cartID,
		TotalPrice:    totalGrossAmount,
		Status:        "Pending",
		TransactionID: "ORDER-" + utils.GenerateUniqueID(),
	}

	paymentURL, err := utils.CreateTransaction(
		transaction.TransactionID,
		int64(totalGrossAmount),
		cart.User.Name,
		cart.User.Email,
		cart.User.PhoneNumber,
		cart.User.Address,
		cart.Items,
	)
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("failed to create payment URL: %w", err)
	}
	transaction.PaymentURL = paymentURL

	createdTransaction, err := s.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("failed to save transaction: %w", err)
	}

	return createdTransaction, nil
}

func (s *TransactionService) HandleMidtransNotification(orderID, status string) error {
	_, err := s.transactionRepo.GetTransactionByOrderID(orderID)
	if err != nil {
		return fmt.Errorf("transaction not found for order ID %s: %w", orderID, err)
	}

	var updatedStatus string
	switch status {
	case "capture", "settlement":
		updatedStatus = "complete"
	case "pending":
		updatedStatus = "pending"
	case "deny", "expire", "cancel":
		updatedStatus = "failed"
	default:
		updatedStatus = "unknown"
	}

	if err := s.transactionRepo.UpdateTransactionStatusByOrderID(orderID, updatedStatus); err != nil {
		return fmt.Errorf("failed to update transaction status: %w", err)
	}

	return nil
}

func (s *TransactionService) GetTransactionByID(transactionID int) (entities.Transaction, error) {
	transaction, err := s.transactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (s *TransactionService) GetAllTransactions() ([]entities.Transaction, error) {
	transactions, err := s.transactionRepo.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) UpdateTransaction(id int, updatedTransaction entities.Transaction) (entities.Transaction, error) {
	if id <= 0 {
		return entities.Transaction{}, errors.New("invalid transaction ID")
	}

	updates := map[string]interface{}{
		"user_id":     updatedTransaction.UserID,
		"cart_id":     updatedTransaction.CartID,
		"total_price": updatedTransaction.TotalPrice + 2000,
		"status":      updatedTransaction.Status,
		"payment_url": updatedTransaction.PaymentURL,
		"updated_at":  utils.GetCurrentTime(),
	}

	transaction, err := s.transactionRepo.UpdateTransaction(id, updates)
	if err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil
}

func (s *TransactionService) UpdateTransactionStatus(transactionID int, status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	return s.transactionRepo.UpdateTransactionStatus(transactionID, status)
}

func (s *TransactionService) DeleteTransaction(transactionID int) error {
	return s.transactionRepo.DeleteTransaction(transactionID)
}
