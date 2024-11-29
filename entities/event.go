package entities

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	Capacity    int       `json:"capacity"`
	Price       int   `json:"price"`
	Description string    `json:"description"`
	UrlPhoto    string    `json:"url_photo"`
}