package pkg

import (
	"fmt"

	"github.com/dijsilva/golang-api-newrelic/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NewRelicApp *newrelic.Application

func InitNewrelicApplication() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Configuration.NewRelicName),
		newrelic.ConfigLicense(config.Configuration.NewRelicToken),
		newrelic.ConfigEnabled(config.Configuration.NewRelicEnabled == "true"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		fmt.Printf("Error to start newrelic app")
	}

	NewRelicApp = app
}
