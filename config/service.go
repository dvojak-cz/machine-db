package config


type WebServerConfig struct {
	Http Http
	Grpc Grpc
}
type Http struct {
	HttpAddress string `envconfig:"HTTP_ADDRESS" default:"0.0.0.0"`
	HttpPort string `envconfig:"HTTP_PORT" default:"8080"`
}
type Grpc struct {
	GrpcPort string `envconfig:"HTTP_PORT" default:"8088"`
	GrpcAddress string `envconfig:"GRPC_ADDRESS" default:"0.0.0.0"`
}



func (c *Http) GetHttpAddress() string {
	return c.HttpAddress + ":" + c.HttpPort
}
