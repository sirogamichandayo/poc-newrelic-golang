package config

import "github.com/dijsilva/golang-api-newrelic/utils"

type configuration struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabse  string
	PostgresPort     string
	NewRelicName     string
	NewRelicToken    string
	NewRelicEnabled  string
}

var Configuration configuration

func SetConfigs() {
	Configuration = configuration{
		PostgresHost:     utils.GetStringEnv(utils.ENV_NAME_POSTGRES_HOST, "localhost"),
		PostgresUser:     utils.GetStringEnv(utils.ENV_NAME_POSTGRES_USER, "user"),
		PostgresPassword: utils.GetStringEnv(utils.ENV_NAME_POSTGRES_PASSWORD, "pass"),
		PostgresDatabse:  utils.GetStringEnv(utils.ENV_NAME_POSTGRES_DATABASE, "users"),
		PostgresPort:     utils.GetStringEnv(utils.ENV_NAME_POSTGRES_PORT, "5432"),
		NewRelicName:     utils.GetStringEnv(utils.ENV_NAME_NEW_RELIC_APP_NAME, "poc-golang-api-newrelic"),
		NewRelicToken:    utils.GetStringEnv(utils.ENV_NAME_NEW_RELIC_TOKEN, ""),
		NewRelicEnabled:  utils.GetStringEnv(utils.ENV_NAME_NEW_RELIC_ENABLED, "true"),
	}
}
