package conf

import "github.com/ilyakaznacheev/cleanenv"

type settings struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

var Settings settings

func InitSettings() {
	cleanenv.ReadEnv(&Settings)
}
