package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdhsaikats/Inventory-Order-Management-Microservice-using-GoLang-Gin/internal/services"
)

type OrderHandler struct {
	Service *services.OrderService
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req services.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.Service.CreateOrder(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}
