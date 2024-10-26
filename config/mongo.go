
package config

type MongoDbConfig struct {
	Host     string `envconfig:"MONGO_HOST" validate:"required"`
	Port     int    `envconfig:"MONGO_PORT" default:"27017" validate:"required"`
	Database string `envconfig:"MONGO_DB" validate:"required"`
	Username string `envconfig:"MONGO_USER" validate:"required"`
	Password string `envconfig:"MONGO_PASS" validate:"required"`
}
