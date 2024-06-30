package services

import "github.com/gin-gonic/gin"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(c *gin.Context) {}

func (s *UserService) Authenticate(c *gin.Context) {}
