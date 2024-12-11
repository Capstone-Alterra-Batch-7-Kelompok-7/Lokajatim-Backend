package request

import (
	"lokajatim/constant"
	"lokajatim/entities"
)

type PaymentRequest struct {
	TransactionID string  `json:"transaction_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
}

func (paymentRequest PaymentRequest) ToEntities() (entities.Payment, error) {
	payment := entities.Payment{
		TransactionID: paymentRequest.TransactionID,
		PaymentMethod: paymentRequest.PaymentMethod,
		PaymentStatus: constant.PaymentStatusPending,
	}
	return payment, nil
}
