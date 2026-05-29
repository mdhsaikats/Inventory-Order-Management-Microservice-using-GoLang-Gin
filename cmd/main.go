package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/config"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/handlers"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/repositories"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/services"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)

	// Repositories
	productRepo := &repositories.ProductRepository{DB: config.DB}
	orderRepo := &repositories.OrderRepository{DB: config.DB}

	// Services
	productService := &services.ProductService{Repo: productRepo}
	orderService := &services.OrderService{
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
	}

	// Handlers
	productHandler := &handlers.ProductHandler{Service: productService}
	orderHandler := &handlers.OrderHandler{Service: orderService}

	api := r.Group("/api")

	api.POST("/products", productHandler.Create)
	api.GET("/products", productHandler.GetAll)
	api.PUT("/products/:id/stock", productHandler.UpdateStock)

	api.POST("/orders", orderHandler.Create)

	r.Run(":8080")
}
