package entities

type Event struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `json:"name"`
	CategoryID  int           `json:"category_id"`
	Category    EventCategory `json:"category"`
	Location    string        `json:"location"`
	DateTime    string        `json:"date_time"`
	Capacity    int           `json:"capacity"`
	Price       int           `json:"price"`
	Description string        `json:"description"`
	UrlPhoto    string        `json:"url_photo"`
	Rating      float32       `json:"rating"`
}
