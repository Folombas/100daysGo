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
	maxLevelXP       = 1000
	codeLinesPerDay  = 67.3
	focusBonusXP     = 25 // Бонус за фокус на Go
)

type Progress struct {
	DaysCount, TotalXP, Level int
	CodeLines                 float64
	FocusDepth                int
	MonthsFocused             int
}

type App struct {
	currentDate time.Time
	progress    map[string]Progress
	theme       string
	rng         *rand.Rand
	dailyFocus  string
}

func NewApp() *App {
	now := time.Now()

	// Глубокое погружение в расчёт дней с учётом 4 января
	hundredDays := max(1, daysSince(hundredDaysStart))
	go365Days := max(1, daysSince(go365Start))

	return &App{
		currentDate: now,
		theme:       "ГЛУБИНА GO: Empty Interface | День погружения: %d",
		rng:         rand.New(rand.NewPCG(uint64(now.UnixNano()), uint64(now.Unix()))),
		dailyFocus:  getDailyFocus(go365Days),
		progress: map[string]Progress{
			"100daysGo": {
				DaysCount:   hundredDays,
				TotalXP:     hundredDays * (15 + focusBonusXP),
				Level:       1 + hundredDays*(15+focusBonusXP)/maxLevelXP,
				CodeLines:   float64(hundredDays) * codeLinesPerDay,
				FocusDepth:  hundredDays / 14,
			},
			"Go365": {
				DaysCount:      go365Days,
				TotalXP:        go365Days * (25 + focusBonusXP),
				Level:          1 + go365Days*(25+focusBonusXP)/maxLevelXP,
				CodeLines:      float64(go365Days) * codeLinesPerDay,
				FocusDepth:     go365Days / 7,
				MonthsFocused:  max(1, go365Days/30),
			},
		},
	}
}

func main() {
	app := NewApp()
	app.renderDeepFocusUI()
}

func (a *App) renderDeepFocusUI() {
	a.printHeader()
	a.printProgress()
	a.printDepthAnalysis()
	a.printDailyChallenge()
	a.printFooter()
}

func (a *App) printHeader() {
	go365 := a.progress["Go365"]
	fmt.Printf("\n%sГОФЕР-ГРЫЗУН ПОГРУЖАЕТСЯ В GO%s\n", ansi("1;36"), ansi("0"))
	fmt.Println("▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰")
	fmt.Printf("📅 %s | 🎯 День %d погружения в Go365\n",
		a.currentDate.Format("02.01.2006"), go365.DaysCount)
	fmt.Printf("🧠 %s\n", fmt.Sprintf(a.theme, go365.FocusDepth))
	fmt.Printf("⚡ Уровень глубины: %d | Месяцев фокуса: %d\n",
		go365.FocusDepth, go365.MonthsFocused)
}

func (a *App) printProgress() {
	fmt.Printf("\n%sПРОГРЕСС ПОГРУЖЕНИЯ%s\n", ansi("1;34"), ansi("0"))

	for name, p := range a.progress {
		percent := p.DaysCount * 100 / map[string]int{"100daysGo": 100, "Go365": 365}[name]
		fmt.Printf("\n▸ %s: День %d | Ур.%d | XP: %d\n", name, p.DaysCount, p.Level, p.TotalXP)
		printDepthBar(percent)
	}

	fmt.Printf("\n📊 Написано строк с фокусом: %.0f (%.1f/день)\n",
		a.progress["100daysGo"].CodeLines + a.progress["Go365"].CodeLines,
		codeLinesPerDay)
	fmt.Printf("⛏️  Уровень концентрации: %s\n", a.getFocusLevel())
}

func (a *App) getFocusLevel() string {
	depth := a.progress["Go365"].FocusDepth
	switch {
	case depth >= 20: return "🔱 ЭКСТРЕМАЛЬНАЯ ГЛУБИНА"
	case depth >= 15: return "🏊 ГЛУБОКОЕ ПОГРУЖЕНИЕ"
	case depth >= 10: return "⚡ СИЛЬНАЯ ФОКУСИРОВКА"
	case depth >= 5:  return "🎯 УМЕРЕННЫЙ ФОКУС"
	default:         return "🌱 НАЧАЛО ПОГРУЖЕНИЯ"
	}
}

func (a *App) printDepthAnalysis() {
	fmt.Printf("\n%sАНАЛИЗ ГЛУБИНЫ: Empty Interface%s\n", ansi("1;35"), ansi("0"))
	fmt.Println("┌─────────────────────────────────────────────────────┐")

	levels := []struct{
		level int
		desc  string
	}{
		{1, "interface{} как любой тип (поверхностно)"},
		{2, "type assertion и type switch"},
		{3, "Отражение (reflect) с empty interface"},
		{4, "Производительность и аллокации"},
		{5, "Использование в stdlib (json, fmt)"},
		{6, "Альтернативы: generics, конкретные типы"},
		{7, "Компиляторные оптимизации"},
		{8, "Внутренняя реализация в рантайме"},
		{9, "Создание type-safe обёрток"},
		{10,"Мастерское владение (уровень контрибьютера)"},
	}

	currentDepth := a.progress["Go365"].FocusDepth
	for _, l := range levels {
		status := "🔒"
		if currentDepth >= l.level {
			status = "✅"
		}
		fmt.Printf("│ %s Ур.%2d: %-45s │\n", status, l.level, l.desc)
	}
	fmt.Println("└─────────────────────────────────────────────────────┘")
}

func (a *App) printDailyChallenge() {
	fmt.Printf("\n%sСЕГОДНЯШНЕЕ ПОГРУЖЕНИЕ (4 января)%s\n", ansi("1;33"), ansi("0"))
	fmt.Printf("💡 Фокус дня: %s\n", a.dailyFocus)

	challenges := []string{
		"1. Разобрать 3 использования interface{} в stdlib",
		"2. Написать type-safe обёртку над interface{}",
		"3. Измерить производительность type assertion",
		"4. Прочитать исходники пакета reflect",
		"5. Написать блог-пост о прозрениях",
	}

	for _, ch := range challenges {
		fmt.Printf("   %s\n", ch)
	}

	fmt.Printf("\n🏆 Достижения глубины: ")
	unlocked := 0
	achievements := []string{"🔱", "🧠", "⚡", "⛏️", "🧬"}
	for i, ach := range achievements {
		if a.progress["Go365"].FocusDepth > i*2 {
			fmt.Printf("%s", ach)
			unlocked++
		} else {
			fmt.Printf("🔒")
		}
	}
	fmt.Printf(" (%d/5)\n", unlocked)
}

func (a *App) printFooter() {
	fmt.Println("\n▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰")
	fmt.Printf("%s💬 ГОФЕР-ГРЫЗУН ГОВОРИТ:%s\n", ansi("1;36"), ansi("0"))

	quotes := []string{
		"«Empty interface — это не дыра в системе типов, а туннель к гибкости»",
		"«Глубина понимания interface{} = глубина понимания всей системы типов Go»",
		"«Не используй interface{} там, где можно использовать дженерики»",
		"«Каждый type assertion — это шаг вглубь системы типов»",
		"«reflect — это микроскоп для изучения interface{}»",
	}

	fmt.Printf("   %s\n\n", quotes[a.rng.IntN(len(quotes))])
	fmt.Printf("%s🚀 ПОГРУЖАЙСЯ ГЛУБЖЕ! ВГРЫЗАЙСЯ В GO С УСЕРДИЕМ СУСЛИКА-ГОФЕРА!%s\n",
		ansi("1;35"), ansi("0"))
}

// --- УТИЛИТЫ ---

func daysSince(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return int(time.Since(t).Hours()/24)
}

func printDepthBar(percent int) {
	width := 40
	filled := percent * width / 100

	bar := ""
	for i := 0; i < width; i++ {
		switch {
		case i < filled/3:
			bar += "█"
		case i < filled*2/3:
			bar += "▓"
		case i < filled:
			bar += "░"
		default:
			bar += " "
		}
	}
	fmt.Printf("   [%s] %d%%\n", bar, percent)
}

func getDailyFocus(day int) string {
	foci := []string{
		"ГЛУБИНА: Интерфейсы — Empty Interface",
		"ФОКУС: Система типов Go",
		"ПОГРУЖЕНИЕ: Компилятор Go",
		"ИЗУЧЕНИЕ: Стандартная библиотека",
		"АНАЛИЗ: Производительность",
		"ПРАКТИКА: Паттерны проектирования",
		"РЕФЛЕКСИЯ: Анализ кода",
	}
	return foci[day%len(foci)]
}

func ansi(code string) string {
	return "\033[" + code + "m"
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

// Простая версия для подсчёта строк (без изменений)
func countCodeLines(dir string) (float64, error) {
	var total float64
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		file, _ := os.Open(path)
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
	return total, nil
}
