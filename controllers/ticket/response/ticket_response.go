package response

import (
	"lokajatim/controllers/auth/response"
	eventResponse "lokajatim/controllers/event/response"
	"lokajatim/entities"
)

// TicketResponse is the response for a ticket
// @Description TicketResponse is the response for a ticket
// @Param ID uint true "Ticket ID"
// @Param EventsID uint true "Event ID"
// @Param TotalAmount int true "Total amount for the ticket"
// @Param UsersID int true "User ID"
// @Param PaymentMethod string true "Payment method"
// @Param Status string true "Ticket status"
// @Param CreatedAt string true "Creation timestamp"
// @Param UpdatedAt string true "Last update timestamp"
type TicketResponse struct {
	ID            uint   `json:"id"`
	EventsID      eventResponse.EventResponse  `json:"events_id"`
	TotalAmount   int    `json:"total_amount"`
	User          response.RegisterResponse `json:"user"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func FromEntity(ticket entities.Ticket, user entities.User, event entities.Event) TicketResponse {
	return TicketResponse{
		ID:            ticket.ID,
		EventsID:      eventResponse.EventFromEntities(event),
		TotalAmount:   ticket.TotalAmount,
		User:          response.RegisterFromEntities(user),
		PaymentMethod: ticket.PaymentMethod,
		Status:        ticket.Status,
		CreatedAt:     ticket.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     ticket.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

