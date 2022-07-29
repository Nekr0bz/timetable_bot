package bot

import tele "gopkg.in/telebot.v3"

func (h *botHandler) Register(b *tele.Bot) {
	initStaticMenus()

	b.Handle(startCMD, h.startHandler)
	b.Handle(debugCMD, h.debugHandler)
	b.Handle(helpCMD, h.helpHandler)
	b.Handle(infoCMD, h.infoHandler)
	b.Handle(profileCMD, h.profileHandler)
	b.Handle(todayCMD, h.todayHandler)
	b.Handle(tomorrowCMD, h.tomorrowHandler)
	b.Handle(weekCMD, h.weekHandler)
	b.Handle(nextWeekCMD, h.nextWeekHandler)

	b.Handle(&startProfileBtn, h.profileHandler)
	b.Handle(&startTimeTableBtn, h.timeTableHandler)

	b.Handle(&profileTimeTableBtn, h.timeTableHandler)
	b.Handle(&profileShareBtn, h.shareHandler)
	//b.Handle(&profileEditBtn, h.debugHandler)

	b.Handle(&timeTableProfileBtn, h.profileHandler)
	b.Handle(&timeTableTodayBtn, h.todayHandler)
	b.Handle(&timeTableTomorrowBtn, h.tomorrowHandler)
	b.Handle(&timeTableWeekBtn, h.weekHandler)
	b.Handle(&timeTableNextWeekBtn, h.nextWeekHandler)
	b.Handle(&timeTableShareBtn, h.shareHandler)
}
