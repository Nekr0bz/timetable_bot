package bot

import (
	"context"
	"fmt"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"github.com/Nekr0bz/timetable_bot/pkg/utils"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"strings"
	"time"
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

func (h *botHandler) makeMenuHandler(msg string, menu *tele.ReplyMarkup) tele.HandlerFunc {
	return func(c tele.Context) error {
		return c.Send(msg, menu)
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

	// TODO: Edit dialogMenu if user is new
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

func (h *botHandler) todayHandler(c tele.Context) error {
	today := time.Now().Format("02 Jan 2006")

	rows := []string{
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
	}

	msg := fmt.Sprintf(dayMsg, today, "example.com", strings.Join(rows, "\n"))

	return c.Send(msg, timeTableMenu, tele.ModeMarkdown)
}

func (h *botHandler) tomorrowHandler(c tele.Context) error {
	tomorrow := time.Now().AddDate(0, 0, 1).Format("02 Jan 2006")

	rows := []string{
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
	}

	msg := fmt.Sprintf(dayMsg, tomorrow, "example.com", strings.Join(rows, "\n"))

	return c.Send(msg, timeTableMenu, tele.ModeMarkdown)
}

func (h *botHandler) weekHandler(c tele.Context) error {
	firstDay := utils.GetMonday(time.Now())
	lastDay := firstDay.AddDate(0, 0, 6)
	date := firstDay.Format("02 Jan 2006") + " - " + lastDay.Format("02 Jan 2006")

	rows := []string{
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
	}

	msg := fmt.Sprintf(dayMsg, date, "example.com", strings.Join(rows, "\n"))

	return c.Send(msg, timeTableMenu, tele.ModeMarkdown)
}

func (h *botHandler) nextWeekHandler(c tele.Context) error {
	firstDay := utils.GetMonday(time.Now().AddDate(0, 0, 7))
	lastDay := firstDay.AddDate(0, 0, 6)
	date := firstDay.Format("02 Jan 2006") + " - " + lastDay.Format("02 Jan 2006")

	rows := []string{
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
		fmt.Sprintf(dayRowMsg, 1, 2, 3, 3, 4),
	}

	msg := fmt.Sprintf(dayMsg, date, "example.com", strings.Join(rows, "\n"))

	return c.Send(msg, timeTableMenu, tele.ModeMarkdown)
}
