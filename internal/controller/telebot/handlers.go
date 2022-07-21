package telebot

import (
	"fmt"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"github.com/Nekr0bz/timetable_bot/pkg/log"
	tele "gopkg.in/telebot.v3"
)

type botHandler struct {
	botUseCase usecase.BotUseCase
}

func NewBotHandler(botUseCase usecase.BotUseCase) *botHandler {
	return &botHandler{botUseCase: botUseCase}
}

func (h *botHandler) Register(b *tele.Bot) {
	b.Handle("/start", h.startHandler)
	b.Handle("/debug", h.debugHandler)
}

func (h *botHandler) debugHandler(c tele.Context) error {
	l := log.Sugar()
	l.Info("BOT CONTEXT ",
		"sender", c.Sender(),
		"Chat", c.Chat(),
		"Data", c.Data(),
		"Args", c.Args(),
	)
	return c.Send("OK")
}

func (h *botHandler) startHandler(c tele.Context) error {
	isNew, err := h.botUseCase.SignUpUser(c)
	if err != nil {
		return err
	}

	return c.Send(fmt.Sprintf("%s", isNew))
}
