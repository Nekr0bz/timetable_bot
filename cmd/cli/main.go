package main

import (
	"github.com/Nekr0bz/timetable_bot/internal/cli"
	"github.com/Nekr0bz/timetable_bot/pkg/logger"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger.InitAppLogger()
	log := logger.GetLogger()

	cliApp := cli.New()

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal("MAIN APP ERROR! ", zap.Error(err))
	}
}
