package bot

import tele "gopkg.in/telebot.v3"

var (
	startMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	startTimeTableBtn = startMenu.Text(timeTableBtnText)
	startProfileBtn   = startMenu.Text(profileBtnText)

	profileMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	profileTimeTableBtn = profileMenu.Text(timeTableBtnText)
	profileShareBtn     = profileMenu.Text(shareBtnText)
	profileEditBtn      = profileMenu.Text(editBtnText)

	timeTableMenu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	timeTableProfileBtn  = timeTableMenu.Text(profileBtnText)
	timeTableTodayBtn    = timeTableMenu.Text(todayBtnText)
	timeTableTomorrowBtn = timeTableMenu.Text(tomorrowBtnText)
	timeTableWeekBtn     = timeTableMenu.Text(weekBtnText)
	timeTableNextWeekBtn = timeTableMenu.Text(nextWeekBtnText)
	timeTableShareBtn    = timeTableMenu.Text(shareBtnText)

	shareMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	shareBtn  = shareMenu.URL(shareReplyBtnText, shareLink)
)

// initStaticMenus Initializing Static Menus
func initStaticMenus() {
	startMenu.Reply(
		startMenu.Row(startTimeTableBtn),
		startMenu.Row(startProfileBtn),
	)
	profileMenu.Reply(
		profileMenu.Row(profileEditBtn, profileTimeTableBtn),
		profileMenu.Row(profileShareBtn),
	)

	timeTableMenu.Reply(
		timeTableMenu.Row(timeTableTodayBtn, timeTableTomorrowBtn),
		timeTableMenu.Row(timeTableWeekBtn, timeTableNextWeekBtn),
		timeTableMenu.Row(timeTableShareBtn, timeTableProfileBtn),
	)

	shareMenu.Inline(
		shareMenu.Row(shareBtn),
	)
}
