package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  App  `yaml:"app"`
	HTTP HTTP `yaml:"http"`
	Log  Log  `yaml:"logger"`
	DB   DB   `yaml:"databases"`
	RMQ  RMQ  `yaml:"rabbitmq"`
}

// NewConfig returns app config (read from root).
func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{}
	if err = cleanenv.ReadConfig("./config.yaml", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
