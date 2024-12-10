package response

import (
	"lokajatim/entities"
	"time"
)

type ProductResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	Description string            `json:"description"`
	Rating      float64           `json:"rating"`
	Photos      []string          `json:"photos"`
	CategoryID  int               `json:"category_id"`
	Category    entities.Category `json:"category"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func ProductFromEntities(product entities.Product, photos []entities.ProductPhoto) ProductResponse {

	photoUrls := make([]string, len(photos))
	for i, photo := range photos {
		photoUrls[i] = photo.UrlPhoto
	}

	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Rating:      product.Rating,
		Photos:      photoUrls,
		CategoryID:  product.CategoryID,
		Category:    product.Category,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
