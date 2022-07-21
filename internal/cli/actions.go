package cli

import (
	"github.com/Nekr0bz/timetable_bot/internal/config"
	"github.com/Nekr0bz/timetable_bot/internal/controller/bot"
	"github.com/Nekr0bz/timetable_bot/internal/parser"
	"github.com/Nekr0bz/timetable_bot/internal/usecase"
	"github.com/Nekr0bz/timetable_bot/internal/usecase/repo"
	"github.com/Nekr0bz/timetable_bot/pkg/logger"
	"github.com/Nekr0bz/timetable_bot/pkg/postgres"
	"github.com/Nekr0bz/timetable_bot/pkg/telebot"
	"github.com/Nekr0bz/timetable_bot/pkg/telebot/middleware"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func RunBot(cCtx *cli.Context) (err error) {
	// Create logger
	log := logger.NewAppLogger()

	// Crete configs
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Config err", zap.Error(err))
	}

	// Create connection to PG
	db, err := postgres.New(cfg.PGConfig.DB, cfg.PGConfig.HOST, cfg.PGConfig.User, log)
	if err != nil {
		log.Fatal("Postgres err", zap.Error(err))
	}

	// Create Telegram Bot
	b, err := telebot.NewTeleBot(cfg.BotConfig.Token)
	if err != nil {
		log.Fatal("Telegram Bot err", zap.Error(err))
	}

	// Set Telegram Bot Middlewares
	b.Use(middleware.Logger(log))

	// Register Telegram Bot Handler
	// TODO: refactor
	botHandler := bot.NewBotHandler(usecase.NewBotUseCase(repo.NewUserRepo(db)))
	botHandler.Register(b)

	// Start...
	log.Info("Start Telegram Bot...")
	b.Start()
	return
}

func RunScheduler(cCtx *cli.Context) (err error) {
	// TODO: share struct {log, cfg, db...} to tasks

	log := logger.GetLogger()

	parser.InitScheduler()
	log.Info("Start Scheduler Tasks...")

	<-gocron.Start()
	return
}
