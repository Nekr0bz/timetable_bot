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
}
