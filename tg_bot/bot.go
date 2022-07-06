package tg_bot

import (
	"github.com/Nekr0bz/timetable_bot/conf"
	"log"
	"sync"
	"time"

	tele "gopkg.in/telebot.v3"
)

var once sync.Once

type MyBot struct{ *tele.Bot }

func makeBot() MyBot {
	pref := tele.Settings{
		Token:  conf.Settings.TelegramBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Telegram Bot has not been created!", err)
	}
	return MyBot{bot}
}

func GetBot() MyBot {
	var tgBot MyBot
	once.Do(func() {
		tgBot = makeBot()
	})
	return tgBot
}

func (bot MyBot) setHandlers() {
	bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})
}

func RunBot() {
	bot := GetBot()
	bot.setHandlers()
	bot.Start()
}
