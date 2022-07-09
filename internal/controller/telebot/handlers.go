package telebot

import tele "gopkg.in/telebot.v3"

type BotUseCase interface {
	GetHello() string
}

type botHandler struct {
	botUseCase BotUseCase
}

func NewBotHandler(botUseCase BotUseCase) *botHandler {
	return &botHandler{botUseCase: botUseCase}
}

func (h *botHandler) Register(b *tele.Bot) {
	b.Handle("/hello", h.helloHandler)
}

func (h *botHandler) helloHandler(c tele.Context) error {
	s := h.botUseCase.GetHello()
	return c.Send(s)
}
