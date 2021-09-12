package config

import "github.com/caarlos0/env/v6"

type EnvConfig struct {
	Port string `env:"PORT,required"`
}

func NewEnvConfig() (*EnvConfig, error) {
	config := EnvConfig{}
	if err := env.Parse(&config); err != nil {
		return &config, err
	}
	return &config, nil
}
