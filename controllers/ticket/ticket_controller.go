package ticket

import (
	"fmt"
	"lokajatim/controllers/base"
	"lokajatim/controllers/pagination"
	"lokajatim/controllers/ticket/response"
	"lokajatim/entities"
	services "lokajatim/services/auth"
	eventService "lokajatim/services/event"
	"lokajatim/services/ticket"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TicketController struct {
	ticketService ticket.TicketInterface
	authService   services.AuthServiceInterface
	eventService  eventService.EventService
}

func NewTicketController(ticketService ticket.TicketInterface, authService services.AuthServiceInterface,eventService eventService.EventService) *TicketController {
	return &TicketController{
		ticketService: ticketService,
		authService:   authService,
		eventService: eventService,
	}
}

// @Summary Get All Tickets
// @Tags Ticket
// @Produce json
// @Success 200 {object} base.BaseResponse{data=[]entities.Ticket}
// @Router /tickets [get]
func (tc *TicketController) GetAllTickets(c echo.Context) error {
	tickets, err := tc.ticketService.GetAllTickets()
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		user, err := tc.authService.GetUserByID(ticket.UsersID)
		if err != nil {
			return base.ErrorResponse(c, fmt.Errorf("user with ID %d not found", ticket.UsersID), nil)
		}

		event, err := tc.eventService.GetByID(ticket.EventsID)
		if err != nil {
			return base.ErrorResponse(c, fmt.Errorf("event with ID %d not found", ticket.EventsID), nil)
		}
		ticketResponse := response.FromEntity(ticket, user, *event)
		ticketResponses = append(ticketResponses, ticketResponse)
	}

	return pagination.SuccessPaginatedResponse(c, ticketResponses, 1, 10, int64(len(ticketResponses)))
}

// @Summary Get Ticket By ID
// @Tags Ticket
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} base.BaseResponse{data=entities.Ticket}
// @Router /tickets/{id} [get]
func (tc *TicketController) GetTicketByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ticket, err := tc.ticketService.GetTicketByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}
	user, err := tc.authService.GetUserByID(ticket.UsersID)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	event, err := tc.eventService.GetByID(ticket.EventsID)
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("event with ID %d not found", ticket.EventsID), nil)
	}

	ticketResponse := response.FromEntity(ticket, user, *event)
	return base.SuccesResponse(c, ticketResponse)
}

// @Summary Create Ticket
// @Tags Ticket
// @Accept json
// @Produce json
// @Param request body entities.Ticket true "Create Ticket"
// @Success 201 {object} base.BaseResponse{data=entities.Ticket}
// @Router /tickets [post]
func (tc *TicketController) CreateTicket(c echo.Context) error {
	var ticket entities.Ticket
	if err := c.Bind(&ticket); err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	user, err := tc.authService.GetUserByID(ticket.UsersID)
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("user not found"), nil)
	}

	createdTicket, err := tc.ticketService.CreateTicket(ticket)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	event, err := tc.eventService.GetByID(createdTicket.EventsID)
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("event with ID %d not found", createdTicket.EventsID), nil)
	}

	return base.SuccesResponse(c, response.FromEntity(createdTicket, user, *event))
}

// @Summary Update Ticket
// @Tags Ticket
// @Accept json
// @Produce json
// @Param id path int true "Ticket ID"
// @Param request body entities.Ticket true "Update Ticket"
// @Success 200 {object} base.BaseResponse{data=entities.Ticket}
// @Router /tickets/{id} [put]
func (tc *TicketController) UpdateTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var ticket entities.Ticket
	if err := c.Bind(&ticket); err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	ticket.ID = uint(id)

	updatedTicket, err := tc.ticketService.UpdateTicket(ticket)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	user, err := tc.authService.GetUserByID(updatedTicket.UsersID)
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("user not found"), nil)
	}
	event, err := tc.eventService.GetByID(updatedTicket.EventsID)
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("event with ID %d not found", updatedTicket.EventsID), nil)
	}
	return base.SuccesResponse(c, response.FromEntity(updatedTicket, user, *event))
}

// @Summary Delete Ticket
// @Tags Ticket
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 204 {object} base.BaseResponse
// @Router /tickets/{id} [delete]
func (tc *TicketController) DeleteTicket(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    _, err := tc.ticketService.GetTicketByID(uint(id))
    if err != nil {
        return base.ErrorResponse(c, fmt.Errorf("ticket with ID %d not found", id), nil)
    }

    err = tc.ticketService.DeleteTicket(uint(id))
    if err != nil {
        return base.ErrorResponse(c, err, nil)
    }

    return base.SuccesResponse(c, map[string]string{"message": "Ticket deleted"})
}
