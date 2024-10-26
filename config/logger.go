package config

type LoggerConfig struct {
	LogFormat string `envconfig:"LOG_FORMAT" default:"json" validate:"required,oneof=text json"`
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info" validate:"required,oneof=trace debug info warn error fatal panic"`
}
