package entities

type ProductPhoto struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UrlPhoto  string `json:"url_photo"`
	ProductID int    `json:"product_id"`
}
