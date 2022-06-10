package server

import (
	"github.com/dijsilva/golang-api-newrelic/adapters"
	"github.com/gin-gonic/gin"
)

func UserHandler(router *gin.RouterGroup, controller adapters.UserController) {
	router.POST("new", controller.CreateUser)
	router.GET("", controller.ListUsers)
}
