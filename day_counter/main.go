package main

import (
	"fmt"
	"time"
)

// Конфигурация
const (
	DATE_FORMAT = "02.01.2006"
	BAR_WIDTH   = 30
)

// Структуры данных
type Challenge struct {
	Name      string
	StartDate string
	TotalDays int
}

type DailyTopic struct {
	Title    string
	Tasks    []string
	MinLines int
}

// Основная функция
func main() {
	// Инициализация данных
	challenges := map[string]Challenge{
		"100daysGo": {"100daysGo", "2025-11-03", 100},
		"Go365":     {"Go365", "2026-01-01", 365},
	}

	today := time.Now()
	go365Start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	go365Day := calcDaysBetween(go365Start, today)

	// Тема дня
	todayTopic := DailyTopic{
		Title: "Generics in Go Programming Language: Type Constraints",
		Tasks: []string{
			"100+ строк кода по теме",
			"Прочитать документацию по type constraints",
			"Написать примеры с comparable и any",
			"Создать конспект в Obsidian",
			"Поделиться инсайтами в TG-канале",
		},
		MinLines: 100,
	}

	// Вывод
	printHeader(today, go365Day)
	printChallengesProgress(challenges, today, go365Day)
	printTopicBox(&todayTopic, go365Day)
	printFocusManifesto()
	printAllowedActivities()
	printFooter(go365Day)
}

// ========== УТИЛИТЫ ==========

func calcDaysBetween(start, end time.Time) int {
	days := int(end.Sub(start).Hours()/24) + 1
	if days < 1 {
		return 1
	}
	return days
}

func getFocusLevel(day int) string {
	switch {
	case day >= 30:
		return "🚀 КОСМИЧЕСКАЯ КОНЦЕНТРАЦИЯ"
	case day >= 20:
		return "🔥 ПОЛНЫЙ ФОКУС"
	case day >= 10:
		return "⚡ ВЫСОКАЯ СКОРОСТЬ"
	default:
		return "🌱 НАЧАЛО ПУТИ"
	}
}

// ========== ВЫВОД ==========

func printHeader(date time.Time, day int) {
	fmt.Printf("\n🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO\n")
	fmt.Println("════════════════════════════════════")
	fmt.Printf("📅 %s | 🔥 День %d абсолютного фокуса\n", date.Format(DATE_FORMAT), day)
	fmt.Printf("🧠 Уровень концентрации: %s\n\n", getFocusLevel(day))
}

func printChallengesProgress(challenges map[string]Challenge, today time.Time, go365Day int) {
	fmt.Println("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ")

	for name, ch := range challenges {
		days := calcChallengeDays(name, ch, today, go365Day)
		percent := days * 100 / ch.TotalDays
		if percent > 100 {
			percent = 100
		}
		level := min(days/10+1, 10)

		fmt.Printf("\n▸ %s: День %d | Ур.%d\n", name, days, level)
		printProgressBar(percent)
	}
}

func calcChallengeDays(name string, ch Challenge, today time.Time, go365Day int) int {
	if name == "Go365" {
		return go365Day
	}
	start, _ := time.Parse("2006-01-02", ch.StartDate)
	days := int(today.Sub(start).Hours()/24) + 1
	if days < 0 {
		return 0
	}
	return days
}

func printProgressBar(percent int) {
	fmt.Print("   [")
	filled := percent * BAR_WIDTH / 100
	for i := 0; i < BAR_WIDTH; i++ {
		if i < filled {
			fmt.Print("█")
		} else {
			fmt.Print("░")
		}
	}
	fmt.Printf("] %d%%\n", percent)
}

func printTopicBox(topic *DailyTopic, day int) {
	fmt.Println("\n📚 ТЕМА ДНЯ")
	fmt.Println("   ┌─────────────────────────────────────────┐")
	fmt.Printf("   │ %-39s │\n", topic.Title)
	fmt.Println("   │                                         │")

	understanding := (day % 10) + 1
	emoji := "🟢"
	switch {
	case understanding <= 3:
		emoji = "🔴"
	case understanding <= 7:
		emoji = "🟡"
	}

	fmt.Printf("   │ %s Уровень понимания: %d/10           │\n", emoji, understanding)
	fmt.Printf("   │   Цель: %d+ строк кода                │\n", topic.MinLines)
	fmt.Println("   │                                         │")

	for i, task := range topic.Tasks {
		if i < 3 { // Показываем только первые 3 задачи
			fmt.Printf("   │   • %-31s │\n", task)
		}
	}

	if len(topic.Tasks) > 3 {
		fmt.Printf("   │   • ...и ещё %d задач               │\n", len(topic.Tasks)-3)
	}

	fmt.Println("   └─────────────────────────────────────────┘")
}

func printFocusManifesto() {
	fmt.Println("\n📜 МАНИФЕСТ ФОКУСА")
	items := []string{
		"БАРЫ/КЛУБЫ                  → ❌ НЕТ",
		"ФИЛЬМЫ/СЕРИАЛЫ              → ❌ НЕТ",
		"ВИДЕОИГРЫ                   → ❌ НЕТ",
		"SCROLL                      → ❌ НЕТ",
		"ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ → ❌ НЕТ",
	}
	printBox(items)
	fmt.Println("   ✅ РАЗРЕШЕНО: GO + КОД + ДОКУМЕНТАЦИЯ")
}

func printAllowedActivities() {
	fmt.Println("\n🎯 ФОКУС НА РАЗВИТИИ")
	items := []string{
		"ПИСАТЬ КОД                     → ✅ ДА",
		"ОСНОВЫ LINUX                   → ✅ ДА",
		"TERMINAL                       → ✅ ДА",
		"DOCKER                         → ✅ ДА",
		"АЛГОРИТМЫ И СТРУКТУРЫ ДАННЫХ   → ✅ ДА",
		"ЧИТАТЬ ДОКУМЕНТАЦИЮ            → ✅ ДА",
		"СОЗДАВАТЬ ПРОЕКТЫ              → ✅ ДА",
		"РЕШАТЬ ЗАДАЧИ НА LEETCODE      → ✅ ДА",
		"ИЗУЧАТЬ АРХИТЕКТУРУ ПО         → ✅ ДА",
		"ПИСАТЬ ТЕСТЫ                   → ✅ ДА",
		"РАБОТАТЬ С GIT                 → ✅ ДА",
		"ЧИТАТЬ ЧУЖОЙ КОД               → ✅ ДА",
	}
	printBox(items)
	fmt.Println("   🔥 НАПИСАТЬ КОД ЛУЧШЕ, ЧЕМ НАПИСАТЬ ОПРАВДАНИЯ ТРОЛЛЯМ")
}

func printBox(items []string) {
	fmt.Println("   ┌─────────────────────────────────────────┐")
	for _, item := range items {
		fmt.Printf("   │ %-40s │\n", item)
	}
	fmt.Println("   └─────────────────────────────────────────┘")
}

func printFooter(day int) {
	fmt.Println("\n════════════════════════════════════")

	quotes := []string{
		"«Каждый день кода — шаг к свободе»",
		"«Распыление создаёт дилетантов. Фокус — мастеров»",
		"«Мои тусовки — это коммиты. Мои друзья — это горутины»",
		"«1 час кода стоит 10 часов скроллинга»",
		"«Тролли кормятся вниманием. Лишай их питания — пиши код»",
		"«Лучший ответ троллю — твой следующий коммит»",
		"«Код не врёт в отличие людей»",
		"«Компилятор — самый честный критик»",
	}

	fmt.Printf("💬 %s\n", quotes[day%len(quotes)])
	fmt.Println("\n🚀 GO ИЛИ НИЧЕГО. КОД ИЛИ НИЧЕГО.")
	fmt.Println("   🔥 НЕ ОТВЕЧАЙ НА ТРОЛЛЕЙ — ОТВЕЧАЙ КОДОМ")
}

// Вспомогательная функция для Go 1.20 и ниже
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
