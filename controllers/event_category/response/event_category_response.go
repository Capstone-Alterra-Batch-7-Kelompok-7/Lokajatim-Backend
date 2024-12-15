package response

import (
	"lokajatim/entities"
	"time"
)

// EventCategoryResponse is the response for the Event Category controller
// @Description EventCategoryResponse is the response for the Event Category controller
// @Param ID int true "ID of the event category"
// @Param Name string true "Name of the event category"
// @Param CreatedAt string true "Created At of the event category"
// @Param UpdatedAt string true "Updated At of the event category"
type EventCategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EventCategoryFromEntities converts entities.EventCategory to EventCategoryResponse
func EventCategoryFromEntities(category entities.EventCategory) EventCategoryResponse {
	return EventCategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
