package main

import (
	"github.com/dijsilva/golang-api-newrelic/config"
	"github.com/dijsilva/golang-api-newrelic/repository"
	"github.com/dijsilva/golang-api-newrelic/server"
)

func main() {
	config.SetConfigs()
	repository.ConnectDatabase()
	httpServer := server.CreateHttpServer()

	httpServer.Run(":" + config.Configuration.AppPort)
}
