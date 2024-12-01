package entities

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	DateTime    string `json:"date_time"`
	Capacity    int       `json:"capacity"`
	Price       int   `json:"price"`
	Description string    `json:"description"`
	UrlPhoto    string    `json:"url_photo"`
}