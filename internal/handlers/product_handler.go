package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/models"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/services"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) Create(c *gin.Context) {
	var body struct {
		Name  string  `json:"name"`
		Stock int     `json:"stock"`
		Price float64 `json:"price"`
	}

	c.ShouldBindJSON(&body)

	err := h.Service.CreateProduct(&models.Product{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	data, _ := h.Service.GetProducts()
	c.JSON(200, data)
}

func (h *ProductHandler) UpdateStock(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var body struct {
		Stock int `json:"stock"`
	}
	c.ShouldBindJSON(&body)

	product, err := h.Service.UpdateStock(uint(id), body.Stock)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}
