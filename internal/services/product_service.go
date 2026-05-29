package services

import (
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.Create(product)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) UpdateStock(id uint, stock int) (*models.Product, error) {
	product, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	product.Stock = stock
	err = s.Repo.Update(product)

	return product, err
}
