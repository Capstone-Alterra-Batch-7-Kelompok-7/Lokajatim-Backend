package product

import (
	"encoding/csv"
	"errors"
	"lokajatim/entities"
	"lokajatim/repositories/product"
	"os"
	"strconv"
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

func (s *ProductService) GetBestProductsPrice() ([]entities.Product, error) {
	return s.productRepository.GetBestProductsPrice()
}

func (s *ProductService) CreateProduct(product entities.Product) (entities.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	return s.productRepository.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepository.DeleteProduct(id)
}

func (s *ProductService) CreateProductPhotos(photos []entities.ProductPhoto) error {
	return s.productRepository.CreateProductPhotos(photos)
}

func (s *ProductService) GetProductPhotos(productID int) ([]entities.ProductPhoto, error) {
	return s.productRepository.GetProductPhotos(productID)
}

func (s *ProductService) UpdateProductPhotos(productID int, photos []entities.ProductPhoto) error {
	return s.productRepository.UpdateProductPhotos(productID, photos)
}

func (s *ProductService) DeleteProductPhotos(productID int) error {
	return s.productRepository.DeleteProductPhotos(productID)
}

func (s *ProductService) ImportProducts(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.New("failed to read CSV file")
	}

	var products []entities.Product
	for i, record := range records {
		if i == 0 {
			continue
		}
		price, _ := strconv.Atoi(record[2])
		stock, _ := strconv.Atoi(record[3])
		categoryID, _ := strconv.Atoi(record[5])
		rating, _ := strconv.ParseFloat(record[6], 64)

		products = append(products, entities.Product{
			Name:        record[1],
			Price:       price,
			Stock:       stock,
			Description: record[4],
			CategoryID:  categoryID,
			Rating:      rating,
		})
	}

	if err := s.productRepository.BulkInsert(products); err != nil {
		return err
	}

	return nil
}
