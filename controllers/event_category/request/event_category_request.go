package request

import "lokajatim/entities"

// EventCategoryRequest is the request for the Event Category endpoint
// @Description EventCategoryRequest is the request for the Event Category endpoint
// @Param Name string true "Name of the event category"
type EventCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (eventCategoryRequest EventCategoryRequest) ToEntities() (entities.EventCategory, error) {
	return entities.EventCategory{
		Name: eventCategoryRequest.Name,
	}, nil
}
