package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
