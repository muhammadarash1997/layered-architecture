package handler

import (
	"layered-architecture/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func (this *userHandler) Create(c *gin.Context)

func (this *userHandler) Update(c *gin.Context)