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
	startDateStr  = "2025-11-03"
	challengeDays = 100
	maxLevelXP    = 1000
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
	Rarity            string
}

var (
	gosha      = Person{"Гоша", 38, "Бывший рэпер-гламурщик", "Стать Go-разработчиком"}
	currentDay = daysSince(startDateStr)
	progress   = calculateProgress()
	theme      = "Conditionals: if"
	r          = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()>>32)))
)

func main() {
	printHeader()
	printProgress()
	printToday()
	printStats()
	printAchievements()
	printFuture()
	printFooter()
	interactiveLineCounter()
}

func daysSince(dateStr string) int {
	t, _ := time.Parse("2006-01-02", dateStr)
	return int(time.Since(t).Hours() / 24)
}

func calculateProgress() Progress {
	xp := 100 + currentDay*10
	return Progress{
		Days:       currentDay,
		XP:         xp,
		Level:      1 + xp/maxLevelXP,
		Streak:     currentDay,
		Confidence: min(100, currentDay*2),
		CodeLines:  float64(currentDay) * 42.5,
	}
}

func printHeader() {
	fmt.Printf("\n%s🔥 100 ДНЕЙ GО С ГОШЕЙ: ОТ БЕЗУМНЫХ ТУСОВОК К ПРОДУМАННОМУ КОДУ 🔥%s\n",
		"\033[1;33m", "\033[0m")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("👤 %s%s%s | %d лет | %s\n",
		"\033[1;36m", gosha.Name, "\033[0m", gosha.Age, gosha.Background)
	fmt.Printf("🎯 %s%s%s\n",
		"\033[1;32m", gosha.Goal, "\033[0m")
	fmt.Printf("📅 %s | День: %d/%d | Тема: %s\n",
		time.Now().Format("02.01.2006"), currentDay, challengeDays, theme)
}

func printProgress() {
	percent := float64(currentDay) / challengeDays * 100
	xpNeeded := progress.Level * maxLevelXP

	fmt.Printf("\n%s🚀 ПРОГРЕСС: %.0f%% завершено (Дней с 30 ноября: %d)%s\n",
		"\033[1;34m", percent, currentDay, "\033[0m")
	fmt.Println(progressBar(percent, 40))
	fmt.Printf("🏆 Уровень %d (%d/%d XP) | 💪 Уверенность: %d%%\n",
		progress.Level, progress.XP, xpNeeded, progress.Confidence)
}

func printToday() {
	motivations := []string{
		"38 лет — идеальный возраст для старта в IT. Твой опыт жизни — твоя суперсила!",
		"Каждая строка кода сегодня — это шаг к свободе от курьерской работы завтра.",
		"Ты не 'поздно начинаешь'. Ты начинаешь в идеальное время с багажом жизненного опыта.",
		"Твоя миссия — не просто выучить Go. Твоя миссия — доказать, что никогда не поздно менять жизнь.",
		"Помни: самые успешные разработчики начинали с нуля. Разница в том, что ты начал с опыта жизни.",
		"Твой рэп научил тебя ритму и структуре. Теперь примени это к коду!",
		"Каждый раз, когда хочется сдаться, вспомни: через год ты будешь жалеть, что НЕ продолжил сегодня.",
	}

	facts := []string{
		"В Go нет наследования классов, только композиция. Как и в жизни — составляй свою судьбу из лучших частей!",
		"Go создан в Google для решения реальных проблем. И твоя проблема — реальна и достойна решения.",
		"Средняя зарплата Go-разработчика в Москве — 220,000 ₽. Это твоя цель через 6 месяцев.",
		"Telegram, Docker, Kubernetes — все они частично написаны на Go. Твой код тоже может изменить мир.",
		"1 горутина в Go = 1 поток выполнения. 1 день твоего челленджа = 1 шаг к новой жизни.",
		"В Go есть поговорка: 'Меньше кода — меньше багов'. В жизни: 'Меньше тусовок — больше смысла'.",
	}

	fmt.Printf("\n%s✨ СЕГОДНЯ ГОВОРИТ СЕРДЦЕ:%s\n", "\033[1;35m", "\033[0m")
	fmt.Printf("   💬 %s\n", motivations[currentDay%len(motivations)])
	fmt.Printf("   💡 %s\n", facts[currentDay%len(facts)])
}

func printStats() {
	cigarettes := currentDay * 15
	beerBottles := currentDay * 3
	studyHours := float64(currentDay) * 2.5
	moneySaved := float64(cigarettes)*15 + float64(beerBottles)*120 + float64(currentDay)*35

	fmt.Printf("\n%s📊 СТАТИСТИКА ПЕРЕРОЖДЕНИЯ:%s\n", "\033[1;36m", "\033[0m")
	fmt.Printf("   🚭 Пропущено сигарет: %d | 🍺 Бутылок пива: %d\n", cigarettes, beerBottles)
	fmt.Printf("   💻 Часов обучения: %.1f | 💰 Сэкономлено: %.0f ₽\n", studyHours, moneySaved)
	fmt.Printf("   📝 Написано строк кода: %.0f | 🔥 Удалено игр: %d\n", progress.CodeLines, 7+currentDay/5)
}

func printAchievements() {
	achievements := []Achievement{
		{"🌱", "Новое начато!", "Первые 24 часа без пагубных привычек", 1, "common"},
		{"⚔️", "Цифровой Гуру", "7 дней чистого кода вместо сериалов", 7, "uncommon"},
		{"💎", "Сердце чемпиона", "Ты прошел четверть пути! 25 дней перемен!", 25, "rare"},
		{"🚀", "Наполовину к звёздам", "50 дней без оглядки назад — только вперёд!", 50, "epic"},
		{"🌟", "Полный круг", "100 дней нового Гоши — легенда в мире кода!", 100, "legendary"},
	}

	unlocked := 0
	for _, a := range achievements {
		if currentDay >= a.Day {
			unlocked++
		}
	}

	fmt.Printf("\n%s🏆 ДОСТИЖЕНИЯ (%d/%d):%s\n", "\033[1;33m", unlocked, len(achievements), "\033[0m")
	for _, a := range achievements {
		if currentDay >= a.Day {
			fmt.Printf("   %s%s %s%s\n", "\033[1;32m", a.Emoji, a.Name, "\033[0m")
		}
	}
}

func printFuture() {
	currentSalary := 80000 + currentDay*1700
	daysToJob := max(0, 45-currentDay)

	fmt.Printf("\n%s🔮 БУДУЩЕЕ ЧЕРЕЗ 100 ДНЕЙ:%s\n", "\033[1;35m", "\033[0m")
	fmt.Printf("   💼 Go-разработчик в Биг-Техе (Текущая: %d ₽/мес → %s250,000 ₽/мес%s)\n",
		currentSalary, "\033[1;32m", "\033[0m")
	fmt.Printf("   🏠 Своя квартира-студия в новом районе у метро (мечта с 30 ноября)\n")
	fmt.Printf("   👵 Родители гордятся тобой (а не тем, как ты 'прославился' в прошлом)\n")
	fmt.Printf("   ⏳ Работа найдется через %d дней. Ты справишься!\n", daysToJob)
}

func printFooter() {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("%s💬 ФИЛОСОФИЯ 38-ЛЕТНЕГО ГОШИ:%s\n", "\033[1;34m", "\033[0m")
	fmt.Println("   \"Я не 'поздно начинаю'. Я начинаю в то время, когда другие сдаются.\"")
	fmt.Println("   \"Моя жизнь до 38 лет — мой 'legacy код'. Теперь я рефакторю свою судьбу.\"")
	fmt.Println("   \"Каждая буква 'G' в 'Go' означает: Growth, Goals, Glory.\"")

	birthdayMessage := "\n%s🎉 НАПОМИНАНИЕ: 30 ноября 2025 года тебе исполнилось 38 лет. " +
		"Это не конец молодости — это начало твоей самой важной главы!%s"
	fmt.Printf(birthdayMessage, "\033[1;33m", "\033[0m")

	fmt.Printf("\n%s🌟 СЕГОДНЯ: УДАЛИЛ 1 ИГРУ + НАПИСАЛ 42 СТРОКИ КОДА! ТВОЙ ПУТЬ ПРОДОЛЖАЕТСЯ! 🌟%s\n",
		"\033[1;32m", "\033[0m")
}

func progressBar(percent float64, width int) string {
	filled := int(percent / 100 * float64(width))
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func interactiveLineCounter() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Print("🔍 Проверить прогресс за другой день (например: day25): ")

	var day string
	fmt.Scanln(&day)

	if !strings.HasPrefix(day, "day") {
		fmt.Printf("%s⚠️ Неправильный формат! Используй: day25%s\n", "\033[1;31m", "\033[0m")
		return
	}

	dir := fmt.Sprintf("../%s", day)
	lines, err := countCodeLines(dir)

	if err != nil {
		fmt.Printf("%s❌ Ошибка: %v%s\n", "\033[1;31m", err, "\033[0m")
		return
	}

	emoji := "✅"
	if lines > 100 {
		emoji = "🔥"
	} else if lines < 10 {
		emoji = "💪"
	}

	fmt.Printf("\n%s%s %s: %.0f строк кода!%s\n",
		"\033[1;32m", emoji, day, lines, "\033[0m")

	if lines > 0 {
		fmt.Printf("%s💡 Совет: Добавь комментарии и рефакторинг!%s\n", "\033[1;34m", "\033[0m")
	}
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
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
			if line != "" && !strings.HasPrefix(line, "//") {
				total++
			}
		}
		return nil
	})
	return total, err
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
