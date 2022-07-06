package main

import (
	"github.com/Nekr0bz/timetable_bot/conf"
	"github.com/Nekr0bz/timetable_bot/tg_bot"
	"github.com/Nekr0bz/timetable_bot/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	conf.InitSettings()
	utils.InitLogger()

	app := &cli.App{
		Name: "Timetable Telegram Bot",
		Commands: []*cli.Command{
			{
				Name: "scheduler",
				Subcommands: []*cli.Command{
					{
						Name:  "run",
						Usage: "Run Scheduler",
						Action: func(cCtx *cli.Context) error {
							conf.RunScheduler()
							return nil
						},
					},
				},
			},
			{
				Name: "bot",
				Subcommands: []*cli.Command{
					{
						Name:  "run",
						Usage: "Run Telegram Bot",
						Action: func(cCtx *cli.Context) error {
							tg_bot.RunBot()
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}