package config

type Postgres struct {
	DSN string `envconfig:"DSN" required:"true"`
}
