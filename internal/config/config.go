package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config is a common struct contains other configs.
type Config struct {
	Server
	Postgres
	Decoder
}

// LoadConfig loads config.
func LoadConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("unable to decode env into struct: %w", err)
	}

	return cfg, nil
}
