package product

import (
	"lokajatim/entities"
	"lokajatim/repositories/product"
)

type ProductService struct {
	productRepository product.ProductRepositoryInterface
}

func NewProductService(productRepo product.ProductRepositoryInterface) *ProductService {
	return &ProductService{productRepository: productRepo}
}

func (s *ProductService) GetProducts() ([]entities.Product, error) {
	return s.productRepository.GetProducts()
}

func (s *ProductService) GetProductByID(id int) (entities.Product, error) {
	return s.productRepository.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product entities.Product) (entities.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product entities.Product) (entities.Product, error) {
	return s.productRepository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepository.DeleteProduct(id)
}
