package response

import (
	"lokajatim/entities"
)

// EventResponse is the response for the event controller
// @Description EventResponse is the response for event data retrieval
// @Param ID int true "ID of the event"
// @Param Name string true "Name of the event"
// @Param Category string true "Category of the event"
// @Param Location string true "Location of the event"
// @Param DateTime string true "Date and time of the event"
// @Param Capacity int true "Capacity of the event"
// @Param Price float64 true "Price of the event"
// @Param Description string false "Description of the event"
// @Param UrlPhoto string false "URL of the event photo"
type EventResponse struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Category    entities.EventCategory `json:"category"`
	Location    string                 `json:"location"`
	DateTime    string                 `json:"date_time"`
	Capacity    int                    `json:"capacity"`
	Price       int                    `json:"price"`
	Description string                 `json:"description,omitempty"`
	UrlPhoto    string                 `json:"url_photo,omitempty"`
	Rating      float32                `json:"rating,omitempty"`
}

// EventFromEntities maps Event entity to EventResponse
func EventFromEntities(event entities.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		Name:        event.Name,
		Category:    event.Category,
		Location:    event.Location,
		DateTime:    event.DateTime,
		Capacity:    event.Capacity,
		Price:       event.Price,
		Description: event.Description,
		UrlPhoto:    event.UrlPhoto,
		Rating:      event.Rating,
	}
}
