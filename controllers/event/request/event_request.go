package request

import (
	"lokajatim/entities"
)

// EventRequest is the request for the event endpoint
// @Description EventRequest is the request for creating or updating an event
// @Param Name string true "Name of the event"
// @Param CategoryID int true "Category ID of the event"
// @Param Location string true "Location of the event"
// @Param DateTime string true "Date and time of the event"
// @Param Capacity int true "Capacity of the event"
// @Param Price float64 true "Price of the event"
// @Param Description string false "Description of the event"
// @Param UrlPhoto string false "URL of the event photo"
type EventRequest struct {
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Location    string  `json:"location"`
	DateTime    string  `json:"date_time"`
	Capacity    int     `json:"capacity"`
	Price       int     `json:"price"`
	Description string  `json:"description,omitempty"`
	UrlPhoto    string  `json:"url_photo,omitempty"`
	Rating      float32 `json:"rating,omitempty"`
}

// ToEntities converts the EventRequest to Event entity
func (eventRequest EventRequest) ToEntities() entities.Event {
	return entities.Event{
		Name:        eventRequest.Name,
		CategoryID:  eventRequest.CategoryID,
		Location:    eventRequest.Location,
		DateTime:    eventRequest.DateTime,
		Capacity:    eventRequest.Capacity,
		Price:       eventRequest.Price,
		Description: eventRequest.Description,
		UrlPhoto:    eventRequest.UrlPhoto,
		Rating:      eventRequest.Rating,
	}
}
