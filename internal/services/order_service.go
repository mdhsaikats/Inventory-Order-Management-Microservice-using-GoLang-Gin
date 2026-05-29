package services

import (
	"errors"

	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/config"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/repositories"
)

type OrderService struct {
	ProductRepo *repositories.ProductRepository
	OrderRepo   *repositories.OrderRepository
}

type OrderRequest struct {
	Items []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func (s *OrderService) CreateOrder(req OrderRequest) (*models.Order, error) {

	tx := config.DB.Begin()

	var total float64
	var items []models.OrderItem

	for _, i := range req.Items {
		product, err := s.ProductRepo.GetByID(i.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if product.Stock < i.Quantity {
			tx.Rollback()
			return nil, errors.New("not enough stock")
		}

		product.Stock -= i.Quantity
		tx.Save(product)

		item := models.OrderItem{
			ProductID: product.ID,
			Quantity:  i.Quantity,
			Price:     product.Price,
		}

		total += float64(i.Quantity) * product.Price
		items = append(items, item)
	}

	order := models.Order{
		TotalPrice: total,
		Status:     "completed",
		Items:      items,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &order, nil
}
