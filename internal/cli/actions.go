package cli

import (
	"github.com/Nekr0bz/timetable_bot/config"
	"github.com/Nekr0bz/timetable_bot/internal/controller/telebot"
	"github.com/Nekr0bz/timetable_bot/internal/scheduler"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"github.com/Nekr0bz/timetable_bot/pkg/log"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/cli/v2"
	tele "gopkg.in/telebot.v3"
	"time"
)

func RunBot(cCtx *cli.Context) (err error) {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Config err", log.FError(err))
	}

	pref := tele.Settings{
		Token:  cfg.TelegramBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Telegram Bot err", log.FError(err))
	}

	botHandler := telebot.NewBotHandler(usecase.NewBotUseCase())
	botHandler.Register(b)

	log.Info("Start Telegram Bot...")
	b.Start()
	return
}

func RunScheduler(cCtx *cli.Context) (err error) {
	scheduler.InitScheduler()
	log.Info("Start Scheduler Tasks...")
	<-gocron.Start()
	return
}
