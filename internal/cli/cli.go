package cli

import "github.com/urfave/cli/v2"

func New() *cli.App {
	return &cli.App{
		Name: "Timetable Telegram Bot",
		Commands: []*cli.Command{
			{
				Name: "scheduler",
				Subcommands: []*cli.Command{
					{
						Name:   "run",
						Usage:  "Run Scheduler",
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
