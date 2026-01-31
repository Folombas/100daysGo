package main

import (
	"fmt"
	"strings"
	"time"
)

// Конфигурация
const (
	DATE_FORMAT   = "02.01.2006"
	BAR_WIDTH     = 25
	DISPLAY_LIMIT = 3
)

// Структуры данных
type Challenge struct {
	Name      string
	StartDate string
	TotalDays int
}

type DailyTopic struct {
	Title      string
	Tasks      []string
	MinLines   int
	Category   string
	Complexity string
	Priority   int
}

type EcoSystemFocus struct {
	Category string
	Skills   []string
	Progress int // 0-100%
}

// Основные данные
var (
	challenges = map[string]Challenge{
		"Go365":     {"Go365", "2026-01-01", 365},
		"100daysGo": {"100daysGo", "2025-11-03", 100},
	}

	todayTopic = DailyTopic{
		Title:      "Channels: Buffered vs Unbuffered",
		MinLines:   100,
		Category:   "Go Core",
		Complexity: "Intermediate",
		Priority:   1,
		Tasks: []string{
			"Изучить документацию по Type Parameters",
			"Разобрать примеры с comparable и any",
			"Написать generic функции для работы с коллекциями",
			"Понять ограничения type constraints",
			"Создать свой generic тип с методами",
			"Прочитать статью о performance implications",
			"Решить 3 задачи на LeetCode с использованием generics",
		},
	}

	ecosystemFocus = []EcoSystemFocus{
		{
			Category: "Go Core",
			Skills:   []string{"Goroutines", "Channels", "Interfaces", "Generics", "Reflection"},
			Progress: 65,
		},
		{
			Category: "Linux & Terminal",
			Skills:   []string{"Bash Scripting", "Systemd", "Networking", "Permissions", "Process Management"},
			Progress: 40,
		},
		{
			Category: "DevOps & Containers",
			Skills:   []string{"Docker", "Docker Compose", "CI/CD", "Kubernetes Basics", "Monitoring"},
			Progress: 30,
		},
		{
			Category: "Databases",
			Skills:   []string{"PostgreSQL", "Redis", "MongoDB", "SQL Optimization", "Migrations"},
			Progress: 25,
		},
		{
			Category: "Backend Development",
			Skills:   []string{"REST APIs", "gRPC", "Authentication", "Testing", "Logging"},
			Progress: 50,
		},
	}

	negativeItems = []string{
		"БАРЫ / КЛУБЫ", "ФИЛЬМЫ / СЕРИАЛЫ", "ВИДЕОИГРЫ",
		"БЕССМЫСЛЕННЫЙ SCROLL", "ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ",
		"СОЦИАЛЬНЫЕ СЕТИ", "НОВОСТНЫЕ ЛЕНТЫ",
	}

	rules = []string{
		"Код > Оправданий", "Коммит > Скроллинга", "Документация > Догадок",
		"Тесты > Уверенности", "Простота > Умности", "Практика > Теории",
		"Git > Памяти", "Docker > 'У меня работает'",
	}

	quotes = []string{
		"«Системный подход к обучению создаёт системного разработчика»",
		"«Каждый коммит — шаг к мастерству в Go и его экосистеме»",
		"«Инвестиции в знания Go, Linux и DevOps окупаются экспоненциально»",
		"«1 час целенаправленного кода на Go стоит 10 часов поверхностного изучения»",
		"«Экосистема — это не только язык, но и окружение, инструменты, практики»",
		"«Контейнеризация знания: Docker для кода, Kubernetes для карьеры»",
		"«Базы данных — память приложения, алгоритмы — его интеллект»",
		"«Полное погружение в Go: от горутин до продакшн-деплоя»",
	}
)

// Основная функция
func main() {
	today := time.Now()
	go365Start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	go365Day := daysBetween(go365Start, today)

	printHeader(today, go365Day)
	printSection("📚 ТЕМА ДНЯ", func() { printTopic(go365Day) })
	printSection("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ", func() { printChallenges(today, go365Day) })
	printSection("🎯 ФОКУС НА ЭКОСИСТЕМЕ GO", func() { printFocus() })
	printSection("📜 МАНИФЕСТ ПОЛНОГО ПОГРУЖЕНИЯ", printManifesto)
	printSection("⚡ ПРАВИЛА GO-РАЗРАБОТЧИКА", printRules)
	printFooter(go365Day)
}

// ========== УТИЛИТЫ ==========

func daysBetween(start, end time.Time) int {
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

func getDaySymbol(day int) string {
	if day <= 30 {
		return "①"
	}
	return fmt.Sprintf("%d", day)
}

func getUnderstandingEmoji(level int) string {
	switch {
	case level >= 8:
		return "🎯"
	case level >= 5:
		return "⚡"
	default:
		return "📚"
	}
}

func progressBar(percent int) string {
	filled := percent * BAR_WIDTH / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("░", BAR_WIDTH-filled)
	return fmt.Sprintf("[%s] %d%%", bar, percent)
}

func printList(items []string, limit int, numbered bool) {
	for i, item := range items {
		if i >= limit {
			fmt.Printf("   ...и ещё %d задач\n", len(items)-limit)
			break
		}
		if numbered {
			fmt.Printf("  %d. %s\n", i+1, item)
		} else {
			fmt.Printf("  • %s\n", item)
		}
	}
}

func printSection(title string, contentFunc func()) {
	fmt.Printf("\n%s\n", title)
	fmt.Println(strings.Repeat("─", 50))
	contentFunc()
}

// ========== КОМПОНЕНТЫ ВЫВОДА ==========

func printHeader(date time.Time, day int) {
	fmt.Println("\n🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO")
	fmt.Println("═" + strings.Repeat("═", 48))
	fmt.Printf("📅 %s | %s День %d абсолютного фокуса\n",
		date.Format(DATE_FORMAT), getDaySymbol(day), day)
	fmt.Printf("🧠 Уровень концентрации: %s\n", getFocusLevel(day))
}

func printTopic(day int) {
	fmt.Printf("\n%s\n", todayTopic.Title)
	fmt.Printf("%s | %s | Приоритет: %d/3\n",
		todayTopic.Category, todayTopic.Complexity, todayTopic.Priority)

	understanding := (day % 10) + 1
	emoji := getUnderstandingEmoji(understanding)

	fmt.Printf("\n%s Уровень понимания: %d/10\n", emoji, understanding)
	fmt.Printf("🎯 Цель: %d+ строк кода\n", todayTopic.MinLines)

	fmt.Println("\n📋 ЗАДАЧИ:")
	printList(todayTopic.Tasks, DISPLAY_LIMIT, true)
}

func printChallenges(today time.Time, go365Day int) {
	fmt.Println()
	challengeOrder := []string{"Go365", "100daysGo"}

	for _, name := range challengeOrder {
		ch := challenges[name]
		days := calcChallengeDays(name, ch, today, go365Day)
		percent := days * 100 / ch.TotalDays
		if percent > 100 {
			percent = 100
		}
		level := min(days/10+1, 10)

		fmt.Printf("%s\n", ch.Name)
		fmt.Printf("  День %d | Уровень %d\n", days, level)
		fmt.Printf("  %s\n\n", progressBar(percent))
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

func printFocus() {
	fmt.Println()
	for _, focus := range ecosystemFocus {
		fmt.Printf("%s\n", focus.Category)
		fmt.Printf("  %s\n", progressBar(focus.Progress))
		printList(focus.Skills, DISPLAY_LIMIT, false)
		fmt.Println()
	}
}

func printManifesto() {
	fmt.Println("\n❌ ЗАПРЕЩЕНО:")
	for _, item := range negativeItems {
		fmt.Printf("  × %s\n", item)
	}

	fmt.Println("\n✅ РАЗРЕШЕНО:")
	fmt.Println("  ✓ GO + КОД + ДОКУМЕНТАЦИЯ")
	fmt.Println("  ✓ ТЕХНИЧЕСКИЕ СТАТЬИ И КНИГИ")
	fmt.Println("  ✓ ОБСУЖДЕНИЯ ТЕХНИЧЕСКИХ ВОПРОСОВ")
	fmt.Println("  ✓ СОЗДАНИЕ ПРОЕКТОВ")
}

func printRules() {
	fmt.Println()
	for i, rule := range rules {
		fmt.Printf("  %d. %s\n", i+1, rule)
	}
}

func printFooter(day int) {
	fmt.Printf("\n💭 %s\n\n", quotes[day%len(quotes)])

	fmt.Println("🚀 GO ИЛИ НИЧЕГО")
	fmt.Println("   КОД ИЛИ НИЧЕГО")
	fmt.Println("   СИСТЕМНЫЙ ПОДХОД ИЛИ НИЧЕГО")
	fmt.Println()
	fmt.Println("🔥 НЕ ОТВЕЧАЙ НА ТРОЛЛЕЙ — ОТВЕЧАЙ КОДОМ")
	fmt.Println("   НЕ ОПРАВДЫВАЙСЯ — КОММИТЬ")
	fmt.Println("   НЕ ОТВЛЕКАЙСЯ — УГЛУБЛЯЙСЯ")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
