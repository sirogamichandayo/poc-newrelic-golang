package utils

import (
	"fmt"
	"os"
)

func GetStringEnv(envName string, defaultValue string) string {
	value, isDefined := os.LookupEnv(envName)
	if !isDefined {
		fmt.Printf("Variable not defined - %s. Using default value\n", envName)
		return defaultValue
	}
	return value
}
