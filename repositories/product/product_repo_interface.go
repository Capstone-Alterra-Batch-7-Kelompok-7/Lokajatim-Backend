package product

import (
	"lokajatim/entities"
)

type ProductRepositoryInterface interface {
	GetProducts() ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(id int, product entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
}
