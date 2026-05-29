package repositories

import (
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
