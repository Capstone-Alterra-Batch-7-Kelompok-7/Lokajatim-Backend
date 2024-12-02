package ticket

import "lokajatim/entities"

type TicketInterface interface {
	GetAllTickets() ([]entities.Ticket, error)
	GetTicketByID(id uint) (entities.Ticket, error)
	CreateTicket(ticket entities.Ticket) (entities.Ticket, error)
	UpdateTicket(ticket entities.Ticket) (entities.Ticket, error)
	DeleteTicket(id uint) error
}