package main

import (
	"github.com/dijsilva/golang-api-newrelic/config"
	"github.com/dijsilva/golang-api-newrelic/pkg"
	"github.com/dijsilva/golang-api-newrelic/repository"
	"github.com/dijsilva/golang-api-newrelic/server"
)

func main() {
	config.SetConfigs()
	repository.ConnectDatabase()
	newRelicApp := pkg.NewNewrelicApplication()
	httpServer := server.CreateHttpServer(newRelicApp)

	httpServer.Run(":" + config.Configuration.AppPort)
}
