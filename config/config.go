package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type (
	Config struct {
		Postgres `envPrefix:"POSTGRES_"`
		GRPC     `envPrefix:"GRPC_"`
		Gateway  `envPrefix:"GATEWAY_"`
	}

	Postgres struct {
		URL string `env:"URL"`
	}

	GRPC struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}

	Gateway struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "load env file")
	}

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Wrap(err, "parse env file")
	}

	return &cfg, nil
}
