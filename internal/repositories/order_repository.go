package repositories

import (
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}
