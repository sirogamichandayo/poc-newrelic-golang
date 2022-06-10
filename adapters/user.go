package adapters

import (
	"net/http"

	"github.com/dijsilva/golang-api-newrelic/dtos"
	"github.com/dijsilva/golang-api-newrelic/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	ListUsers(ctx *gin.Context)
}

type userController struct {
	service services.UserService
}

func CreateUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var userDTO dtos.User

	if err := ctx.ShouldBind(&userDTO); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := c.service.Create(userDTO, ctx)

	if err.Err != nil {
		ctx.AbortWithStatus(err.Status())
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *userController) ListUsers(ctx *gin.Context) {

	users, err := c.service.List(ctx)

	if err.Err != nil {
		ctx.AbortWithStatus(err.Status())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
