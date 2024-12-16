package ticket_test

import (
	"lokajatim/entities"
	"lokajatim/services/ticket"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTicketRepository adalah mock dari TicketRepository
type MockTicketRepository struct {
	mock.Mock
}

func (m *MockTicketRepository) GetAll() ([]entities.Ticket, error) {
	args := m.Called()
	return args.Get(0).([]entities.Ticket), args.Error(1)
}

func (m *MockTicketRepository) GetByID(id uint) (entities.Ticket, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Ticket), args.Error(1)
}

func (m *MockTicketRepository) Create(ticket entities.Ticket) (entities.Ticket, error) {
	args := m.Called(ticket)
	return args.Get(0).(entities.Ticket), args.Error(1)
}

func (m *MockTicketRepository) Update(ticket entities.Ticket) (entities.Ticket, error) {
	args := m.Called(ticket)
	return args.Get(0).(entities.Ticket), args.Error(1)
}

func (m *MockTicketRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllTickets(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	ticketService := ticket.NewTicketService(mockRepo)

	expectedTickets := []entities.Ticket{
		{ID: 1},
		{ID: 2},
	}

	mockRepo.On("GetAll").Return(expectedTickets, nil)

	tickets, err := ticketService.GetAllTickets()

	assert.NoError(t, err)
	assert.Equal(t, expectedTickets, tickets)
	mockRepo.AssertExpectations(t)
}

func TestGetTicketByID(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	ticketService := ticket.NewTicketService(mockRepo)

	expectedTicket := entities.Ticket{ID: 1}

	mockRepo.On("GetByID", uint(1)).Return(expectedTicket, nil)

	ticket, err := ticketService.GetTicketByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTicket, ticket)
	mockRepo.AssertExpectations(t)
}

func TestCreateTicket(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	ticketService := ticket.NewTicketService(mockRepo)

	newTicket := entities.Ticket{ID:1}
	expectedTicket := entities.Ticket{ID: 1}

	mockRepo.On("Create", newTicket).Return(expectedTicket, nil)

	ticket, err := ticketService.CreateTicket(newTicket)

	assert.NoError(t, err)
	assert.Equal(t, expectedTicket, ticket)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTicket(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	ticketService := ticket.NewTicketService(mockRepo)

	updatedTicket := entities.Ticket{ID: 1}
	mockRepo.On("Update", updatedTicket).Return(updatedTicket, nil)

	ticket, err := ticketService.UpdateTicket(updatedTicket)

	assert.NoError(t, err)
	assert.Equal(t, updatedTicket, ticket)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTicket(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	ticketService := ticket.NewTicketService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := ticketService.DeleteTicket(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}