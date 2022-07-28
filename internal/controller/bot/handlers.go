package bot

import (
	"context"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type botHandler struct {
	ctx        context.Context // TODO: I haven't found a better solution yet.
	log        *zap.Logger
	botUseCase usecase.BotUseCase
}

func NewBotHandler(ctx context.Context, log *zap.Logger, botUseCase usecase.BotUseCase) *botHandler {
	return &botHandler{
		ctx:        ctx,
		log:        log,
		botUseCase: botUseCase,
	}
}

func (h *botHandler) Register(b *tele.Bot) {
	initStaticMenus()

	b.Handle(startCMD, h.startHandler)
	b.Handle(debugCMD, h.debugHandler)
	b.Handle(helpCMD, h.helpHandler)
	b.Handle(infoCMD, h.infoHandler)
}

// TODO: remove this handler
func (h *botHandler) debugHandler(c tele.Context) error {
	l := h.log.Sugar()
	l.Info("BOT CONTEXT ",
		"sender", c.Sender(),
		"Chat", c.Chat(),
		"Data", c.Data(),
		"Args", c.Args(),
	)
	return c.Send("OK")
}

func (h *botHandler) startHandler(c tele.Context) error {
	isNew, err := h.botUseCase.SignUpUser(h.ctx, c)
	if err != nil {
		h.log.Error("SignUpUser err", zap.Error(err))
		return err
	}

	// TODO: Edit menu if user is new
	if isNew {
		return c.Send(startMsg, startMenu)
	}

	return c.Send(startMsg, startMenu)
}

func (h *botHandler) helpHandler(c tele.Context) error {
	return c.Send(helpMsg)
}

func (h *botHandler) infoHandler(c tele.Context) error {
	return c.Send(infoMsg)
}
