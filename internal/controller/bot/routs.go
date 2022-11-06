package bot

import (
	tele "gopkg.in/telebot.v3"
)

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

	b.Handle(&timeTableProfileBtn, h.profileHandler)
	b.Handle(&timeTableTodayBtn, h.todayHandler)
	b.Handle(&timeTableTomorrowBtn, h.tomorrowHandler)
	b.Handle(&timeTableWeekBtn, h.weekHandler)
	b.Handle(&timeTableNextWeekBtn, h.nextWeekHandler)
	b.Handle(&timeTableShareBtn, h.shareHandler)

	h.registerDynamicMenus(b)
}

func (h *botHandler) registerDynamicMenus(b *tele.Bot) {
	universities, err := h.botUseCase.Universities(h.ctx)
	if err != nil {
		return
	}

	h.registerChoiceUniversitiesMenus(b, makeChoiceUniversitiesMenu(universities...))
}

func (h *botHandler) registerChoiceUniversitiesMenus(b *tele.Bot, choiceMenu dialogMenu) {
	b.Handle(&profileEditBtn, h.makeMenuHandler(choiceUniversityMsg, choiceMenu.menu))
	b.Handle(&choiceMenu.backBtn, h.startHandler)

	for _, dbtn := range choiceMenu.dialogBtns {
		b.Handle(&dbtn.btn, h.makeMenuHandler(choiceGroupTypeMsg, dbtn.dialogMenu.menu))
		h.registerChoiceGroupTypeMenus(b, dbtn.dialogMenu)
	}
}

func (h *botHandler) registerChoiceGroupTypeMenus(b *tele.Bot, choiceMenu dialogMenu) {
	b.Handle(&choiceMenu.backBtn, h.makeMenuHandler(choiceUniversityMsg, choiceMenu.parentMenu))

	for _, dbtn := range choiceMenu.dialogBtns {
		b.Handle(&dbtn.btn, h.makeMenuHandler(choiceFacultyMsg, dbtn.dialogMenu.menu))
		h.registerChoiceFacultyMenus(b, dbtn.dialogMenu)
	}
}

func (h *botHandler) registerChoiceFacultyMenus(b *tele.Bot, choiceMenu dialogMenu) {
	b.Handle(&choiceMenu.backBtn, h.makeMenuHandler(choiceGroupTypeMsg, choiceMenu.parentMenu))

	for _, dbtn := range choiceMenu.dialogBtns {
		b.Handle(&dbtn.btn, h.makeMenuHandler(choiceCourseMsg, dbtn.dialogMenu.menu))
		h.registerChoiceCourseMenus(b, dbtn.dialogMenu)
	}
}

func (h *botHandler) registerChoiceCourseMenus(b *tele.Bot, choiceMenu dialogMenu) {
	b.Handle(&choiceMenu.backBtn, h.makeMenuHandler(choiceFacultyMsg, choiceMenu.parentMenu))

	for _, dbtn := range choiceMenu.dialogBtns {
		b.Handle(&dbtn.btn, h.makeMenuHandler(choiceGroupMsg, dbtn.dialogMenu.menu))
		h.registerChoiceGroupMenus(b, dbtn.dialogMenu)
	}
}

func (h *botHandler) registerChoiceGroupMenus(b *tele.Bot, choiceMenu dialogMenu) {
	b.Handle(&choiceMenu.backBtn, h.makeMenuHandler(choiceCourseMsg, choiceMenu.parentMenu))

	for _, dbtn := range choiceMenu.dialogBtns {
		b.Handle(&dbtn.btn, h.profileHandler)
	}
}
