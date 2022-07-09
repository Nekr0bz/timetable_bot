package main

import (
	"github.com/Nekr0bz/timetable_bot/internal/cli"
	"github.com/Nekr0bz/timetable_bot/pkg/log"
	"os"
)

func main() {
	cliApp := cli.New()

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal("cli error!", log.FError(err))
	}
}
