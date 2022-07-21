package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BotConfig `yaml:"bot"`
	PGConfig  `yaml:"pg"`
}

type BotConfig struct {
	Token string `env-required:"true" env:"TELEGRAM_BOT_TOKEN"`
}

type PGConfig struct {
	User string `env-required:"true" env:"POSTGRES_USER" yaml:"user"`
	DB   string `env-required:"true" env:"POSTGRES_DB" yaml:"db"`
	HOST string `env-required:"true" env:"POSTGRES_HOST" yaml:"host"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./internal/config/config.yml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
