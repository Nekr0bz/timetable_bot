package scheduler

import (
	"github.com/Nekr0bz/timetable_bot/internal/parser"
	"github.com/jasonlvhit/gocron"
)

func InitScheduler() {
	gocron.Every(1).Second().Do(parser.TaskParseSomeOne)
	// ... tasks ...
}
