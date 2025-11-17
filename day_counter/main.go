package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 🧮 Основные данные
var (
	startDate     = time.Date(2025, time.October, 25, 0, 0, 0, 0, time.UTC)
	today         = time.Now().UTC()
	currentDay    = calculateCurrentDay()
	stats         = initStats()
	growth        = initGrowth()
	dailyEvents   = generateDailyEvents(3)
	momQuote      = getRandomQuote(momQuotes)
	mentorAdvice  = getRandomAdvice()
	factOfTheDay  = getRandomFact()
	motivation    = getDailyMotivation()
	unlockedAchvs = countUnlockedAchievements()
	activeQuests  = countActiveQuests()
)

// 📅 Автовычисление текущего дня челленджа
func calculateCurrentDay() int {
	days := int(today.Sub(startDate).Hours() / 24)
	if days < 0 {
		return 0
	}
	if days > 100 {
		return 100
	}
	return days
}

// 🧠 Инициализация статистики
func initStats() ChallengeStats {
	percent := float64(currentDay) / 100 * 100
	level := 1 + (100+currentDay*10)/1000 // Исправлено вычисление уровня

	return ChallengeStats{
		DaysCompleted:   currentDay,
		DaysRemaining:   100 - currentDay,
		ProgressPercent: percent,
		Level:           level,
		Experience:      100 + currentDay*10, // ✅ Эквивалентно xp
		NextLevelXP:     level * 1000,
		WillpowerLevel:  getWillpowerLevel(currentDay),
		MentalState:     getMentalState(currentDay),
		CurrentMood:     getCurrentMood(currentDay),
		CodingPower:     min(10+currentDay*5, 1000),
		Streak:          currentDay,
		MaxStreak:       currentDay,
	}
}

// 🌱 Инициализация личностного роста
func initGrowth() PersonalGrowth {
	stress := max(0, 100-currentDay*2)
	confidence := min(100, currentDay*2)
	anxiety := max(0, 100-currentDay*3)

	return PersonalGrowth{
		GamingSkipped:    currentDay * 2,
		StudyHours:       float64(currentDay) * 1.5,
		CodeLines:        currentDay * 50,
		ConfidenceLevel:  confidence,
		StressLevel:      stress,
		SocialEnergy:     100 - anxiety,
		MomPrideLevel:    min(100, currentDay),
		RealLifeHours:    currentDay * 3,
		DaysWithoutPanic: currentDay,
	}
}

// 🎯 Основная функция
func main() {
	rand.Seed(today.UnixNano())

	drawHeroCard()
	drawProgressStats()
	drawDailyEvents()
	drawPersonalGrowth()
	drawAchievements()
	drawFutureProspects()
	drawFooter()
}

// 🎨 Визуальные блоки вывода
func drawHeroCard() {
	fmt.Println("\n🚀 100daysGo: HARD CORE 🚀")
	fmt.Println("══════════════════════════════════════════════")
	fmt.Printf("👤 Имя: Гоша | Возраст: 37 | Нейротип: СДВГ+ОКР+социофоб\n")
	fmt.Printf("🎯 Миссия: Из курьера в Golang-разработчика за 100 дней\n")
	fmt.Printf("📅 Сегодня: %s | Day%d челленджа\n", today.Format("02.01.2006"), currentDay)
	fmt.Printf("📚 Тема дня: Numeric Types - Boolean\n")
}

func drawProgressStats() {
	fmt.Printf("\n🔥 ПРОГРЕСС Day%d/%d (%.0f%%)\n", currentDay, 100, stats.ProgressPercent)
	fmt.Println(generateProgressBar(stats.ProgressPercent, 25))

	fmt.Printf("🏆 Уровень: %s (Lvl %d | %d/%d XP)\n",
		getDevLevel(currentDay), stats.Level, stats.Experience, stats.NextLevelXP)
	fmt.Printf("💪 Сила: %s | 🧠 Состояние: %s | 😄 Настроение: %s\n",
		stats.WillpowerLevel, stats.MentalState, stats.CurrentMood)
	fmt.Printf("💡 Сила кода: %d | 🔥 Серия: %d дней\n", stats.CodingPower, stats.Streak)
}

func drawDailyEvents() {
	fmt.Printf("\n⚡ СЕГОДНЯ:\n")
	fmt.Printf("   💬 Мама: \"%s\"\n", momQuote)
	fmt.Printf("   🧙‍♂️ Совет ментора: \"%s\" %s\n", mentorAdvice.Message, mentorAdvice.Emoji)
	fmt.Printf("   💫 Мотивация: %s\n", motivation)
	fmt.Printf("   🎲 Факт о Go: %s\n", factOfTheDay)

	fmt.Println("\n🎲 СЛУЧАЙНЫЕ СОБЫТИЯ:")
	for _, e := range dailyEvents {
		fmt.Printf("%s %s\n", getEventEmoji(e.Type), e.Description)
	}
}

func drawPersonalGrowth() {
	fmt.Printf("\n🌱 ЛИЧНЫЙ РОСТ:\n")
	fmt.Printf("   🎮 Пропущено игр: %d сессий\n", growth.GamingSkipped)
	fmt.Printf("   💻 Написано кода: %d строк\n", growth.CodeLines)
	fmt.Printf("   📚 Часов обучения: %.1f\n", growth.StudyHours)
	fmt.Printf("   🌍 Часов в реальной жизни: %d\n", growth.RealLifeHours)
	fmt.Printf("   😌 Уверенность: %d/100 | 😨 Тревожность: %d/100\n",
		growth.ConfidenceLevel, 100-growth.SocialEnergy)
	fmt.Printf("   👵 Гордость мамы: %d/100 | 🆘 Дней без паники: %d\n",
		growth.MomPrideLevel, growth.DaysWithoutPanic)
}

func drawAchievements() {
	fmt.Printf("\n🏆 ДОСТИЖЕНИЯ: %d/%d разблокировано\n", unlockedAchvs, len(achievements))
	for _, a := range achievements {
		if a.isUnlocked(currentDay) {
			fmt.Printf("   %s %s: %s\n", getRarityEmoji(a.Type), a.Name, a.Description)
		}
	}

	fmt.Printf("\n📜 КВЕСТЫ: %d активно\n", activeQuests)
	for _, q := range quests {
		if q.isActive(currentDay) && !q.Completed {
			fmt.Printf("   ➤ %s: %s\n", q.Name, q.Description)
		}
	}
}

func drawFutureProspects() {
	fmt.Printf("\n💰 ПЕРСПЕКТИВЫ:\n")
	fmt.Printf("   💸 Текущая ЗП: ~%d руб/мес → %d руб/мес через %d дней\n",
		80000+currentDay*500, 80000+100*1200, stats.DaysRemaining)
	fmt.Printf("   🏡 Через %d дней: квартира у метро\n", max(0, 100-currentDay))
	fmt.Printf("   👨‍👩‍👧 Через %d дней: семья гордится тобой\n", max(0, 80-currentDay))
}

func drawFooter() {
	fmt.Println("\n══════════════════════════════════════════════")
	fmt.Println("💡 ФИЛОСОФИЯ ДНЯ:")
	fmt.Println("   \"Boolean — это не true/false. Это твой выбор: сдаваться или идти вперёд.\"")
	fmt.Printf("   👵 Мама: \"Ну ладно, я вижу ты стараешься... может, через год купишь мне дачу?\"\n")
}

func generateProgressBar(percent float64, width int) string {
	filled := int(percent/100*float64(width) + 0.5)
	empty := width - filled

	var bar strings.Builder
	for i := 0; i < filled; i++ {
		bar.WriteString("🟩")
	}
	for i := 0; i < empty; i++ {
		bar.WriteString("⬜")
	}
	return bar.String()
}

func getEventEmoji(t string) string {
	switch t {
	case "obstacle":
		return "🚧"
	case "victory":
		return "🎉"
	case "challenge":
		return "⚔️"
	case "quest":
		return "📜"
	}
	return "❓"
}

func getRarityEmoji(t string) string {
	switch t {
	case "common":
		return "⚪"
	case "rare":
		return "🔵"
	case "epic":
		return "🟣"
	case "legendary":
		return "🟡"
	}
	return "❓"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 🧩 Сокращённые структуры данных
type ChallengeStats struct {
	DaysCompleted   int
	DaysRemaining   int
	ProgressPercent float64
	Level           int
	Experience      int
	NextLevelXP     int
	WillpowerLevel  string
	MentalState     string
	CurrentMood     string
	CodingPower     int
	Streak          int
	MaxStreak       int
}

type PersonalGrowth struct {
	GamingSkipped    int
	StudyHours       float64
	CodeLines        int
	ConfidenceLevel  int
	StressLevel      int
	SocialEnergy     int
	MomPrideLevel    int
	RealLifeHours    int
	DaysWithoutPanic int
}

type DailyEvent struct {
	Type        string
	Description string
}

type Achievement struct {
	Name        string
	Description string
	Type        string
	DayUnlock   int
}

func (a *Achievement) isUnlocked(day int) bool {
	return day >= a.DayUnlock
}

type Quest struct {
	Name        string
	Description string
	DayStart    int
	Completed   bool
}

func (q *Quest) isActive(day int) bool {
	return day >= q.DayStart && !q.Completed
}

type MentorAdvice struct {
	Message string
	Emoji   string
}

// 🎪 Данные челленджа
var (
	momQuotes = []string{
		"Опять за компом? Соседский Коля уже вторую машину купил!",
		"37 лет, а всё в компьютерные игры играешь!",
		"Может, лучше бы пошёл и развёз пару заказов?",
		"Когда уже станешь нормальным женатым мужичком?",
		"Ну хоть бы жену нашёл, как все!",
	}

	obstacles = []string{
		"Сосед сверлит стену во время изучения замыканий",
		"Кошка прошлась по клавиатуре и закоммитила",
		"Мама требует вынести мусор во время дебага",
		"Интернет отключился в самый важный момент",
		"СДВГ: начал изучать интерфейсы, переключился на каналы",
	}

	victories = []string{
		"Победил панику при виде error handling!",
		"Написал первую горутину без deadlock!",
		"Починил баг одним символом после 3 часов поиска!",
		"Понял разницу между slice и array без гугления!",
		"Рефакторинг прошёл успешно — ничего не сломал!",
	}

	goFacts = []string{
		"Go создан тремя легендарными программистами: Роб Пайк, Кен Томпсон, Роберт Гризмер",
		"Горутины легче потоков ОС — их могут быть миллионы!",
		"Девиз Go: 'Do not communicate by sharing memory; instead, share memory by communicating'",
		"Go может компилироваться в WebAssembly!",
		"Go формат кода автоматически применяется через gofmt",
	}

	adviceList = []MentorAdvice{
		{"Не бойся ошибок — они твои лучшие учителя", "📚"},
		{"СДВГ — это не проклятие, а суперсила в программировании", "⚡"},
		{"Каждая строка кода — это кирпичик в твою карьеру", "🧱"},
		{"Ты сильнее своих зависимостей — докажи это!", "💪"},
		{"ОКР помогает писать чистый, структурированный код", "🧼"},
	}

	motivations = []string{
		"Каждая строка кода — это шаг от зависимости к свободе!",
		"Сегодня ты стал на день ближе к карьере мечты!",
		"СДВГ и ОКР — твои суперсилы в программировании!",
		"Игры украли прошлое, Go вернёт будущее!",
		"37 лет — идеальный возраст для перезагрузки!",
	}

	achievements = []Achievement{
		{"Первый день", "Выжил после первого дня", "common", 1},
		{"Неделя без срывов", "7 дней кода подряд", "common", 7},
		{"Месяц без игр", "30 дней без игр", "rare", 30},
		{"Полпути", "50 дней пройдено", "epic", 50},
		{"Самурай кода", "100 дней без срывов", "legendary", 100},
	}

	quests = []Quest{
		{"День 1", "Написать первую программу", 1, false},
		{"День 10", "Создать структуру и методы", 10, false},
		{"День 20", "Создать HTTP-сервер", 20, false},
		{"День 30", "Написать тесты", 30, false},
		{"День 50", "Создать CLI-инструмент", 50, false},
		{"День 100", "Запустить проект в продакшн", 100, false},
	}
)

// 🧮 Генераторы данных
func generateDailyEvents(count int) []DailyEvent {
	events := make([]DailyEvent, 0, count)
	types := []string{"obstacle", "victory", "challenge", "quest"}

	for i := 0; i < count; i++ {
		t := types[rand.Intn(len(types))]
		desc := ""

		switch t {
		case "obstacle":
			desc = obstacles[rand.Intn(len(obstacles))]
		case "victory":
			desc = victories[rand.Intn(len(victories))]
		case "challenge":
			desc = "Вызов: " + []string{"Написать Hello World", "Создать функцию сложения", "Разобраться с указателями", "Написать тест"}[rand.Intn(4)]
		case "quest":
			desc = "Квест: " + []string{"Прочитать доку", "Написать 50 строк кода", "Создать репозиторий", "Написать README"}[rand.Intn(4)]
		}

		events = append(events, DailyEvent{t, desc})
	}
	return events
}

// 🧠 Уровни и состояния
func getWillpowerLevel(day int) string {
	levels := []string{"Стеклянный", "Бумажный", "Картонный", "Деревянный", "Железный", "Стальной", "Алмазный"}
	return levels[min(day/15, len(levels)-1)]
}

func getMentalState(day int) string {
	states := []string{
		"Паника и отрицание", "Гнев на компилятор", "Торг с собой",
		"Депрессия от ошибок", "Принятие и просветление", "Поток и продуктивность",
	}
	return states[min(day/20, len(states)-1)]
}

func getCurrentMood(day int) string {
	moods := []string{
		"Ожидание старта", "Энтузиазм", "Формирование привычки",
		"Стабильный прогресс", "Преодоление трудностей", "Уверенность",
	}
	return moods[min(day/17, len(moods)-1)]
}

func getDevLevel(day int) string {
	levels := []string{
		"Новичок 🐣", "Ученик 📚", "Интерн 🔧", "Junior 💻",
		"Middle 🚀", "Senior 🏆", "Гуру 🧙", "Легенда 🌟",
	}
	return levels[min(day/15, len(levels)-1)]
}

// 🎲 Рандомайзеры
func getRandomQuote(quotes []string) string {
	return quotes[rand.Intn(len(quotes))]
}

func getRandomAdvice() MentorAdvice {
	return adviceList[rand.Intn(len(adviceList))]
}

func getRandomFact() string {
	return goFacts[rand.Intn(len(goFacts))]
}

func getDailyMotivation() string {
	return motivations[rand.Intn(len(motivations))]
}

func countUnlockedAchievements() int {
	count := 0
	for _, a := range achievements {
		if a.isUnlocked(currentDay) {
			count++
		}
	}
	return count
}

func countActiveQuests() int {
	count := 0
	for _, q := range quests {
		if q.isActive(currentDay) {
			count++
		}
	}
	return count
}
