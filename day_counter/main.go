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
	startDateStr     = "2025-11-03"
	challengeDays    = 100
	maxLevelXP       = 1000
	codeLinesPerDay  = 42.5
	cigsPerDay       = 15
	beerPerDay       = 3
	studyHoursPerDay = 2.5
	cigCost          = 15
	beerCost         = 120
	dailySavings     = 35
	baseSalary       = 80000
	salaryIncrease   = 1700
)

type Person struct {
	Name       string
	Age        int
	Background string
	Goal       string
}

type Progress struct {
	Days, XP, Level, Streak, Confidence int
	CodeLines                           float64
}

type Achievement struct {
	Emoji, Name, Desc string
	Day               int
}

type App struct {
	gosha        Person
	currentDay   int
	progress     Progress
	theme        string
	prng         *rand.Rand
	motivations  []string
	facts        []string
	achievements []Achievement
}

func NewApp() *App {
	currentDay := daysSince(startDateStr)
	progress := calculateProgress(currentDay)
	return &App{
		gosha: Person{
			Name:       "Гоша",
			Age:        38,
			Background: "Бывший рэпер-гламурщик",
			Goal:       "Стать Go-разработчиком",
		},
		currentDay: currentDay,
		progress:   progress,
		theme:      "Functions: Multiple Return Values",
		prng:       rand.New(rand.NewSource(time.Now().UnixNano())),
		motivations: []string{
			"38 лет — идеальный возраст для старта в IT. Твой опыт жизни — твоя суперсила!",
			"Каждая строка кода сегодня — это шаг к свободе от курьерской работы завтра.",
			"Ты не 'поздно начинаешь'. Ты начинаешь в идеальное время с багажом жизненного опыта.",
			"Твоя миссия — не просто выучить Go. Твоя миссия — доказать, что никогда не поздно менять жизнь.",
			"Помни: самые успешные разработчики начинали с нуля. Разница в том, что ты начал с опыта жизни.",
			"Твой рэп научил тебя ритму и структуре. Теперь примени это к коду!",
			"Каждый раз, когда хочется сдаться, вспомни: через год ты будешь жалеть, что НЕ продолжил сегодня.",
		},
		facts: []string{
			"В Go нет наследования классов, только композиция. Как и в жизни — составляй свою судьбу из лучших частей!",
			"Go создан в Google для решения реальных проблем. И твоя проблема — реальна и достойна решения.",
			"Средняя зарплата Go-разработчика в Москве — 220,000 ₽. Это твоя цель через 6 месяцев.",
			"Telegram, Docker, Kubernetes — все они частично написаны на Go. Твой код тоже может изменить мир.",
			"1 горутина в Go = 1 поток выполнения. 1 день твоего челленджа = 1 шаг к новой жизни.",
			"В Go есть поговорка: 'Меньше кода — меньше багов'. В жизни: 'Меньше тусовок — больше смысла'.",
		},
		achievements: []Achievement{
			{"🌱", "Новое начато!", "Первые 24 часа без пагубных привычек", 1},
			{"⚔️", "Цифровой Гуру", "7 дней чистого кода вместо сериалов", 7},
			{"💎", "Сердце чемпиона", "Ты прошел четверть пути! 25 дней перемен!", 25},
			{"🚀", "Наполовину к звёздам", "50 дней без оглядки назад — только вперёд!", 50},
			{"🌟", "Полный круг", "100 дней нового Гоши — легенда в мире кода!", 100},
		},
	}
}

func main() {
	app := NewApp()
	app.printHeader()
	app.printProgress()
	app.printToday()
	app.printStats()
	app.printAchievements()
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

func calculateProgress(days int) Progress {
	xp := 100 + days*10
	return Progress{
		Days:       days,
		XP:         xp,
		Level:      1 + xp/maxLevelXP,
		Streak:     days,
		Confidence: min(100, days*2),
		CodeLines:  float64(days) * codeLinesPerDay,
	}
}

func (a *App) printHeader() {
	fmt.Printf("\n%s🔥 100 ДНЕЙ GО С ГОШЕЙ: ОТ БЕЗУМНЫХ ТУСОВОК К ПРОДУМАННОМУ КОДУ 🔥%s\n",
		ansi("1;33"), ansi("0"))
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("👤 %s%s%s | %d лет | %s\n",
		ansi("1;36"), a.gosha.Name, ansi("0"), a.gosha.Age, a.gosha.Background)
	fmt.Printf("🎯 %s%s%s\n",
		ansi("1;32"), a.gosha.Goal, ansi("0"))
	fmt.Printf("📅 %s | День: %d/%d | Тема: %s\n",
		time.Now().Format("02.01.2006"), a.currentDay, challengeDays, a.theme)
}

func (a *App) printProgress() {
	percent := float64(a.currentDay) / challengeDays * 100
	xpNeeded := a.progress.Level * maxLevelXP

	fmt.Printf("\n%s🚀 ПРОГРЕСС: %.0f%% завершено (Дней с 30 ноября: %d)%s\n",
		ansi("1;34"), percent, a.currentDay, ansi("0"))
	fmt.Println(progressBar(percent, 40))
	fmt.Printf("🏆 Уровень %d (%d/%d XP) | 💪 Уверенность: %d%%\n",
		a.progress.Level, a.progress.XP, xpNeeded, a.progress.Confidence)
}

func (a *App) printToday() {
	fmt.Printf("\n%s✨ СЕГОДНЯ ГОВОРИТ СЕРДЦЕ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💬 %s\n", a.motivations[a.currentDay%len(a.motivations)])
	fmt.Printf("   💡 %s\n", a.facts[a.currentDay%len(a.facts)])
}

func (a *App) printStats() {
	cigarettes := a.currentDay * cigsPerDay
	beerBottles := a.currentDay * beerPerDay
	studyHours := float64(a.currentDay) * studyHoursPerDay
	moneySaved := float64(cigarettes)*cigCost + float64(beerBottles)*beerCost + float64(a.currentDay)*dailySavings

	fmt.Printf("\n%s📊 СТАТИСТИКА ПЕРЕРОЖДЕНИЯ:%s\n", ansi("1;36"), ansi("0"))
	fmt.Printf("   🚭 Пропущено сигарет: %d | 🍺 Бутылок пива: %d\n", cigarettes, beerBottles)
	fmt.Printf("   💻 Часов обучения: %.1f | 💰 Сэкономлено: %.0f ₽\n", studyHours, moneySaved)
	fmt.Printf("   📝 Написано строк кода: %.0f | 🔥 Удалено игр: %d\n",
		a.progress.CodeLines, 7+a.currentDay/5)
}

func (a *App) printAchievements() {
	unlocked := 0
	for _, ach := range a.achievements {
		if a.currentDay >= ach.Day {
			unlocked++
		}
	}

	fmt.Printf("\n%s🏆 ДОСТИЖЕНИЯ (%d/%d):%s\n", ansi("1;33"), unlocked, len(a.achievements), ansi("0"))
	for _, ach := range a.achievements {
		if a.currentDay >= ach.Day {
			fmt.Printf("   %s%s %s%s\n", ansi("1;32"), ach.Emoji, ach.Name, ansi("0"))
		}
	}
}

func (a *App) printFuture() {
	currentSalary := baseSalary + a.currentDay*salaryIncrease
	daysToJob := max(0, 45-a.currentDay)

	fmt.Printf("\n%s🔮 БУДУЩЕЕ ЧЕРЕЗ 100 ДНЕЙ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💼 Go-разработчик в Биг-Техе (Текущая: %d ₽/мес → %s250,000 ₽/мес%s)\n",
		currentSalary, ansi("1;32"), ansi("0"))
	fmt.Printf("   🏠 Своя квартира-студия в новом районе у метро (мечта с 30 ноября)\n")
	fmt.Printf("   👵 Родители гордятся тобой (а не тем, как ты 'прославился' в прошлом)\n")
	fmt.Printf("   ⏳ Работа найдется через %d дней. Ты справишься!\n", daysToJob)
}

func (a *App) printFooter() {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("%s💬 ФИЛОСОФИЯ 38-ЛЕТНЕГО ГОШИ:%s\n", ansi("1;34"), ansi("0"))
	fmt.Println("   \"Я не 'поздно начинаю'. Я начинаю в то время, когда другие сдаются.\"")
	fmt.Println("   \"Моя жизнь до 38 лет — мой 'legacy код'. Теперь я рефакторю свою судьбу.\"")
	fmt.Println("   \"Каждая буква 'G' в 'Go' означает: Growth, Goals, Glory.\"")

	birthdayMessage := "\n%s🎉 НАПОМИНАНИЕ: 30 ноября 2025 года тебе исполнилось 38 лет. " +
		"Это не конец молодости — это начало твоей самой важной главы!%s"
	fmt.Printf(birthdayMessage, ansi("1;33"), ansi("0"))

	fmt.Printf("\n%s🌟 СЕГОДНЯ: УДАЛИЛ 1 ИГРУ + НАПИСАЛ %.0f СТРОКИ КОДА! ТВОЙ ПУТЬ ПРОДОЛЖАЕТСЯ! 🌟%s\n",
		ansi("1;32"), a.progress.CodeLines, ansi("0"))
}

func progressBar(percent float64, width int) string {
	filled := int(percent/100*float64(width) + 0.5)
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func (a *App) interactiveLineCounter() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Print("🔍 Проверить прогресс за другой день (например: day25): ")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return
	}

	if !strings.HasPrefix(input, "day") {
		fmt.Printf("%s⚠️ Неправильный формат! Используй: day25%s\n", ansi("1;31"), ansi("0"))
		return
	}

	dirPath := filepath.Join("..", input)

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

	fmt.Printf("\n%s%s %s: %.0f строк кода!%s\n",
		ansi("1;32"), emoji, input, lines, ansi("0"))

	if lines > 0 {
		fmt.Printf("%s💡 Совет: Добавь комментарии и рефакторинг!%s\n", ansi("1;34"), ansi("0"))
	}
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".go" {
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
			if line == "" || strings.HasPrefix(line, "//") {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
