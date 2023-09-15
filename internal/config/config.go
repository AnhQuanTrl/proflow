package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	// AppConfig contains application configuration.
	AppConfig struct {
		LogLevel string
	}
)

// InitConfig initializes application configuration from environment variables.
//
// If the .env file is present in the caller's directory, it will be loaded.
func InitConfig() (AppConfig, error) {
	appConfig := AppConfig{}
	if err := godotenv.Load(); err != nil {
		return appConfig, err
	}

	err := envconfig.Process("", &appConfig)
	return appConfig, err
}
