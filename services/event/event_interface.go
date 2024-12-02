package event

import "lokajatim/entities"

type EventService interface {
	GetAll() ([]entities.Event, error)
	GetByID(id uint) (*entities.Event, error)
	Create(event *entities.Event) error
	Update(event *entities.Event) error
	Delete(id uint) error
}