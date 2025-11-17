package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 🧮 Структуры данных
type ChallengeStats struct {
	TotalDays         int
	DaysCompleted     int
	DaysRemaining     int
	ProgressPercent   float64
	CurrentStreak     int
	LongestStreak     int
	LastStudyDate     time.Time
	ProductivityScore int
	WillpowerLevel    string
	MentalState       string
	CodingPower       int
	Level             int
	Experience        int
	NextLevelXP       int
	ExpGainedToday    int
	CurrentMood       string
}

type PersonalGrowth struct {
	GamingAvoided       int
	AdultContentAvoided int
	StudyHours          float64
	SkillsLearned       []string
	LifeCrisesSurvived  int
	MomComplaints       int
	SocialAnxietyLevel  int
	StressLevel         int
	ConfidenceLevel     int
	ProgrammingHours    int
}

type DailyEvent struct {
	Type        string // "obstacle", "victory", "challenge", "quest"
	Description string
}

type Achievement struct {
	Name        string
	Description string
	Type        string // "common", "rare", "epic", "legendary"
	Unlocked    bool
	Date        time.Time
}

type Quest struct {
	Name        string
	Description string
	Completed   bool
	RewardXP    int
}

type MentorAdvice struct {
	Message string
	Emoji   string
}

// 🎪 Глобальные переменные
var (
	stats        ChallengeStats
	growth       PersonalGrowth
	dailyEvents  []DailyEvent
	achievements []Achievement
	quests       []Quest
	momQuotes    []string
	obstacles    []string
	victories    []string
	goFacts      []string
	adviceList   []MentorAdvice
)

// 🎯 Инициализация данных
func initChallenge() {
	// Статистика
	stats = ChallengeStats{
		TotalDays:       100,
		DaysCompleted:   0,
		DaysRemaining:   100,
		ProgressPercent: 0,
		CurrentStreak:   0,
		LongestStreak:   0,
		Level:           1,
		Experience:      0,
		NextLevelXP:     100,
		ExpGainedToday:  0,
		WillpowerLevel:  "Стеклянный",
		MentalState:     "Паника и отрицание",
		CodingPower:     10,
		CurrentMood:     "Ожидание старта",
	}

	// Личный рост
	growth = PersonalGrowth{
		GamingAvoided:       0,
		AdultContentAvoided: 0,
		StudyHours:          0,
		SkillsLearned:       []string{},
		LifeCrisesSurvived:  0,
		MomComplaints:       0,
		SocialAnxietyLevel:  100,
		StressLevel:         100,
		ConfidenceLevel:     0,
		ProgrammingHours:    0,
	}

	// Мамин упрёки
	momQuotes = []string{
		"Опять за компом сидишь? Может, лучше бы в магазин сходил?",
		"В твоём возрасте нормальные люди уже детей кормят, а ты всё в компуктере своём сидишь!",
		"Соседский Коля уже вторую машину купил, а ты всё курьерничаешь!",
		"Сколько можно в этих играх-то сидеть? Работать надо!",
		"Ты бы хоть раз в жизни нормально заработал, а не эти копейки на доставке!",
		"Может, лучше бы вышел и развёз пару заказов? Хоть какие-то деньги!",
		"Нашёл бы уже себе жену, как все нормальные мужички!",
		"И долго ты ещё будешь на моей шее сидеть?",
	}

	// Препятствия
	obstacles = []string{
		"Сосед сверлит стену ровно в момент, когда ты пытаешься понять замыкания",
		"Комп завис в самый разгар изучения горутин",
		"Мама требует вынести мусор во время дебага критического бага",
		"Кошка Муська прошлася по клавиатуре и закоммитила случайные символы",
		"ОКР заставляет переписывать код 10 раз из-за форматирования",
		"СДВГ: начал изучать интерфейсы, переключился на каналы, потом на тесты...",
		"Нет света/Интернета в самый разгар кодинга",
		"Клиент звонит в 3 часа ночи с 'срочным' заказиком в загородный коттеджный посёлок",
		"Сломалась клавиатура, а запасная — в шкафу",
		"Началась миграция данных на сервере, и ты не можешь запустить код",
	}

	// Победы
	victories = []string{
		"Победил панику при виде error handling!",
		"Написал первую работающую горутину без deadlock!",
		"Починил баг, который искал 3 часа, одним символом!",
		"Понял разницу между slice и array без гугления!",
		"Рефакторинг прошёл успешно - ничего не сломал!",
		"Написал первый HTTP-сервер на Go!",
		"Создал свою первую библиотеку!",
		"Успешно прошёл первый собеседование по Go!",
		"Написал первый CLI-инструмент!",
		"Запустил свой первый проект в продакшн!",
	}

	// Факты о Go
	goFacts = []string{
		"Go был создан в Google тремя легендарными айтишниками-программистами: Робом Пайком, Кеном Томпсоном и Робертом Гризмером",
		"Горутины легче потоков ОС - их могут быть миллионы!",
		"Go имеет сборщик мусора, но нет исключений как в Java",
		"interface{} может содержать любое значение - это мощно и страшно одновременно",
		"Каналы - это typed conduits для связи между горутинами",
		"Go формат кода автоматически применяется gofmt",
		"Девиз Go: 'Do not communicate by sharing memory; instead, share memory by communicating'",
		"Go был создан в 2007 году и выпущен в 2009",
		"Go поддерживает встроенные тесты и бенчмарки через пакет testing",
		"Go может компилироваться в WebAssembly!",
	}

	// Советы ментора
	adviceList = []MentorAdvice{
		{"Не бойся ошибок - они твои лучшие учителя", "📚"},
		{"Каждый день кода - это шаг к финансовой свободе", "💰"},
		{"СДВГ - это не проклятие, а суперсила в программировании", "⚡"},
		{"ОКР помогает писать чистый, структурированный код", "🧼"},
		{"Социофобия - не преграда, а твой внутренний голос разума", "🧘"},
		{"Мама не против программирования - она против тунеядства", "👵"},
		{"Изучение Go - это инвестиция в будущее", "📈"},
		{"Ты не один в этом пути - миллионы программистов прошли через всё это", "👥"},
		{"Каждая написанная строка кода - это кирпичик в твою карьеру", "🧱"},
		{"Ты сильнее своих зависимостей - докажи это себе и всем!", "💪"},
	}

	// Инициализация достижений
	initAchievements()
	// Инициализация квестов
	initQuests()
}

func initAchievements() {
	achievements = []Achievement{
		{"Первый день", "Выжил после первого дня", "common", false, time.Time{}},
		{"Неделя без срывов", "Продержался 7 дней подряд", "common", false, time.Time{}},
		{"Месяц без игр", "Продержался 30 дней без игр", "rare", false, time.Time{}},
		{"Полпути", "Преодолел 50 дней", "epic", false, time.Time{}},
		{"Самурай кода", "Продержался 100 дней без срывов", "legendary", false, time.Time{}},
		{"Победитель паники", "Победил панику при виде interface{}", "common", false, time.Time{}},
		{"Горутинный маг", "Написал первую работающую горутину", "common", false, time.Time{}},
		{"Багоискатель", "Починил первый критический баг", "common", false, time.Time{}},
		{"Мамин герой", "Мама впервые похвалила за программирование", "rare", false, time.Time{}},
		{"Код-ас", "Написал 1000 строк кода", "epic", false, time.Time{}},
	}
}

func initQuests() {
	quests = []Quest{
		{"День 1", "Написать первую программу на Go", false, 50},
		{"День 5", "Изучить функции и вернуть несколько значений", false, 75},
		{"День 10", "Создать свою структуру и методы", false, 100},
		{"День 15", "Написать первую горутину", false, 125},
		{"День 20", "Создать простой HTTP-сервер", false, 150},
		{"День 25", "Использовать каналы", false, 175},
		{"День 30", "Написать unit-тесты", false, 200},
		{"День 50", "Создать CLI-инструмент", false, 300},
		{"День 75", "Развернуть приложение с Docker", false, 400},
		{"День 100", "Запустить проект в продакшн", false, 1000},
	}
}

// 🧠 Вычисления
func calculateStats(day int) {
	stats.DaysCompleted = day
	stats.DaysRemaining = stats.TotalDays - stats.DaysCompleted
	stats.ProgressPercent = float64(stats.DaysCompleted) / float64(stats.TotalDays) * 100

	// XP и уровни
	baseXP := 100
	stats.ExpGainedToday = baseXP + (day * 10)
	stats.Experience += stats.ExpGainedToday

	// Уровень
	stats.Level = 1 + (stats.Experience / 1000)
	stats.NextLevelXP = stats.Level * 1000

	// Воля, сила, настроение
	willpowerLevels := []string{"Стеклянный", "Бумажный", "Картонный", "Деревянный", "Железный", "Стальной", "Малахитовый", "Аметистовый", "Алмазный", "Богатырский"}
	mentalStates := []string{
		"Паника и отрицание", "Гнев на компилятор", "Торг с самим собой",
		"Депрессия от ошибок", "Принятие и просветление", "Поток и продуктивность",
		"Просветление Golang-Программиста", "Мастер Go", "Легенда кода",
	}
	moods := []string{
		"Ожидание старта", "Энтузиазм старта", "Формирование привычки",
		"Стабильный прогресс", "Преодоление трудностей", "Уверенность в себе",
		"Просветление Go-Программиста", "Гуру Go-кода",
	}

	willIndex := day / 10
	if willIndex >= len(willpowerLevels) {
		willIndex = len(willpowerLevels) - 1
	}
	mentalIndex := day / 12
	if mentalIndex >= len(mentalStates) {
		mentalIndex = len(mentalStates) - 1
	}
	moodIndex := day / 15
	if moodIndex >= len(moods) {
		moodIndex = len(moods) - 1
	}

	stats.WillpowerLevel = willpowerLevels[willIndex]
	stats.MentalState = mentalStates[mentalIndex]
	stats.CurrentMood = moods[moodIndex]

	stats.CodingPower = 10 + (day * 5)
	if stats.CodingPower > 1000 {
		stats.CodingPower = 1000
	}

	// Уровень стресса уменьшается
	growth.StressLevel = 100 - (day * 2)
	if growth.StressLevel < 0 {
		growth.StressLevel = 0
	}
	// Уверенность растёт
	growth.ConfidenceLevel = day * 2
	if growth.ConfidenceLevel > 100 {
		growth.ConfidenceLevel = 100
	}
	// Тревожность уменьшается
	growth.SocialAnxietyLevel = 100 - (day * 2)
	if growth.SocialAnxietyLevel < 0 {
		growth.SocialAnxietyLevel = 0
	}
}

// 🎨 Прогресс-бар
func generateProgressBar(percent float64, width int) string {
	filled := int(percent / 100 * float64(width))
	empty := width - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "🟩"
	}
	for i := 0; i < empty; i++ {
		bar += "⬜"
	}

	return bar
}

// 🧮 Уровень разработчика
func getLevelByDay(day int) string {
	levels := []string{
		"Новичок 🐣", "Ученик 📚", "Интерн 🔧", "Junior Golang-Программист 💻",
		"Опытный Middle Golang-Разработчик 🚀", "Продвинутый Senior Golang Developer 🏆",
		"Просвещённый Golang-Гуру 🧙", "Легенда Go-Кода 🌟", "Golang-Богатырь ⚡",
	}

	levelIndex := day / 15
	if levelIndex >= len(levels) {
		levelIndex = len(levels) - 1
	}
	return levels[levelIndex]
}

// 🎲 Случайное событие
func getRandomEvent() DailyEvent {
	eventTypes := []string{"obstacle", "victory", "challenge", "quest"}
	eventType := eventTypes[rand.Intn(len(eventTypes))]

	var description string
	switch eventType {
	case "obstacle":
		description = obstacles[rand.Intn(len(obstacles))]
	case "victory":
		description = victories[rand.Intn(len(victories))]
	case "challenge":
		challenges := []string{
			"Напиши программу, которая выводит 'Hello, Go!' 10 раз",
			"Создай функцию, которая складывает два числа",
			"Напиши программу с использованием структуры Person",
			"Создай горутину, которая выводит числа от 1 до 5",
			"Напиши тест для простой функции",
		}
		description = "Ежедневный вызов: " + challenges[rand.Intn(len(challenges))]
	case "quest":
		quests := []string{
			"Прочитай 1 главу документации Go",
			"Напиши 50 строк кода",
			"Создай GitHub-репозиторий для своего проекта",
			"Напиши README.md для своего проекта",
			"Закоммить изменения в репозиторий",
		}
		description = "Ежедневный квест: " + quests[rand.Intn(len(quests))]
	}

	return DailyEvent{Type: eventType, Description: description}
}

// 🏆 Проверка достижений
func checkAchievements(day int) {
	for i := range achievements {
		if !achievements[i].Unlocked {
			switch achievements[i].Name {
			case "Первый день":
				if day >= 1 {
					achievements[i].Unlocked = true
					achievements[i].Date = time.Now()
				}
			case "Неделя без срывов":
				if day >= 7 {
					achievements[i].Unlocked = true
					achievements[i].Date = time.Now()
				}
			case "Месяц без игр":
				if day >= 30 {
					achievements[i].Unlocked = true
					achievements[i].Date = time.Now()
				}
			case "Полпути":
				if day >= 50 {
					achievements[i].Unlocked = true
					achievements[i].Date = time.Now()
				}
			case "Самурай кода":
				if day >= 100 {
					achievements[i].Unlocked = true
					achievements[i].Date = time.Now()
				}
			}
		}
	}
}

// 🧘‍♂️ Ментор Go
func getMentorAdvice() MentorAdvice {
	return adviceList[rand.Intn(len(adviceList))]
}

// 🎪 Основная функция
func main() {
	rand.Seed(time.Now().UnixNano())
	initChallenge()

	// Текущий день (из аргумента командной строки, если есть)
	day := 0
	if len(os.Args) > 1 {
		if d, err := strconv.Atoi(os.Args[1]); err == nil {
			day = d
		}
	}

	// Обновляем статистику
	calculateStats(day)
	checkAchievements(day)

	// Генерируем события
	dailyEvents = []DailyEvent{}
	for i := 0; i < 3; i++ {
		dailyEvents = append(dailyEvents, getRandomEvent())
	}

	// Получаем совет ментора
	advice := getMentorAdvice()

	// 🎨 Вывод
	fmt.Println("")
	fmt.Println("🚀 100daysGo: HARD REBOOT 🚀")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Printf("👤 Герой: Гоша, 37 лет | СДВГ+ОКР+Социофобия\n")
	fmt.Printf("🎯 Миссия: Превратиться из курьера в Go-разработчика\n")
	fmt.Printf("💔 Исходное состояние: Зависимость, бедность, мамины упрёки\n")
	fmt.Printf("❤️ Целевое состояние: Финансовая свобода, уважение, карьера\n")
	fmt.Printf("📚 Сегодня %s мы изучаем: Numeric Types - Boolean\n", "17 ноября 2025")
	fmt.Println()

	// 📅 Информация о дне
	fmt.Printf("🔥 ДЕНЬ БИТВЫ: Day%d\n", day)
	fmt.Printf("📊 Статистика героя:\n")
	fmt.Printf("   🎮 Уровень: %s (Lvl %d)\n", getLevelByDay(day), stats.Level)
	fmt.Printf("   🧠 Опыт: %d/%d XP\n", stats.Experience, stats.NextLevelXP)
	fmt.Printf("   💪 Сила воли: %s\n", stats.WillpowerLevel)
	fmt.Printf("   🧘‍♂️ Ментальное состояние: %s\n", stats.MentalState)
	fmt.Printf("   😊 Настроение: %s\n", stats.CurrentMood)
	fmt.Printf("   💻 Сила кодинга: %d/1000\n", stats.CodingPower)
	fmt.Println()

	// 🎯 Прогресс-бар
	progressBar := generateProgressBar(stats.ProgressPercent, 20)
	fmt.Printf("📈 Прогресс: %s\n", progressBar)
	fmt.Printf("   %d/%d дней (%.1f%%) | Осталось: %d дней\n",
		stats.DaysCompleted, stats.TotalDays, stats.ProgressPercent, stats.DaysRemaining)
	fmt.Printf("   🔥 Серия без срывов: %d дней\n", stats.CurrentStreak)
	fmt.Println()

	// 💫 Мотивация
	fmt.Printf("💫 СЕГОДНЯШНЯЯ МОТИВАЦИЯ:\n")
	motivations := []string{
		"Каждая строка кода - это шаг от зависимости к свободе!",
		"Сегодня ты стал на один день ближе к карьере мечты!",
		"Твой упорство впечатляет! Продолжай в том же духе!",
		"Мама будет гордиться тобой, когда увидит твои результаты!",
		"СДВГ и ОКР - твои суперсилы в программировании!",
		"Игры украли прошлое, Go вернёт будущее!",
		"37 лет - идеальный возраст для перезагрузки!",
		"Социофобия отступает перед уверенностью в своих навыках!",
		"Каждая работающая программа - это удар по бедности!",
		"Ты не просто учишь Go - ты переписываешь свою судьбу!",
	}
	fmt.Printf("   %s\n", motivations[day%len(motivations)])
	fmt.Println()

	// 🧘‍♂️ Совет ментора
	fmt.Printf("🎓 СОВЕТ МЕНТОРА GO:\n")
	fmt.Printf("   %s %s\n", advice.Message, advice.Emoji)
	fmt.Println()

	// 🎪 Случайные события
	fmt.Printf("🎪 СЕГОДНЯШНИЕ СОБЫТИЯ:\n")
	for _, event := range dailyEvents {
		emoji := "❓"
		switch event.Type {
		case "obstacle":
			emoji = "🚧"
		case "victory":
			emoji = "🏆"
		case "challenge":
			emoji = "🎯"
		case "quest":
			emoji = "📜"
		}
		fmt.Printf("   %s %s\n", emoji, event.Description)
	}

	// Мамин упрёк
	fmt.Printf("   👵 Мама: \"%s\"\n", momQuotes[rand.Intn(len(momQuotes))])
	fmt.Println()

	// 🎯 Статистика личного роста
	fmt.Println("🌟 СТАТИСТИКА ЛИЧНОГО РОСТА:")
	growth.GamingAvoided = day * 2
	growth.AdultContentAvoided = day * 3
	growth.StudyHours = float64(day) * 1.5
	growth.ProgrammingHours = day * 2
	growth.LifeCrisesSurvived = day / 7
	growth.MomComplaints = day / 3

	fmt.Printf("   🎮 Игровых сессий пропущено: ~%d\n", growth.GamingAvoided)
	fmt.Printf("   🔞 Вредного контента проигнорировано: ~%d раз\n", growth.AdultContentAvoided)
	fmt.Printf("   📚 Часов изучения: ~%.1f часов\n", growth.StudyHours)
	fmt.Printf("   💻 Часов программирования: %d часов\n", growth.ProgrammingHours)
	fmt.Printf("   😨 Уровень стресса: %d/100\n", growth.StressLevel)
	fmt.Printf("   💪 Уровень уверенности: %d/100\n", growth.ConfidenceLevel)
	fmt.Printf("   😮 Уровень социофобии: %d/100\n", growth.SocialAnxietyLevel)
	fmt.Printf("   🆘 Стрессов и депрессий пережито: %d\n", growth.LifeCrisesSurvived)
	fmt.Printf("   👵 Маминых жалоб и упрёков: %d\n", growth.MomComplaints)

	// 🏆 Достижения
	fmt.Println()
	fmt.Println("🏆 ДОСТИЖЕНИЯ:")
	unlockedCount := 0
	for _, achievement := range achievements {
		if achievement.Unlocked {
			unlockedCount++
			emoji := "❓"
			switch achievement.Type {
			case "common":
				emoji = "🟢"
			case "rare":
				emoji = "🔵"
			case "epic":
				emoji = "🟣"
			case "legendary":
				emoji = "🟡"
			}
			fmt.Printf("   %s %s: %s\n", emoji, achievement.Name, achievement.Description)
		}
	}
	fmt.Printf("   📊 Разблокировано: %d/%d\n", unlockedCount, len(achievements))

	// 📜 Активные квесты
	fmt.Println()
	fmt.Println("📜 АКТИВНЫЕ КВЕСТЫ:")
	activeQuests := 0
	for i := range quests {
		if day >= i*10 && !quests[i].Completed {
			activeQuests++
			status := "❌"
			if quests[i].Completed {
				status = "✅"
			}
			fmt.Printf("   %s %s: %s (Награда: +%d XP)\n", status, quests[i].Name, quests[i].Description, quests[i].RewardXP)
		}
	}
	fmt.Printf("   📊 Активно: %d\n", activeQuests)

	// 💰 Перспективы
	fmt.Println()
	fmt.Println("💰 ПЕРСПЕКТИВЫ:")
	salaryProgress := 80000 + (day * 1200)
	fmt.Printf("   💸 Через %d дней: ~%d руб/мес\n", stats.DaysRemaining, salaryProgress)
	fmt.Printf("   💰 Через %d дней: Возможность финансово помогать маме\n", stats.DaysRemaining/2)
	fmt.Printf("   👑 Через %d дней: Статус 'Уважаемый Golang-Разработчик'\n", stats.DaysRemaining)

	// 🎲 Факт о Go
	fmt.Println()
	fmt.Printf("🎲 ФАКТ О GO: %s\n", goFacts[rand.Intn(len(goFacts))])

	// 🧠 Заключение
	fmt.Println()
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("💡 Напутствие: Каждая ошибка компиляции - это урок.")
	fmt.Println("   Каждая работающая программа - это победа.")
	fmt.Println("   Ты не просто учишь Go - ты переписываешь свою судьбу!")
	fmt.Println("   Мама гордится твоим упорством, скоро будет гордиться и зарплатой!")
	fmt.Println("   👵 \"Ну хоть что-то полезное из тебя выйдет, надеюсь...\"")
}
