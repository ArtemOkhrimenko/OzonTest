package config

type Server struct {
	HttpPort uint16 `envconfig:"HTTP_PORT" required:"true"`
	GrpcPort uint16 `envconfig:"GRPC_PORT" required:"true"`
	Host     string `envconfig:"HOST" required:"true"`
}
