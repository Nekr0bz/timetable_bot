package bot

// Commands endpoints
const (
	startCMD    = "/start"
	debugCMD    = "/debug"
	helpCMD     = "/help"
	infoCMD     = "/info"
	todayCMD    = "/today"
	tomorrowCMD = "/tomorrow"
	weekCMD     = "/week"
	nextWeekCMD = "/nextweek"
	profileCMD  = "/profile"
)

// Messages for commands
const (
	helpMsg = `
📖 Список команд:

/today - Расписание на сегодня ⬇️
/tomorrow - Расписание на завтра ➡️
/week - Расписание на неделю ⬇️
/nextweek - Расписание на следующую неделю ➡️
/profile - Профиль 👤
/help - Список команд 📖
/info - О нас 📚
/start - Старт 🚀
`

	infoMsg = `
Здравствуй, я TimeTableBot! 👋
Помогаю узнавать расписание занятий в университетах Санкт-Петербурга. 😎 

Чтобы узнать свое расписание нажми /start и заполни "профиль", по этим параметрам я буду искать твое расписание.
`

	startMsg = `
Здравствуй, я TimeTableBot! 👋
Помогаю узнавать расписание занятий в университетах Санкт-Петербурга. 😎 

Подробнее /info
`
)

// Buttons texts
const (
	timeTableBtnText = "Расписание 📚"
	profileBtnText   = "Профиль 👤"
	todayBtnText     = "Сегодня ⬇️"
	tomorrowBtnText  = "Завтра ➡️"
	weekBtnText      = "Неделя ⬇️"
	nextWeekBtnText  = "Следующая неделя ➡️"
	shareBtnText     = "Поделиться 🔗"
	editBtnText      = "Редактировать ⚙️"
)
