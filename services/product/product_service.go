package product

import (
	"encoding/csv"
	"fmt"
	"lokajatim/entities"
	"lokajatim/repositories/product"
	"os"
	"strconv"
	"strings"
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
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	var products []entities.Product
	var photos []entities.ProductPhoto

	for i, record := range records {
		if i == 0 {
			continue
		}

		price, err := strconv.Atoi(record[2])
		if err != nil {
			return fmt.Errorf("invalid price on row %d: %w", i+1, err)
		}
		stock, err := strconv.Atoi(record[3])
		if err != nil {
			return fmt.Errorf("invalid stock on row %d: %w", i+1, err)
		}
		categoryID, err := strconv.Atoi(record[5])
		if err != nil {
			return fmt.Errorf("invalid category ID on row %d: %w", i+1, err)
		}

		product := entities.Product{
			Name:        record[1],
			Price:       price,
			Stock:       stock,
			Description: record[4],
			CategoryID:  categoryID,
		}
		products = append(products, product)

		if len(record) > 7 && record[7] != "" {
			photoURLs := strings.Split(record[7], ";")
			for _, url := range photoURLs {
				photos = append(photos, entities.ProductPhoto{
					UrlPhoto:  url,
					ProductID: 0,
				})
			}
		}
	}

	if err := s.productRepository.BulkInsert(products); err != nil {
		return fmt.Errorf("failed to insert products: %w", err)
	}

	for _, product := range products {
		for j := range photos {
			if photos[j].ProductID == 0 {
				photos[j].ProductID = product.ID
			}
		}
	}

	if err := s.productRepository.BulkInsertPhotos(photos); err != nil {
		return fmt.Errorf("failed to insert product photos: %w", err)
	}

	return nil
}
