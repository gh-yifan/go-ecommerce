package services

import "github.com/gin-gonic/gin"

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}
func (s *ProductService) ListProducts(c *gin.Context) {

}

func (s *ProductService) GetProduct(c *gin.Context) {}

func (s *ProductService) CreateProduct(c *gin.Context) {}

func (s *ProductService) UpdateProduct(c *gin.Context) {}

func (s *ProductService) DeleteProduct(c *gin.Context) {}
