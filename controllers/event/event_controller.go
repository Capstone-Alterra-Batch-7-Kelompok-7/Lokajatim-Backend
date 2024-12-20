package event

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/event/request"
	"lokajatim/controllers/event/response"
	"lokajatim/controllers/pagination"
	"lokajatim/services/event"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	service event.EventService
}

func NewEventController(service event.EventService) *EventController {
	return &EventController{service}
}

// @Summary Get All Events
// @Description Get a list of all events
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} base.BaseResponse{data=[]response.EventResponse}
// @Failure 400 {object} base.BaseResponse
// @Router /events [get]
func (c *EventController) GetAllEvents(ctx echo.Context) error {
	events, err := c.service.GetAll()
	if err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}

	// Map entities to EventResponse and return them
	var eventResponses []response.EventResponse
	for _, e := range events {
		eventResponses = append(eventResponses, response.EventFromEntities(e))
	}
	return pagination.SuccessPaginatedResponse(ctx, events, 1, 10, int64(len(eventResponses)))
}

// @Summary Get Event by ID
// @Description Get details of an event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} base.BaseResponse{data=response.EventResponse}
// @Failure 404 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /events/{id} [get]
func (c *EventController) GetEventByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	event, err := c.service.GetByID(uint(id))
	if err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}
	return base.SuccesResponse(ctx, response.EventFromEntities(*event))
}

// @Summary Get Event by Best Price
// @Description Get details of an event with the best price
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} base.BaseResponse{data=[]response.EventResponse}
// @Failure 400 {object} base.BaseResponse
// @Router /events/best [get]
func (c *EventController) GetByBestPrice(ctx echo.Context) error {
	events, err := c.service.GetbyBestPrice()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to fetch events with the best price",
		})
	}

	return base.SuccesResponse(ctx, events)
}

// @Summary Create Event
// @Description Create a new event
// @Tags Event
// @Accept json
// @Produce json
// @Param request body request.EventRequest true "Event data"
// @Success 200 {object} base.BaseResponse{data=response.EventResponse}
// @Failure 400 {object} base.BaseResponse
// @Router /events [post]
func (c *EventController) CreateEvent(ctx echo.Context) error {
	var eventRequest request.EventRequest
	if err := ctx.Bind(&eventRequest); err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}

	event := eventRequest.ToEntities()
	err := c.service.Create(&event)
	if err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}

	return base.SuccesResponse(ctx, response.EventFromEntities(event))
}

// @Summary Update Event
// @Description Update an existing event
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param request body request.EventRequest true "Updated event data"
// @Success 200 {object} base.BaseResponse{data=response.EventResponse}
// @Failure 400 {object} base.BaseResponse
// @Failure 404 {object} base.BaseResponse
// @Router /events/{id} [put]
func (c *EventController) UpdateEvent(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var eventRequest request.EventRequest
	if err := ctx.Bind(&eventRequest); err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}

	event := eventRequest.ToEntities()
	event.ID = uint(id)
	err := c.service.Update(&event)
	if err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}

	return base.SuccesResponse(ctx, response.EventFromEntities(event))
}

// @Summary Delete Event
// @Description Delete an event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} base.BaseResponse{data=string}
// @Failure 400 {object} base.BaseResponse
// @Failure 404 {object} base.BaseResponse
// @Router /events/{id} [delete]
func (c *EventController) DeleteEvent(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.Delete(uint(id))
	if err != nil {
		return base.ErrorResponse(ctx, err, nil)
	}
	return base.SuccesResponse(ctx, map[string]string{"message": "Event deleted"})
}
