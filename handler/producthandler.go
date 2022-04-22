package handler

import (
	"fmt"
	"g-mysql/entity"
	"g-mysql/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerProduct struct {
	service service.IProductService
}

func NewHandlerProduct(service service.IProductService) *HandlerProduct {
	return &HandlerProduct{service}
}

func (h *HandlerProduct) SaveData(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, "Wrong input!")
	}
	a, b := h.service.SaveProduct(product)
	if b != nil {
		fmt.Println("Error!")
	}
	c.JSON(http.StatusOK, a)
}

func (h *HandlerProduct) FindAllData(c *gin.Context) {
	var product []entity.Product
	a, b := h.service.FindAllProduct(product)
	if b != nil {
		panic(b.Error())
	}
	c.JSON(200, a)
}
