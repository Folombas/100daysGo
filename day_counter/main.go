package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	hundredDaysStart = "2025-11-03" // Начало 100daysGo
	go365Start       = "2026-01-01" // Начало Go365
	hundredDaysTotal = 100
	go365TotalDays   = 365
	maxLevelXP       = 1000
	codeLinesPerDay  = 42.5
	deletedGames     = 7 // Количество удалённых игр на старте 2026
)

type Person struct {
	Name       string
	Age        int
	Background string
	Goal       string
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
	prng         *rand.Rand
	motivations  []string
	achievements []Achievement
}

func NewApp() *App {
	now := time.Now()
	hundredDays := daysSince(hundredDaysStart)
	go365Days := daysSince(go365Start)
	if go365Days < 0 {
		go365Days = 0
	}

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

			Go365DaysCount: go365Days,
			Go365XP:        go365Days * 15,
			Go365Level:     1 + go365Days*15/maxLevelXP,

			CodeLines: float64(hundredDays+go365Days) * codeLinesPerDay,
		},
		theme: "2026: Глубина вместо ширины. Только Go - Value Receivers",
		prng:  rand.New(rand.NewSource(now.UnixNano())),
		motivations: []string{
			"Твой GTX 1060 больше не рендерит Unreal Engine — он компилирует твоё будущее в Go!",
			"20 лет распыления закончились. Сегодня ты удалил 7 игр. Каждый день — ещё одна игра вместо кода.",
			"В 2025 ты прыгал между Python и Java. В 2026 ты прыгаешь только по уровням в Go.",
			"Гофер внутри тебя голоден. Накорми его строчками кода, а не FPS в играх.",
			"Каждый коммит в Go365 — это кирпич в фундаменте твоей новой профессии.",
			"Не 10 языков поверхностно. Не 10 движков. Только Go. Глубоко. Серьёзно. До победного.",
			"Твой рэп научил тебя ритму. Теперь найди ритм в goroutines и channels.",
		},
		achievements: []Achievement{
			{"🔥", "Фокус-2026", "Первый день без игр и сериалов. Только Go.", false},
			{"🚀", "Двойной челлендж", "100daysGo + Go365 = непрерывный рост", false},
			{"🎯", "Хардкорный выбор", "Удалены Unity, IntelliJ, Unreal Engine. Только VS Code + Go", false},
			{"🐍➡️🐹", "От Змеи к Гоферу", "Полный переход с Python на Go. Символично!", false},
			{"💻", "GTX 1060 Upgrade", "Видеокарта теперь майнит знания, а не FPS", false},
		},
	}
}

func main() {
	app := NewApp()
	app.unlockAchievements()
	app.printHeader()
	app.printProgress()
	app.printToday()
	app.printStats()
	app.printFuture()
	app.printFooter()
	app.interactiveLineCounter()
}

func daysSince(dateStr string) int {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(fmt.Sprintf("invalid start date: %v", err))
	}
	return int(time.Since(t).Hours() / 24)
}

func (a *App) unlockAchievements() {
	// Автоматически разблокируем достижения на основе прогресса
	if a.progress.Go365DaysCount >= 1 {
		a.achievements[0].Unlocked = true // Фокус-2026
	}
	if a.progress.HundredDaysCount > 0 && a.progress.Go365DaysCount > 0 {
		a.achievements[1].Unlocked = true // Двойной челлендж
	}
	if a.progress.Go365DaysCount >= 3 {
		a.achievements[2].Unlocked = true // Хардкорный выбор
	}
	if a.progress.HundredDaysCount > 50 && a.progress.Go365DaysCount > 0 {
		a.achievements[3].Unlocked = true // От Змеи к Гоферу
	}
	if deletedGames > 0 {
		a.achievements[4].Unlocked = true // GTX 1060 Upgrade
	}
}

func (a *App) printHeader() {
	fmt.Printf("\n%s🔥 2026: ГОД ФОКУСА НА GO | 100daysGo + Go365 🔥%s\n",
		ansi("1;33"), ansi("0"))
	fmt.Println(strings.Repeat("═", 70))
	fmt.Printf("👤 %s%s%s | %d лет | %s\n",
		ansi("1;36"), a.gosha.Name, ansi("0"), a.gosha.Age, a.gosha.Background)
	fmt.Printf("🎯 %s%s%s\n",
		ansi("1;32"), a.gosha.Goal, ansi("0"))
	fmt.Printf("📅 %s | 100daysGo: День %d/%d | Go365: День %d/%d\n",
		a.currentDate.Format("02.01.2006"),
		a.progress.HundredDaysCount, hundredDaysTotal,
		a.progress.Go365DaysCount, go365TotalDays)
	fmt.Printf("📚 Тема дня: %s%s%s\n", ansi("1;34"), a.theme, ansi("0"))
}

func (a *App) printProgress() {
	hundredDaysPercent := float64(a.progress.HundredDaysCount) / hundredDaysTotal * 100
	go365Percent := float64(a.progress.Go365DaysCount) / go365TotalDays * 100

	fmt.Printf("\n%s🚀 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ:%s\n", ansi("1;34"), ansi("0"))

	// Прогресс 100daysGo
	fmt.Printf("%s▸ 100daysGo:%s %.0f%% завершено | Уровень: %d | XP: %d/%d\n",
		ansi("1;36"), ansi("0"),
		hundredDaysPercent,
		a.progress.HundredDaysLevel,
		a.progress.HundredDaysXP,
		hundredDaysTotal*10)
	fmt.Println(progressBar(hundredDaysPercent, 50))

	// Прогресс Go365
	fmt.Printf("%s▸ Go365:%s %.1f%% завершено | Уровень: %d | XP: %d/%d\n",
		ansi("1;32"), ansi("0"),
		go365Percent,
		a.progress.Go365Level,
		a.progress.Go365XP,
		go365TotalDays*15)
	fmt.Println(progressBar(go365Percent, 50))
}

func (a *App) printToday() {
	fmt.Printf("\n%s💡 СУТЬ 1 ЯНВАРЯ 2026:%s Почему фокус на Go — твой последний шанс%s\n",
		ansi("1;31"), ansi("1;33"), ansi("0"))
	fmt.Println("   ┌──────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("   │ ❌ ПРОШЛОЕ (2023-2025):                                                      │")
	fmt.Println("   │   - Январь 2025: Python (Год Змеи) → Май: Переключение на Go                 │")
	fmt.Println("   │   - Unity (C#) → Unreal Engine (C++) → IntelliJ (Java) → VS Code (JS)        │")
	fmt.Println("   │   - GTX 1060 тонула в лаве Unreal Engine 5, а не в компиляции Go             │")
	fmt.Println("   │   - 10 лет распыления вместо глубины                                         │")
	fmt.Println("   │                                                                              │")
	fmt.Println("   │ ✅ НАСТОЯЩЕЕ (01.01.2026):                                                   │")
	fmt.Println("   │   - 8:00 утра. Чай с вкусняшками. Новый день. Новый фокус.                   │")
	fmt.Println("   │   - Все игры удалены. Свободное время → Go365                                │")
	fmt.Println("   │   - Только один путь: от \"fmt.Println(hello)\" до Production-кода           │")
	fmt.Println("   │   - Гофер — мой персонаж. Каждый день — прокачка уровня!                     │")
	fmt.Println("   └──────────────────────────────────────────────────────────────────────────────┘")

	fmt.Printf("\n%s✨ МОТИВАЦИЯ ДНЯ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💬 %s\n", a.motivations[a.currentDate.YearDay()%len(a.motivations)])
}

func (a *App) printStats() {
	totalDays := a.progress.HundredDaysCount + a.progress.Go365DaysCount
	learningHours := float64(totalDays) * 2.5
	freedomHours := float64(deletedGames) * 3.0 // 3 часа на игру

	fmt.Printf("\n%s📊 СТАТИСТИКА ПРЕВРАЩЕНИЯ:%s\n", ansi("1;36"), ansi("0"))
	fmt.Printf("   🎮 Удалено игр: %d (освобождено %.1f часов/день)\n", deletedGames, freedomHours)
	fmt.Printf("   💻 Написано строк кода: %.0f (100daysGo + Go365)\n", a.progress.CodeLines)
	fmt.Printf("   ⏳ Часов на обучение: %.1f | Среднее: 2.5 часа/день\n", learningHours)
	fmt.Printf("   📁 Репозиториев: 2 (100daysGo + Go365/Go1)\n")
	fmt.Printf("   🚫 Заблокировано: Unity Hub, IntelliJ IDEA, Unreal Engine Launcher\n")
}

func (a *App) printAchievements() {
	unlocked := 0
	for _, ach := range a.achievements {
		if ach.Unlocked {
			unlocked++
		}
	}

	fmt.Printf("\n%s🏆 ДОСТИЖЕНИЯ (%d/%d):%s\n", ansi("1;33"), unlocked, len(a.achievements), ansi("0"))
	for _, ach := range a.achievements {
		status := "🔒"
		style := ansi("1;37") // Серый для закрытых
		if ach.Unlocked {
			status = "✅"
			style = ansi("1;32") // Зелёный для открытых
		}
		fmt.Printf("   %s%s %s: %s%s\n", style, status, ach.Name, ach.Desc, ansi("0"))
	}
}

func (a *App) printFuture() {
	// Расчёт зарплаты с учётом двух челленджей
	baseSalary := 120000
	salaryGrowth := 1800 * (a.progress.HundredDaysCount + a.progress.Go365DaysCount)
	currentSalary := baseSalary + salaryGrowth
	projectedSalary := 350000 // Прогноз через год

	fmt.Printf("\n%s🔮 БУДУЩЕЕ ПОСЛЕ 2026:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💼 Go-разработчик: %s%d ₽/мес → %d ₽/мес%s (через год)\n",
		ansi("1;31"), currentSalary, projectedSalary, ansi("0"))
	fmt.Printf("   📈 Карьера: Junior (сейчас) → Middle (июнь 2028) → Senior (декабрь 2029)\n")
	fmt.Printf("   🏠 Свобода: Работа из любой точки мира. Больше нет сугробов и луж!\n")
	fmt.Printf("   🎮 GTX 1060: Теперь греет не игровые сцены, а Docker-контейнеры с Go-кодом\n")
	fmt.Printf("   ⏳ Финал 100daysGo: %d дней | Старт Go365: %d дней до Senior\n",
		hundredDaysTotal-a.progress.HundredDaysCount,
		go365TotalDays-a.progress.Go365DaysCount)
}

func (a *App) printFooter() {
	fmt.Println(strings.Repeat("═", 70))
	fmt.Printf("%s💬 КЛЯТВА ГОШИ НА 2026 ГОД:%s\n", ansi("1;34"), ansi("0"))
	fmt.Println("   \"Больше никаких 'попробую C#' или 'вдруг Unity'!\"")
	fmt.Println("   \"Каждый день — 1 коммит в Go365. Каждая строка — шаг к свободе.\"")
	fmt.Println("   \"Мой Гофер сильнее всех боссов в играх. Его оружие — goroutines и channels.\"")

	fmt.Printf("\n%s🎉 01.01.2026: ИСТОРИЧЕСКИЙ ДЕНЬ%s\n", ansi("1;33"), ansi("0"))
	fmt.Println("   - Утром: Удалены все игры с GTX 1060")
	fmt.Println("   - Днём: Запущен челлендж Go365")
	fmt.Println("   - Вечером: Написан первый коммит в Go365/Go1")
	fmt.Printf("   - Итог: %sСФОКУСИРОВАН. СОБРАН. ГОТОВ%s\n", ansi("1;32"), ansi("0"))

	fmt.Printf("\n%s🚀 СЛЕДУЮЩИЙ УРОВЕНЬ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   День 2 задача: Реализовать REST API для Go365-дневника")
}

func (a *App) interactiveLineCounter() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Printf("%s🔍 Проверить прогресс:%s\n", ansi("1;36"), ansi("0"))
	fmt.Println("   - Для 100daysGo: введите день (например: 25)")
	fmt.Println("   - Для Go365: введите дату (например: 2026-01-01)")
	fmt.Print("   Ваш выбор: ")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return
	}

	var dirPath string
	if strings.Contains(input, "-") {
		// Go365 формат: 2026-01-01
		dirPath = filepath.Join("..", "Go365", input)
	} else {
		// 100daysGo формат: day25
		dirPath = filepath.Join("..", fmt.Sprintf("day%s", input))
	}

	lines, err := countCodeLines(dirPath)
	if err != nil {
		fmt.Printf("%s❌ Ошибка: %v%s\n", ansi("1;31"), err, ansi("0"))
		return
	}

	emoji := "✅"
	switch {
	case lines > 100:
		emoji = "🔥"
	case lines < 10:
		emoji = "💪"
	}

	fmt.Printf("\n%s%s Прогресс за %s: %.0f строк кода!%s\n",
		ansi("1;32"), emoji, input, lines, ansi("0"))

	if lines > 0 {
		fmt.Printf("%s💡 Совет:%s Добавь тесты и документацию!%s\n",
			ansi("1;34"), ansi("1;33"), ansi("0"))
		fmt.Printf("%s🚀 Напоминание:%s В IT ценится глубина, а не количество языков. Продолжай!%s\n",
			ansi("1;35"), ansi("1;36"), ansi("0"))
	}
}

// --- Вспомогательные функции (без изменений) ---
func progressBar(percent float64, width int) string {
	filled := int(percent/100*float64(width) + 0.5)
	return fmt.Sprintf("[%s%s] %.0f%%",
		strings.Repeat("█", filled),
		strings.Repeat("░", width-filled),
		percent)
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || (filepath.Ext(path) != ".go" && filepath.Ext(path) != ".md") {
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
			if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
				continue
			}
			total++
		}
		return scanner.Err()
	})
	return total, err
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
