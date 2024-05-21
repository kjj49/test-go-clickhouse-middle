package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config
	Config struct {
		HTTP       `yaml:"http"`
		ClickHouse `yaml:"clickhouse"`
		Log        `yaml:"logger"`
	}

	// HTTP
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// ClickHouse
	ClickHouse struct {
		DB       string `env-required:"true" yaml:"database" env:"CLICKHOUSE_DATABASE"`
		Username string `env-required:"true" yaml:"username" env:"CLICKHOUSE_USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"CLICKHOUSE_PASSWORD"`
		Host     string `env-required:"true" yaml:"host" env:"CLICKHOUSE_HOST"`
		Port     int    `env-required:"true" yaml:"port" env:"CLICKHOUSE_PORT"`
	}

	// Log
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
