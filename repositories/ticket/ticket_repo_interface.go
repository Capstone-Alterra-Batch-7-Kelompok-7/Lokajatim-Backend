package ticket

import "lokajatim/entities"

type TicketRepository interface {
	GetAll() ([]entities.Ticket, error)
	GetByID(id uint) (entities.Ticket, error)
	Create(ticket entities.Ticket) (entities.Ticket, error)
	Update(ticket entities.Ticket) (entities.Ticket, error)
	Delete(id uint) error
}