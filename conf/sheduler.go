package conf

import (
	"github.com/Nekr0bz/timetable_bot/parsers"
	"github.com/jasonlvhit/gocron"
)

func RunScheduler() {
	gocron.Every(1).Second().Do(parsers.TaskParseSomeOne)

	<-gocron.Start()
}