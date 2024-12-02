package requests

import (
	"lokajatim/entities"
)

// CreateTicketRequest is the request for creating a ticket
// @Description CreateTicketRequest is the request for creating a ticket
// @Param EventsID uint true "Event ID"
// @Param TotalAmount int true "Total amount for the ticket"
// @Param UsersID uint true "User ID"
// @Param PaymentMethod string true "Payment method"
// @Param Status string true "Ticket status"
type CreateTicketRequest struct {
	EventsID      uint   `json:"events_id" validate:"required"`
	TotalAmount   int    `json:"total_amount" validate:"required,min=1"`
	UsersID       int   `json:"users_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	Status        string `json:"status" validate:"required,oneof='pending' 'paid' 'cancelled'"`
}

// UpdateTicketRequest is the request for updating a ticket
// @Description UpdateTicketRequest is the request for updating a ticket
// @Param EventsID uint false "Event ID"
// @Param TotalAmount int false "Total amount for the ticket"
// @Param UsersID uint false "User ID"
// @Param PaymentMethod string false "Payment method"
// @Param Status string false "Ticket status"
type UpdateTicketRequest struct {
	EventsID      uint   `json:"events_id" validate:"omitempty"`
	TotalAmount   int    `json:"total_amount" validate:"omitempty,min=1"`
	UsersID       int   `json:"users_id" validate:"omitempty"`
	PaymentMethod string `json:"payment_method" validate:"omitempty"`
	Status        string `json:"status" validate:"omitempty,oneof='pending' 'paid' 'cancelled'"`
}

func (req CreateTicketRequest) ToEntity() entities.Ticket {
	return entities.Ticket{
		EventsID:      req.EventsID,
		TotalAmount:   req.TotalAmount,
		UsersID:       req.UsersID,
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
	}
}

func (req UpdateTicketRequest) ToEntity() entities.Ticket {
	return entities.Ticket{
		EventsID:      req.EventsID,
		TotalAmount:   req.TotalAmount,
		UsersID:       req.UsersID,
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
	}
}
