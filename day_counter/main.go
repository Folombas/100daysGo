package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const (
	hundredDaysStart = "2025-11-03"
	go365Start       = "2026-01-01"
	maxLevelXP       = 1000
	focusBonusXP     = 25
)

type Progress struct {
	DaysCount, TotalXP, Level int
	FocusDepth                int
	MonthsFocused             int
}

type App struct {
	currentDate time.Time
	progress    map[string]Progress
	rng         *rand.Rand
	DailyTopic  string // Тема дня (менять здесь каждый день)
}

func NewApp() *App {
	now := time.Now()
	hundredDays := max(1, daysSince(hundredDaysStart))
	go365Days := max(1, daysSince(go365Start))

	return &App{
		currentDate: now,
		rng:         rand.New(rand.NewPCG(uint64(now.UnixNano()), uint64(now.Unix()))),
		DailyTopic:  "Interfaces: Type Assertions", // ← МЕНЯТЬ ЭТУ СТРОКУ КАЖДЫЙ ДЕНЬ
		progress: map[string]Progress{
			"100daysGo": {
				DaysCount:  hundredDays,
				TotalXP:    hundredDays * (15 + focusBonusXP),
				Level:      1 + hundredDays*(15+focusBonusXP)/maxLevelXP,
				FocusDepth: hundredDays / 10,
			},
			"Go365": {
				DaysCount:     go365Days,
				TotalXP:       go365Days * (25 + focusBonusXP),
				Level:         1 + go365Days*(25+focusBonusXP)/maxLevelXP,
				FocusDepth:    go365Days / 5,
				MonthsFocused: max(1, go365Days/30),
			},
		},
	}
}

func main() {
	app := NewApp()
	app.renderFocusUI()
}

func (a *App) renderFocusUI() {
	a.printHeader()
	a.printProgress()
	a.printFocusManifesto()
	a.printDailyTopic() // Новый раздел
	a.printFooter()
}

func (a *App) printHeader() {
	go365 := a.progress["Go365"]
	fmt.Printf("\n%s🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO%s\n", ansi("1;31"), ansi("0"))
	fmt.Println("▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰")
	fmt.Printf("📅 %s | 🔥 День %d абсолютного фокуса\n",
		a.currentDate.Format("02.01.2006"), go365.DaysCount)
	fmt.Printf("🧠 АБСОЛЮТНЫЙ ОТКАЗ ОТ РАСПЫЛЕНИЯ | Уровень глубины: %d\n", go365.FocusDepth)
	fmt.Printf("⚡ Месяцев без тусовок/баров/игр: %d\n", go365.MonthsFocused)
}

func (a *App) printProgress() {
	fmt.Printf("\n%sПРОГРЕСС ОТКАЗА ОТ РАСПЫЛЕНИЯ%s\n", ansi("1;34"), ansi("0"))

	for name, p := range a.progress {
		percent := p.DaysCount * 100 / map[string]int{"100daysGo": 100, "Go365": 365}[name]
		fmt.Printf("\n▸ %s: День %d | Ур.%d | Фокус-XP: %d\n", name, p.DaysCount, p.Level, p.TotalXP)
		printFocusBar(percent)
	}

	fmt.Printf("\n⛏️  Уровень концентрации: %s\n", a.getFocusLevel())
}

func (a *App) getFocusLevel() string {
	depth := a.progress["Go365"].FocusDepth
	switch {
	case depth >= 20:
		return "🚫 АБСОЛЮТНЫЙ АСКЕТИЗМ (только Go)"
	case depth >= 15:
		return "🔥 ПОЛНЫЙ ОТКАЗ ОТ РАЗВЛЕЧЕНИЙ"
	case depth >= 10:
		return "⚡ ЖЁСТКАЯ ФОКУСИРОВКА"
	case depth >= 5:
		return "🎯 ОТКАЗ ОТ ТУСОВОК"
	default:
		return "🌱 НАЧАЛО АСКЕЗЫ"
	}
}

func (a *App) printFocusManifesto() {
	fmt.Printf("\n%s🚫 МАНИФЕСТ ОТКАЗА ОТ РАСПЫЛЕНИЯ%s\n", ansi("1;31"), ansi("0"))
	fmt.Println("┌─────────────────────────────────────────────────────┐")

	manifesto := []string{
		"│ 1. БАРЫ/РЕСТОРАНЫ      → ❌ ЗАПРЕЩЕНО               │",
		"│ 2. КЛУБЫ/ТУСОВКИ       → ❌ ЗАПРЕЩЕНО               │",
		"│ 3. ФИЛЬМЫ/СЕРИАЛЫ      → ❌ ЗАПРЕЩЕНО               │",
		"│ 4. ВИДЕОИГРЫ           → ❌ ЗАПРЕЩЕНО               │",
		"│ 5. СОЦСЕТИ/SCROLL      → ❌ ЗАПРЕЩЕНО               │",
		"│ 6. ПУСТЫЕ РАЗГОВОРЫ    → ❌ ЗАПРЕЩЕНО               │",
		"│ 7. СПОРТИВНЫЕ СТАВКИ   → ❌ ЗАПРЕЩЕНО               │",
		"│ 8. ШОППИНГ РАДИ КАЙФА  → ❌ ЗАПРЕЩЕНО               │",
		"│                                                  │",
		"│ ✅ РАЗРЕШЕНО ТОЛЬКО: GO + КОД + ДОКУМЕНТАЦИЯ      │",
	}

	for _, line := range manifesto {
		fmt.Println(line)
	}
	fmt.Println("└─────────────────────────────────────────────────────┘")
}

// НОВЫЙ РАЗДЕЛ: Тема дня
func (a *App) printDailyTopic() {
	fmt.Printf("\n%s📚 ТЕМА ДНЯ: %s%s\n", ansi("1;36"), ansi("0"), a.DailyTopic)

	// Конкретные задачи по теме
	tasks := map[string][]string{
		"Embedding Interfaces": {
			"1. Изучить композицию структур через встраивание",
			"2. Практика: создать 3 примера embedding",
			"3. Понять разницу между embedding и inheritance",
			"4. Прочитать главу 'Embedding' в 'Go in Action'",
			"5. Написать блог-пост с примерами",
		},
		"Concurrency Patterns": {
			"1. Worker pool на горутинах",
			"2. Fan-in/fan-out шаблоны",
			"3. Context для отмены операций",
			"4. Select с таймаутами",
			"5. Реализация rate limiter",
		},
		"Standard Library": {
			"1. Изучить 5 пакетов из stdlib",
			"2. Прочитать исходники fmt или net/http",
			"3. Написать wrapper над существующим пакетом",
			"4. Разобрать внутреннее устройство",
			"5. Создать cheatsheet по пакету",
		},
	}

	// Показываем задачи для текущей темы или общие
	if topicTasks, exists := tasks[a.DailyTopic]; exists {
		for _, task := range topicTasks {
			fmt.Printf("   %s\n", task)
		}
	} else {
		// Общие задачи, если темы нет в списке
		generalTasks := []string{
			"1. 100+ строк кода по теме дня",
			"2. Прочитать документацию",
			"3. Написать примеры",
			"4. Сделать mind map концепции",
			"5. Поделиться инсайтами в блоге",
		}
		for _, task := range generalTasks {
			fmt.Printf("   %s\n", task)
		}
	}

	// Уровень понимания темы (на основе дня)
	understandingLevel := a.progress["Go365"].FocusDepth % 10
	if understandingLevel == 0 {
		understandingLevel = 1
	}
	fmt.Printf("\n   🎯 Уровень понимания темы: %d/10\n", understandingLevel)
}

func (a *App) printFooter() {
	fmt.Println("\n▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰")
	fmt.Printf("%s💬 КРЕДО АСКЕТА-ПРОГРАММИСТА:%s\n", ansi("1;36"), ansi("0"))

	quotes := []string{
		"«Каждая несыгранная игра — это 100 строк кода. Каждый непосещённый бар — это новый скилл.»",
		"«Распыление создаёт дилетантов. Аскеза создаёт мастеров. Я выбираю аскезу.»",
		"«Мои тусовки — это коммиты. Мои вечеринки — это дебаг сессии. Мои друзья — это горутины.»",
		"«Отказ от 1 часа фильма = +1 час к мастерству в Go. Математика простая.»",
		"«Бар забирает деньги и время. Go даёт деньги и свободу. Выбор очевиден.»",
	}

	fmt.Printf("   %s\n\n", quotes[a.rng.IntN(len(quotes))])
	fmt.Printf("%s🚀 GO ИЛИ НИЧЕГО. КОД ИЛИ НИЧЕГО.%s\n", ansi("1;35"), ansi("0"))
}

func daysSince(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return int(time.Since(t).Hours() / 24)
}

func printFocusBar(percent int) {
	width := 40
	filled := percent * width / 100
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	fmt.Printf("   [%s] %d%%\n", bar, percent)
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
