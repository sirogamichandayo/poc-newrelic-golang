package server

import (
	"github.com/dijsilva/golang-api-newrelic/adapters"
	"github.com/dijsilva/golang-api-newrelic/pkg/github"
	"github.com/dijsilva/golang-api-newrelic/repository"
	"github.com/dijsilva/golang-api-newrelic/server/middleware"
	"github.com/dijsilva/golang-api-newrelic/services"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func CreateHttpServer(newRelicApp *newrelic.Application) *gin.Engine {

	app := gin.Default()

	app.Use(
		nrgin.Middleware(newRelicApp),
		middleware.RegisterLogger(newRelicApp),
	)

	userRepository := repository.CreateUserRepository()
	githubClient := github.NewGithubClient()
	userService := services.CreateUserService(userRepository, githubClient)
	userController := adapters.CreateUserController(userService)

	userRoutes := app.Group("users")

	UserHandler(userRoutes, userController)

	return app
}
