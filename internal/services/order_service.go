package services

import "github.com/gin-gonic/gin"

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(c *gin.Context) {
	
}
