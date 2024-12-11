package response

import (
	"lokajatim/entities"
	"time"
)

type PaymentResponse struct {
	ID            int                  `json:"id"`
	TransactionID string               `json:"transaction_id"`
	Transaction   entities.Transaction `json:"transaction"`
	PaymentMethod string               `json:"payment_method"`
	PaymentStatus string               `json:"payment_status"`
	PaymentURL    string               `json:"payment_url"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

func PaymentFromEntity(payment entities.Payment) PaymentResponse {
	return PaymentResponse{
		ID:            payment.ID,
		TransactionID: payment.TransactionID,
		Transaction:   payment.Transaction,
		PaymentMethod: payment.PaymentMethod,
		PaymentStatus: payment.PaymentStatus,
		PaymentURL:    payment.PaymentURL,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
}
