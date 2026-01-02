package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	hundredDaysStart = "2025-11-03"
	go365Start       = "2026-01-01"
	hundredDaysTotal = 100
	go365TotalDays   = 365
	maxLevelXP       = 1000
	codeLinesPerDay  = 42.5
	deletedGames     = 7
)

type Person struct {
	Name, Background, Goal string
	Age                    int
}

type Progress struct {
	HundredDaysCount, HundredDaysXP, HundredDaysLevel int
	Go365DaysCount, Go365XP, Go365Level               int
	CodeLines                                         float64
}

type Achievement struct {
	Emoji, Name, Desc string
	Unlocked          bool
}

type App struct {
	gosha        Person
	currentDate  time.Time
	progress     Progress
	theme        string
	rng          *rand.Rand
	motivations  []string
	achievements []Achievement
	dailyThemes  []string
	dailyEvents  []string
}

func NewApp() *App {
	now := time.Now()
	hundredDays := calculateDaysSince(hundredDaysStart)
	go365Days := max(0, calculateDaysSince(go365Start))

	return &App{
		gosha: Person{
			Name:       "Гоша",
			Age:        38,
			Background: "Бывший игроман с опытом метаний между Python/Java/C#/C++/JS",
			Goal:       "Стать Junior Go-разработчиком в 2026. Никаких переключений!",
		},
		currentDate: now,
		progress: Progress{
			HundredDaysCount: hundredDays,
			HundredDaysXP:    min(hundredDays*10, hundredDaysTotal*10),
			HundredDaysLevel: 1 + hundredDays*10/maxLevelXP,
			Go365DaysCount:   go365Days,
			Go365XP:          go365Days * 15,
			Go365Level:       1 + go365Days*15/maxLevelXP,
			CodeLines:        float64(hundredDays+go365Days) * codeLinesPerDay,
		},
		theme: "2026: Глубина вместо ширины. Только Go",
		rng:   rand.New(rand.NewPCG(uint64(now.UnixNano()), uint64(now.Unix()))),
		motivations: []string{
			"В 2025 ты прыгал между Python и Java. В 2026 ты прыгаешь только по уровням в Go.",
			"Каждая строчка кода на Go — шаг к новой профессии. Никаких отступлений!",
			"Твой GTX 1060 больше не рендерит Unreal Engine — он компилирует твоё будущее в Go!",
			"Гофер внутри тебя голоден. Накорми его строчками кода, а не FPS в играх.",
			"Вчера ты был курьером. Сегодня ты — программист. Завтра — Go-разработчик.",
			"Интерфейсы в Go — твои друзья. Они не спрашивают твоё имя, только методы.",
			"Сборщик мусора убирает память. Ты убираешь сомнения. Доверяй runtime.",
			"Методы определяют твою сущность. Функции — только действия. Ты — метод, Гоша.",
		},
		achievements: []Achievement{
			{"🔥", "Фокус-2026", "Первый день без игр и сериалов. Только Go.", false},
			{"🚀", "Двойной челлендж", "100daysGo + Go365 = непрерывный рост", false},
			{"🎯", "Хардкорный выбор", "Удалены Unity, IntelliJ, Unreal Engine. Только VS Code + Go", false},
			{"🐍➡️🐹", "От Змеи к Гоферу", "Полный переход с Python на Go. Символично!", false},
			{"💻", "GTX 1060 Upgrade", "Видеокарта теперь майнит знания, а не FPS", false},
		},
		dailyThemes: []string{
			"Почему фокус на Go — твой последний шанс",
			"Как Go спасёт тебя от распыления",
			"Глубина вместо ширины: Путь к мастерству в Go",
			"Почему только Go — ключ к твоему будущему",
			"Как один язык превратит тебя в профессионала",
			"Сборщик мусора в жизни: Убираем сомнения, как мусор в памяти",
			"Интерфейсы — не только в Go, но и в жизни: Не важно, кто ты, важно, что ты можешь",
		},
		dailyEvents: []string{
			"Утром: Удалил все игры с GTX 1060",
			"Днём: Написал первый коммит в Go365",
			"Вечером: Прошел 10к шагов по заснеженным улицам",
			"Ночью: Написал Telegram-бота для учёта расходов",
			"Вечером: Прочитал главу про interfaces в Effective Go",
			"Утром: Запустил Go-сервер для учёта прогресса",
			"Днём: Решил 5 задач на LeetCode на Go",
			"Вечером: Написал свой первый middleware для Gin",
			"Ночью: Изучил основы gRPC и написал простой сервис",
		},
	}
}

func main() {
	app := NewApp()
	app.unlockAchievements()
	app.renderUI()
}

func (a *App) renderUI() {
	a.printHeader()
	a.printProgressSection()
	a.printDailyInsight()
	a.printStatsSection()
	a.printAchievementsSection()
	a.printFutureSection()
	a.printFooter()
	a.interactiveCheck()
}

// --- СЕКЦИИ ИНТЕРФЕЙСА ---

func (a *App) printHeader() {
	a.printTitle("🔥 2026: ГОД ФОКУСА НА GO | 100daysGo + Go365 🔥", "33")
	a.printLine("═", 70)
	a.printfColored("👤 %s | %d лет | %s\n", "36", a.gosha.Name, a.gosha.Age, a.gosha.Background)
	a.printfColored("🎯 %s\n", "32", a.gosha.Goal)
	a.printf("📅 %s | 100daysGo: День %d/%d | Go365: День %d/%d\n",
		a.currentDate.Format("02.01.2006"),
		a.progress.HundredDaysCount, hundredDaysTotal,
		a.progress.Go365DaysCount, go365TotalDays)
	a.printfColored("📚 Тема дня: %s\n", "34", a.theme)
}

func (a *App) printProgressSection() {
	a.printSectionHeader("🚀 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ", "34")

	hundredDaysPercent := a.progress.HundredDaysCount * 100 / hundredDaysTotal
	go365Percent := a.progress.Go365DaysCount * 100 / go365TotalDays

	a.printfColored("▸ 100daysGo: %.0f%% завершено | Уровень: %d | XP: %d/%d\n", "36",
		float64(hundredDaysPercent),
		a.progress.HundredDaysLevel,
		a.progress.HundredDaysXP,
		hundredDaysTotal*10)
	a.printProgressBar(hundredDaysPercent)

	a.printfColored("▸ Go365: %.1f%% завершено | Уровень: %d | XP: %d/%d\n", "32",
		float64(go365Percent),
		a.progress.Go365Level,
		a.progress.Go365XP,
		go365TotalDays*15)
	a.printProgressBar(go365Percent)
}

func (a *App) printDailyInsight() {
	// Формат даты: "02 ЯНВАРЯ 2026"
	monthNames := []string{"ЯНВАРЯ", "ФЕВРАЛЯ", "МАРТА", "АПРЕЛЯ", "МАЯ", "ИЮНЯ",
		"ИЮЛЯ", "АВГУСТА", "СЕНТЯБРЯ", "ОКТЯБРЯ", "НОЯБРЯ", "ДЕКАБРЯ"}
	dateLabel := fmt.Sprintf("%02d %s %d",
		a.currentDate.Day(),
		monthNames[a.currentDate.Month()-1],
		a.currentDate.Year())

	theme := a.getRandomItem(a.dailyThemes)
	motivation := a.getRandomItem(a.motivations)
	events := a.getRandomItems(a.dailyEvents, a.rng.IntN(3)+1)

	a.printSectionHeader(fmt.Sprintf("💡 СУТЬ %s: %s", dateLabel, theme), "31;1")
	a.printBlock(56, func() {
		a.printf("❌ ПРОШЛОЕ (2023-2025):\n")
		a.printBullet("Январь 2025: Python (Год Змеи) → Май: Переключение на Go")
		a.printBullet("Unity (C#) → Unreal Engine (C++) → IntelliJ (Java) → VS Code (JS)")
		a.printBullet("GTX 1060 тонула в лаве Unreal Engine 5, а не в компиляции Go")
		a.printBullet("10 лет распыления вместо глубины")

		a.printf("\n✅ НАСТОЯЩЕЕ (%s):\n", a.currentDate.Format("02.01.2006"))
		for _, event := range events {
			a.printBullet(event)
		}
	})

	a.printSectionHeader("✨ МОТИВАЦИЯ ДНЯ", "35")
	a.printf("💬 %s\n", motivation)
}

func (a *App) printStatsSection() {
	totalDays := a.progress.HundredDaysCount + a.progress.Go365DaysCount
	learningHours := float64(totalDays) * 2.5
	freedomHours := float64(deletedGames) * 3.0

	a.printSectionHeader("📊 СТАТИСТИКА ПРЕВРАЩЕНИЯ", "36")
	a.printBullet(fmt.Sprintf("Удалено игр: %d (освобождено %.1f часов/день)", deletedGames, freedomHours))
	a.printBullet(fmt.Sprintf("Написано строк кода: %.0f (100daysGo + Go365)", a.progress.CodeLines))
	a.printBullet(fmt.Sprintf("Часов на обучение: %.1f | Среднее: 2.5 часа/день", learningHours))
	a.printBullet("Репозиториев: 2 (100daysGo + Go365/Go1)")
	a.printBullet("Заблокировано: Unity Hub, IntelliJ IDEA, Unreal Engine Launcher")
}

func (a *App) printAchievementsSection() {
	unlocked := countUnlocked(a.achievements)
	a.printSectionHeader(fmt.Sprintf("🏆 ДОСТИЖЕНИЯ (%d/%d)", unlocked, len(a.achievements)), "33")

	for _, ach := range a.achievements {
		status := "🔒"
		color := "37" // Серый
		if ach.Unlocked {
			status = "✅"
			color = "32" // Зеленый
		}
		a.printfColored("   %s %s: %s\n", color, status, ach.Name, ach.Desc)
	}
}

func (a *App) printFutureSection() {
	currentSalary := 120000 + 1800*(a.progress.HundredDaysCount+a.progress.Go365DaysCount)

	a.printSectionHeader("🔮 БУДУЩЕЕ ПОСЛЕ 2026", "35")
	a.printf("💼 Go-разработчик: %s%d ₽/мес → %d ₽/мес%s (через год)\n",
		ansi("31;1"), currentSalary, 350000, ansi("0"))
	a.printBullet("Карьера: Junior (сейчас) → Middle (июнь 2028) → Senior (декабрь 2029)")
	a.printBullet("Свобода: Работа из любой точки мира. Больше нет сугробов и луж!")
	a.printBullet("GTX 1060: Теперь греет Docker-контейнеры с Go-кодом")
	a.printBullet(fmt.Sprintf("Финал 100daysGo: %d дней | До Senior: %d дней",
		hundredDaysTotal-a.progress.HundredDaysCount,
		go365TotalDays-a.progress.Go365DaysCount))
}

func (a *App) printFooter() {
	a.printLine("═", 70)
	a.printSectionHeader("💬 АЙТИ-НАСТРОЙ ГОШИ НА 2026 ГОД", "34")
	a.printBullet("Больше никаких 'попробую C#' или 'вдруг Unity'!")
	a.printBullet("Каждый день — 1 коммит в Go365. Каждая строка — шаг к свободе.")
	a.printBullet("Мой Гофер сильнее всех боссов в играх. Его оружие — goroutines и channels.")

	a.printSectionHeader(fmt.Sprintf("🎉 %s: ПРОГРАММИСТСКИЙ ДЕНЬ", a.currentDate.Format("02.01.2006")), "33")
	for _, event := range a.getRandomItems(a.dailyEvents, 3) {
		a.printBullet(event)
	}
	a.printf("\n%s🚀 ПОМНИ: В IT ценится глубина, а не количество языков. Продолжай копать!%s\n",
		ansi("35;1"), ansi("0"))
}

// --- ВСПОМОГАТЕЛЬНЫЕ МЕТОДЫ ---

func (a *App) unlockAchievements() {
	a.achievements[0].Unlocked = a.progress.Go365DaysCount >= 1
	a.achievements[1].Unlocked = a.progress.HundredDaysCount > 0 && a.progress.Go365DaysCount > 0
	a.achievements[2].Unlocked = a.progress.Go365DaysCount >= 3
	a.achievements[3].Unlocked = a.progress.HundredDaysCount > 50 && a.progress.Go365DaysCount > 0
	a.achievements[4].Unlocked = deletedGames > 0
}

func (a *App) interactiveCheck() {
	a.printLine("═", 70)
	a.printSectionHeader("🔍 Проверить прогресс", "36")

	fmt.Println("   - Для 100daysGo: введите день (например: 25)")
	fmt.Println("   - Для Go365: введите дату (например: 2026-01-01)")
	fmt.Print("   Ваш выбор: ")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return
	}

	dirPath := a.getProgressPath(input)
	lines, err := countCodeLines(dirPath)
	if err != nil {
		a.printfColored("❌ Ошибка: %v\n", "31", err)
		return
	}

	emoji := "✅"
	if lines > 100 {
		emoji = "🔥"
	} else if lines < 10 {
		emoji = "💪"
	}

	a.printfColored("\n%s Прогресс за %s: %.0f строк кода!\n", "32;1", emoji, input, lines)
	if lines > 0 {
		a.printfColored("💡 Совет: Добавь тесты и документацию!\n", "34;1")
	}
}

func (a *App) getProgressPath(input string) string {
	if strings.Contains(input, "-") {
		return filepath.Join("..", "Go365", input)
	}
	return filepath.Join("..", fmt.Sprintf("day%s", input))
}

// --- УНИВЕРСАЛЬНЫЕ УТИЛИТЫ ---

func (a *App) getRandomItem(items []string) string {
	return items[a.rng.IntN(len(items))]
}

func (a *App) getRandomItems(items []string, count int) []string {
	result := make([]string, 0, count)
	used := make(map[int]bool)

	for len(result) < count && len(result) < len(items) {
		idx := a.rng.IntN(len(items))
		if !used[idx] {
			used[idx] = true
			result = append(result, items[idx])
		}
	}
	return result
}

func countUnlocked(achievements []Achievement) int {
	count := 0
	for _, a := range achievements {
		if a.Unlocked {
			count++
		}
	}
	return count
}

func calculateDaysSince(dateStr string) int {
	t, _ := time.Parse("2006-01-02", dateStr)
	return int(time.Since(t).Hours() / 24)
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !isCodeFile(path) {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "#") {
				total++
			}
		}
		return scanner.Err()
	})
	return total, err
}

func isCodeFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".go" || ext == ".md"
}

// --- ФОРМАТТЕРЫ И ЦВЕТА ---

func (a *App) printTitle(text, color string) {
	fmt.Printf("%s%s%s\n", ansi(color+";1"), text, ansi("0"))
}

func (a *App) printSectionHeader(text, color string) {
	fmt.Printf("\n%s%s%s\n", ansi(color+";1"), text, ansi("0"))
}

func (a *App) printLine(char string, count int) {
	fmt.Println(strings.Repeat(char, count))
}

func (a *App) printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

func (a *App) printfColored(format, color string, args ...any) {
	fmt.Printf("%s"+format+"%s", append([]any{ansi(color)}, append(args, ansi("0"))...)...)
}

func (a *App) printBlock(width int, content func()) {
	fmt.Println("   ┌" + strings.Repeat("─", width) + "┐")
	content()
	fmt.Println("   └" + strings.Repeat("─", width) + "┘")
}

func (a *App) printBullet(text string) {
	fmt.Printf("   │   - %s\n", text)
}

func (a *App) printProgressBar(percent int) {
	width := 50
	filled := percent * width / 100
	fmt.Printf("[%s%s] %d%%\n",
		strings.Repeat("█", filled),
		strings.Repeat("░", width-filled),
		percent)
}

func ansi(code string) string {
	return "\033[" + code + "m"
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
