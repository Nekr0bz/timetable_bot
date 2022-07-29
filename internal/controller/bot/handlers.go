package bot

import (
	"context"
	"fmt"
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

func (h *botHandler) profileHandler(c tele.Context) error {
	// TODO: get user profile
	msg := fmt.Sprintf(profileMsg, 1, 2, 3, 4, 5, 6)
	return c.Send(msg, profileMenu)
}

func (h *botHandler) timeTableHandler(c tele.Context) error {
	return c.Send(choiceDateMsg, timeTableMenu)
}

func (h *botHandler) shareHandler(c tele.Context) error {
	return c.Send(shareTextMsg, shareMenu)
}

// TODO: implement this handlers...
func (h *botHandler) todayHandler(c tele.Context) error {
	return c.Send("не готово", profileMenu)
}

func (h *botHandler) tomorrowHandler(c tele.Context) error {
	return c.Send("не готово", profileMenu)
}

func (h *botHandler) weekHandler(c tele.Context) error {
	return c.Send("не готово", profileMenu)
}

func (h *botHandler) nextWeekHandler(c tele.Context) error {
	return c.Send("не готово", profileMenu)
}
