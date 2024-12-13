package event

import "lokajatim/entities"

type EventRepository interface {
	GetAll() ([]entities.Event, error)
	GetByID(id uint) (*entities.Event, error)
	GetbyBestPrice() ([]entities.Event, error)
	Create(event *entities.Event) error
	Update(event *entities.Event) error
	Delete(id uint) error
}