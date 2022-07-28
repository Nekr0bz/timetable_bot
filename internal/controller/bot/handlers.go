package bot

import (
	"fmt"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"github.com/Nekr0bz/timetable_bot/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type botHandler struct {
	botUseCase usecase.BotUseCase
}

func NewBotHandler(botUseCase usecase.BotUseCase) *botHandler {
	return &botHandler{botUseCase: botUseCase}
}

func (h *botHandler) Register(b *tele.Bot) {
	b.Handle(startCMD, h.startHandler)
	b.Handle(debugCMD, h.debugHandler)
	b.Handle(helpCMD, h.helpHandler)
	b.Handle(infoCMD, h.infoHandler)
}

// TODO: remove this handler
func (h *botHandler) debugHandler(c tele.Context) error {
	l := logger.GetLogger().Sugar()
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

func (h *botHandler) helpHandler(c tele.Context) error {
	return c.Send(helpMsg)
}

func (h *botHandler) infoHandler(c tele.Context) error {
	return c.Send(infoMsg)
}
