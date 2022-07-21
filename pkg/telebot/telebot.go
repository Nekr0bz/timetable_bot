package telebot

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

func NewTeleBot(token string) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	return tele.NewBot(pref)
}
