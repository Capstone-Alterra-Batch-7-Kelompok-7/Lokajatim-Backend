package event_category

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/event_category/request"
	"lokajatim/controllers/event_category/response"
	"lokajatim/controllers/pagination"
	"lokajatim/services/event_category"

	"strconv"

	"github.com/labstack/echo/v4"
)

type EventCategoryController struct {
	EventCategoryService event_category.EventCategoryServiceImpl
}

func NewEventCategoryController(service event_category.EventCategoryServiceImpl) *EventCategoryController {
	return &EventCategoryController{EventCategoryService: service}
}

// @Summary Get all event categories
// @Description Get all event categories
// @Tags Event Category
// @Accept json
// @Produce json
// @Success 200 {object} []response.EventCategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /event-categories [get]
func (c *EventCategoryController) GetEventCategories(ctx echo.Context) error {
	categories, err := c.EventCategoryService.GetAll()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to get event categories",
		})
	}
	eventCategoryResponses := make([]response.EventCategoryResponse, len(categories))
	for i, category := range categories {
		eventCategoryResponses[i] = response.EventCategoryFromEntities(category)
	}
	return pagination.SuccessPaginatedResponse(ctx, eventCategoryResponses, 1, 10, int64(len(eventCategoryResponses)))
}

// @Summary Get event category by ID
// @Description Get event category by ID
// @Tags Event Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the event category"
// @Success 200 {object} response.EventCategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /event-categories/{id} [get]
func (c *EventCategoryController) GetEventCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.EventCategoryService.GetByID(id)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to get event category",
		})
	}
	return base.SuccesResponse(ctx, response.EventCategoryFromEntities(category))
}

// @Summary Create event category
// @Description Create event category
// @Tags Event Category
// @Accept json
// @Produce json
// @Param request body request.EventCategoryRequest true "Event Category Request"
// @Success 201 {object} response.EventCategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /event-categories [post]
func (c *EventCategoryController) CreateEventCategory(ctx echo.Context) error {
	req := new(request.EventCategoryRequest)
	if err := ctx.Bind(req); err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	category, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	createdCategory, err := c.EventCategoryService.Create(category)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to create event category",
		})
	}

	return base.SuccesResponse(ctx, response.EventCategoryFromEntities(createdCategory))
}

// @Summary Update event category
// @Description Update event category
// @Tags Event Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the event category"
// @Param request body request.EventCategoryRequest true "Event Category Request"
// @Success 200 {object} response.EventCategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /event-categories/{id} [put]
func (c *EventCategoryController) UpdateEventCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	req := new(request.EventCategoryRequest)
	if err := ctx.Bind(req); err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	category, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	category.ID = id
	updatedCategory, err := c.EventCategoryService.Update(category)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to update event category",
		})
	}

	return base.SuccesResponse(ctx, response.EventCategoryFromEntities(updatedCategory))
}

// @Summary Delete event category
// @Description Delete event category
// @Tags Event Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the event category"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /event-categories/{id} [delete]
func (c *EventCategoryController) DeleteEventCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.EventCategoryService.Delete(id)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to delete event category",
		})
	}
	return base.SuccesResponse(ctx, "Event category deleted successfully")
}
