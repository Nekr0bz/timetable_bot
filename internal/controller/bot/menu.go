package bot

import tele "gopkg.in/telebot.v3"

var (
	startMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	startTimeTableBtn = startMenu.Text(timeTableBtnText)
	startProfileBtn   = startMenu.Text(profileBtnText)
)

// initStaticMenus Initializing Static Menus
func initStaticMenus() {
	startMenu.Reply(
		startMenu.Row(startTimeTableBtn),
		startMenu.Row(startProfileBtn),
	)
}
