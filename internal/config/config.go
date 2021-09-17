package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port string `env:"PORT,required"`
}

func NewConfig() (*Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return &config, err
	}
	return &config, nil
}
