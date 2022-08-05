package bot

import (
	"github.com/Nekr0bz/timetable_bot/internal/entity"
	tele "gopkg.in/telebot.v3"
)

// static menus and buttons
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

type dialogMenu struct {
	menu       *tele.ReplyMarkup
	backBtn    tele.Btn
	parentMenu *tele.ReplyMarkup
	dialogBtns []dialogBtn
}

type dialogBtn struct {
	btn        tele.Btn
	dialogMenu dialogMenu
}

func makeChoiceUniversitiesMenu(universities ...entity.University) dialogMenu {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}
	backBtn := menu.Text(backToStartBtnText)

	dialogBtns := make([]dialogBtn, 0, len(universities))
	rows := make([]tele.Row, 0, len(universities)+1)

	rows = append(rows, menu.Row(backBtn))

	for _, university := range universities {
		btn := menu.Text(university.Name)
		dialogBtns = append(dialogBtns, dialogBtn{
			btn:        btn,
			dialogMenu: makeChoiceGroupTypeMenu(&menu, university.GroupTypes),
		})
		rows = append(rows, menu.Row(btn))
	}

	menu.Reply(rows...)
	return dialogMenu{
		menu:       &menu,
		backBtn:    backBtn,
		dialogBtns: dialogBtns,
	}
}

func makeChoiceGroupTypeMenu(parentMenu *tele.ReplyMarkup, groupTypes []entity.GroupType) dialogMenu {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}
	backBtn := menu.Text(backToChoiceUniversityBtnText)

	btns := make([]tele.Btn, 0, len(groupTypes))
	dialogBtns := make([]dialogBtn, 0, len(groupTypes))

	for _, groupType := range groupTypes {
		btn := menu.Text(groupType.Name)
		btns = append(btns, btn)
		dialogBtns = append(dialogBtns, dialogBtn{
			btn:        btn,
			dialogMenu: makeChoiceFacultyMenu(&menu, groupType.Faculties),
		})
	}

	menu.Reply(
		menu.Row(btns...),
		menu.Row(backBtn),
	)

	return dialogMenu{
		menu:       &menu,
		backBtn:    backBtn,
		parentMenu: parentMenu,
		dialogBtns: dialogBtns,
	}
}

func makeChoiceFacultyMenu(parentMenu *tele.ReplyMarkup, faculties []entity.Faculty) dialogMenu {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}
	backBtn := menu.Text(backToChoiceGroupTypeBtnText)

	dialogBtns := make([]dialogBtn, 0, len(faculties))
	rows := make([]tele.Row, 0, len(faculties)+1)

	rows = append(rows, menu.Row(backBtn))

	for _, faculty := range faculties {
		btn := menu.Text(faculty.Name)
		dialogBtns = append(dialogBtns, dialogBtn{
			btn:        btn,
			dialogMenu: makeChoiceCourseMenu(&menu, faculty.Courses),
		})
		rows = append(rows, menu.Row(btn))
	}

	menu.Reply(rows...)
	return dialogMenu{
		menu:       &menu,
		backBtn:    backBtn,
		parentMenu: parentMenu,
		dialogBtns: dialogBtns,
	}
}

func makeChoiceCourseMenu(parentMenu *tele.ReplyMarkup, courses []entity.Course) dialogMenu {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}
	backBtn := menu.Text(backToChoiceFacultyBtnText)

	btns := make([]tele.Btn, 0, len(courses))
	dialogBtns := make([]dialogBtn, 0, len(courses))

	for _, course := range courses {
		btn := menu.Text(course.Name)
		btns = append(btns, btn)
		dialogBtns = append(dialogBtns, dialogBtn{
			btn:        btn,
			dialogMenu: makeChoiceGroupMenu(&menu, course.Groups),
		})
	}

	menu.Reply(
		menu.Row(btns...),
		menu.Row(backBtn),
	)

	return dialogMenu{
		menu:       &menu,
		backBtn:    backBtn,
		parentMenu: parentMenu,
		dialogBtns: dialogBtns,
	}
}

func makeChoiceGroupMenu(parentMenu *tele.ReplyMarkup, groups []entity.Group) dialogMenu {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}
	backBtn := menu.Text(backToChoiceCourseBtnText)

	dialogBtns := make([]dialogBtn, 0, len(groups))
	rows := make([]tele.Row, 0, len(groups)+1)

	rows = append(rows, menu.Row(backBtn))

	for _, group := range groups {
		btn := menu.Text(group.Name)
		dialogBtns = append(dialogBtns, dialogBtn{
			btn: btn,
		})
		rows = append(rows, menu.Row(btn))
	}

	menu.Reply(rows...)
	return dialogMenu{
		menu:       &menu,
		backBtn:    backBtn,
		parentMenu: parentMenu,
		dialogBtns: dialogBtns,
	}
}
