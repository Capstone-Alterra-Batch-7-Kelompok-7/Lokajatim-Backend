package ticket

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) GetAll() ([]entities.Ticket, error) {
	var tickets []entities.Ticket
	err := r.db.Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) GetByID(id uint) (entities.Ticket, error) {
	var ticket entities.Ticket
	err := r.db.First(&ticket, id).Error
	return ticket, err
}

func (r *ticketRepository) Create(ticket entities.Ticket) (entities.Ticket, error) {
	err := r.db.Create(&ticket).Error
	return ticket, err
}

func (r *ticketRepository) Update(ticket entities.Ticket) (entities.Ticket, error) {
	var existingTicket entities.Ticket
	err := r.db.First(&existingTicket, ticket.ID).Error
	if err != nil {
		return entities.Ticket{}, err
	}

	ticket.CreatedAt = existingTicket.CreatedAt

	err = r.db.Save(&ticket).Error
	return ticket, err
}

func (r *ticketRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.Ticket{}, id).Error
	return err
}
