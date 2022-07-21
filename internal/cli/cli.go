package cli

import "github.com/urfave/cli/v2"

func New() *cli.App {
	return &cli.App{
		Name: "Timetable Telegram Bot",
		Commands: []*cli.Command{
			{
				Name: "parser",
				Subcommands: []*cli.Command{
					{
						Name:   "scheduler",
						Usage:  "Run Scheduler of Parsers",
						Action: RunScheduler,
					},
				},
			},
			{
				Name: "bot",
				Subcommands: []*cli.Command{
					{
						Name:   "run",
						Usage:  "Run Telegram Bot",
						Action: RunBot,
					},
				},
			},
		},
	}
}
