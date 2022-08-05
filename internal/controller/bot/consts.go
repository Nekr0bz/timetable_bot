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
/nextweek - Расписание на след. неделю ➡️
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

	profileMsg = `
Профиль %v
---------------------------—
%v
---------------------------—
Университет: %v
Факультет: %v
Курс: %v
Группа: %v

---------------------------—
Подробнее /info
`

	dayMsg = `
📆 %v 
🗂 [источник](%v)
%v
`

	dayRowMsg = "   %v %v %v %v %v"
)

// Dialogs messages
const (
	choiceDateMsg       = "Выбери день или неделю 📆"
	shareTextMsg        = "Если тебе нравится бот, то поделись им с одногруппниками 💟"
	shareLink           = "https://t.me/share/url?url=Я узнаю расписание нашего университета в TimeTableBot.\nПрисоединяйся\nt.me/TimetableClassesBot"
	choiceUniversityMsg = "Выбери университет:"
	choiceGroupTypeMsg  = "Тип группы:"
	choiceFacultyMsg    = "Выбери факультет:"
	choiceCourseMsg     = "Выбери курс:"
	choiceGroupMsg      = "Выбери группу:"
)

// Buttons texts (should be unique)
const (
	timeTableBtnText              = "Расписание 📚"
	profileBtnText                = "Профиль 👤"
	todayBtnText                  = "Сегодня ⬇️"
	tomorrowBtnText               = "Завтра ➡️"
	weekBtnText                   = "Неделя ⬇️"
	nextWeekBtnText               = "Следующая неделя ➡️"
	shareBtnText                  = "Поделиться 🔗"
	shareReplyBtnText             = "Поделиться 🛸"
	editBtnText                   = "Редактировать ⚙️"
	backToStartBtnText            = "Назад ⬅️"
	backToChoiceUniversityBtnText = "‎Назад ⬅️"
	backToChoiceGroupTypeBtnText  = "‎Назад ⬅️‎"
	backToChoiceFacultyBtnText    = "‎Назад‎ ⬅️‎"
	backToChoiceCourseBtnText     = "‎Н‎азад‎ ⬅️‎"
)
