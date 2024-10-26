package config

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	Configs  configs
	validate *validator.Validate
)

type configs struct {
	MongoDbConfig MongoDbConfig
	WebServerConfig WebServerConfig
	LoggerConfig LoggerConfig
}

func init() {
	validate = validator.New()
}

func LoadFromEnv() error {
	if err := envconfig.Process("", &Configs); err != nil {
		log.WithError(err).Warning("Failed to load environment variables")
		return err
	}
	return validateConfig()
}

func validateConfig() error {
	if err := validate.Struct(&Configs); err != nil {
		log.WithError(err).Fatal("Failed to validate environment variables")
	}
	return nil
}
