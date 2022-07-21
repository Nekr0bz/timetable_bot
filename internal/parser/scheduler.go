package parser

import (
	"github.com/jasonlvhit/gocron"
)

func InitScheduler() {
	gocron.Every(1).Second().Do(TaskParseSomeOne)
	// ... tasks ...
}
