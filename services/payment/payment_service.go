package payment

import (
	"errors"
	"log"
	"lokajatim/entities"
	"lokajatim/repositories/payment"
	"lokajatim/repositories/transaction"
	"lokajatim/utils"
	"strconv"
)

type PaymentService struct {
	paymentRepo     payment.PaymentRepoInterface
	transactionRepo transaction.TransactionRepositoryInterface
}

func NewPaymentService(paymentRepo payment.PaymentRepoInterface, transactionRepo transaction.TransactionRepositoryInterface) PaymentService {
	return PaymentService{
		paymentRepo:     paymentRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *PaymentService) ProcessMidtransNotification(notification utils.NotificationRequest) error {

	transactionID, err := strconv.Atoi(notification.TransactionID)
	if err != nil {
		log.Println("Failed to convert TransactionID to int:", err)
		return err
	}

	transaction, err := s.transactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		log.Println("Failed to get transaction:", err)
		return err
	}

	var newStatus string
	var paymentStatus string

	switch notification.TransactionStatus {
	case "settlement":
		newStatus = "completed"
		paymentStatus = "success"
	case "pending":
		newStatus = "pending"
		paymentStatus = "pending"
	case "deny", "expire", "cancel":
		newStatus = "failed"
		paymentStatus = "failed"
	default:
		return errors.New("unknown transaction status")
	}

	err = s.paymentRepo.UpdatePaymentStatus(transaction.ID, paymentStatus)
	if err != nil {
		log.Println("Failed to update payment status:", err)
		return err
	}

	err = s.transactionRepo.UpdateTransactionStatus(transaction.ID, newStatus)
	if err != nil {
		log.Println("Failed to update transaction status:", err)
		return err
	}

	return nil
}

func (s *PaymentService) CreatePayment(payment entities.Payment) (entities.Payment, error) {
	return s.paymentRepo.CreatePayment(payment)
}

func (s *PaymentService) GetPaymentByID(paymentID int) (entities.Payment, error) {
	return s.paymentRepo.GetPaymentByID(paymentID)
}

func (s *PaymentService) UpdatePaymentStatus(paymentID int, status string) error {
	return s.paymentRepo.UpdatePaymentStatus(paymentID, status)
}

func (s *PaymentService) DeletePayment(paymentID int) error {
	return s.paymentRepo.DeletePayment(paymentID)
}
