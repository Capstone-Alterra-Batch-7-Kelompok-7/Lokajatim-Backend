package ticket

import (
	"lokajatim/entities"
	"lokajatim/repositories/ticket"
)

type ticketService struct {
	ticketRepo ticket.TicketRepository
}

func NewTicketService(repo ticket.TicketRepository) TicketInterface {
	return &ticketService{ticketRepo: repo}
}

func (s *ticketService) GetAllTickets() ([]entities.Ticket, error) {
	return s.ticketRepo.GetAll()
}

func (s *ticketService) GetTicketByID(id uint) (entities.Ticket, error) {
	return s.ticketRepo.GetByID(id)
}

func (s *ticketService) CreateTicket(ticket entities.Ticket) (entities.Ticket, error) {
	return s.ticketRepo.Create(ticket)
}

func (s *ticketService) UpdateTicket(ticket entities.Ticket) (entities.Ticket, error) {
	return s.ticketRepo.Update(ticket)
}

func (s *ticketService) DeleteTicket(id uint) error {
	return s.ticketRepo.Delete(id)
}
