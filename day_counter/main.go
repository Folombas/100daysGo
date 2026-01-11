package main

import (
	"fmt"
	"time"
)

// Конфигурация
const (
	DATE_FORMAT = "02.01.2006"
	BAR_WIDTH   = 25
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

// Основная функция
func main() {
	// Инициализация данных
	challenges := map[string]Challenge{
		"Go365":     {"Go365", "2026-01-01", 365},
		"100daysGo": {"100daysGo", "2025-11-03", 100},
	}

	today := time.Now()
	go365Start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	go365Day := calcDaysBetween(go365Start, today)

	// Тема дня
	todayTopic := DailyTopic{
		Title:      "Generics in Go Programming Language: Type Constraints",
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

	// Фокус на экосистеме
	ecosystemFocus := []EcoSystemFocus{
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

	// Вывод
	printHeader(today, go365Day)
	printTopic(&todayTopic, go365Day)
	printChallengesProgress(challenges, today, go365Day)
	printEcosystemFocus(ecosystemFocus)
	printFocusManifesto()
	printDevRules()
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

// ========== ВЫВОД ==========

func printHeader(date time.Time, day int) {
	fmt.Println()
	fmt.Println("🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO")
	fmt.Println("═" + repeatString("═", 48))
	fmt.Printf("📅 %s | %d День %d абсолютного фокуса\n",
		date.Format(DATE_FORMAT), getDaySymbol(day), day)
	fmt.Printf("🧠 Уровень концентрации: %s\n\n", getFocusLevel(day))
}

func getDaySymbol(day int) string {
	// Используем кружочки для дней
	if day <= 30 {
		return "①"
	}
	// Для больших чисел показываем номер
	return fmt.Sprintf("%d", day)
}

func printTopic(topic *DailyTopic, day int) {
	fmt.Println("📚 ТЕМА ДНЯ")
	fmt.Println()
	fmt.Printf("%s\n", topic.Title)
	fmt.Printf("%s | %s | Приоритет: %d/3\n",
		topic.Category, topic.Complexity, topic.Priority)

	understanding := (day % 10) + 1
	emoji := getUnderstandingEmoji(understanding)

	fmt.Printf("\n%s Уровень понимания: %d/10\n", emoji, understanding)
	fmt.Printf("🎯 Цель: %d+ строк кода\n\n", topic.MinLines)

	fmt.Println("📋 ЗАДАЧИ:")
	for i, task := range topic.Tasks {
		fmt.Printf("  %d. %s\n", i+1, task)
		if i == 2 && len(topic.Tasks) > 4 {
			remaining := len(topic.Tasks) - 3
			fmt.Printf("     ...и ещё %d задач\n", remaining)
			break
		}
	}

	fmt.Println("\n" + repeatString("─", 50))
}

func printChallengesProgress(challenges map[string]Challenge, today time.Time, go365Day int) {
	fmt.Println("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ")
	fmt.Println()

	// Сначала Go365, потом 100daysGo
	order := []string{"Go365", "100daysGo"}

	for _, name := range order {
		ch := challenges[name]
		days := calcChallengeDays(name, ch, today, go365Day)
		percent := days * 100 / ch.TotalDays
		if percent > 100 {
			percent = 100
		}
		level := min(days/10+1, 10)

		fmt.Printf("%s\n", name)
		fmt.Printf("  День %d | Уровень %d\n", days, level)
		printSimpleProgressBar(percent)
		fmt.Println()
	}

	fmt.Println(repeatString("─", 50))
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

func printSimpleProgressBar(percent int) {
	fmt.Print("  [")
	filled := percent * BAR_WIDTH / 100
	for i := 0; i < BAR_WIDTH; i++ {
		if i < filled {
			fmt.Print("█")
		} else {
			fmt.Print("░")
		}
	}
	fmt.Printf("] %d%%", percent)
}

func printEcosystemFocus(focuses []EcoSystemFocus) {
	fmt.Println("🎯 ФОКУС НА ЭКОСИСТЕМЕ GO")
	fmt.Println()

	for _, focus := range focuses {
		fmt.Printf("%s\n", focus.Category)
		printSimpleProgressBar(focus.Progress)
		fmt.Println()

		// Показываем первые 3 навыка
		for i := 0; i < len(focus.Skills) && i < 3; i++ {
			fmt.Printf("  • %s\n", focus.Skills[i])
		}
		if len(focus.Skills) > 3 {
			fmt.Printf("    +%d more...\n", len(focus.Skills)-3)
		}
		fmt.Println()
	}

	fmt.Println(repeatString("─", 50))
}

func printFocusManifesto() {
	fmt.Println("📜 МАНИФЕСТ ПОЛНОГО ПОГРУЖЕНИЯ")
	fmt.Println()

	fmt.Println("❌ ЗАПРЕЩЕНО:")
	negativeItems := []string{
		"БАРЫ / КЛУБЫ",
		"ФИЛЬМЫ / СЕРИАЛЫ",
		"ВИДЕОИГРЫ",
		"БЕССМЫСЛЕННЫЙ SCROLL",
		"ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ",
		"СОЦИАЛЬНЫЕ СЕТИ",
		"НОВОСТНЫЕ ЛЕНТЫ",
	}

	for _, item := range negativeItems {
		fmt.Printf("  × %s\n", item)
	}

	fmt.Println("\n✅ РАЗРЕШЕНО:")
	fmt.Println("  ✓ GO + КОД + ДОКУМЕНТАЦИЯ")
	fmt.Println("  ✓ ТЕХНИЧЕСКИЕ СТАТЬИ И КНИГИ")
	fmt.Println("  ✓ ОБСУЖДЕНИЯ ТЕХНИЧЕСКИХ ВОПРОСОВ")
	fmt.Println("  ✓ СОЗДАНИЕ ПРОЕКТОВ")
	fmt.Println()

	fmt.Println(repeatString("─", 50))
}

func printDevRules() {
	fmt.Println("⚡ ПРАВИЛА GO-РАЗРАБОТЧИКА")
	fmt.Println()

	rules := []string{
		"1. Код > Оправданий",
		"2. Коммит > Скроллинга",
		"3. Документация > Догадок",
		"4. Тесты > Уверенности",
		"5. Простота > Умности",
		"6. Практика > Теории",
		"7. Git > Памяти",
		"8. Docker > 'У меня работает'",
	}

	for _, rule := range rules {
		fmt.Printf("  %s\n", rule)
	}

	fmt.Println()
	fmt.Println(repeatString("─", 50))
}

func printFooter(day int) {
	fmt.Println()

	quotes := []string{
		"«Системный подход к обучению создаёт системного разработчика»",
		"«Каждый коммит — шаг к мастерству в Go и его экосистеме»",
		"«Инвестиции в знания Go, Linux и DevOps окупаются экспоненциально»",
		"«1 час целенаправленного кода на Go стоит 10 часов поверхностного изучения»",
		"«Экосистема — это не только язык, но и окружение, инструменты, практики»",
		"«Контейнеризация знания: Docker для кода, Kubernetes для карьеры»",
		"«Базы данных — память приложения, алгоритмы — его интеллект»",
		"«Полное погружение в Go: от горутин до продакшн-деплоя»",
	}

	fmt.Printf("💭 %s\n\n", quotes[day%len(quotes)])

	fmt.Println("🚀 GO ИЛИ НИЧЕГО")
	fmt.Println("   КОД ИЛИ НИЧЕГО")
	fmt.Println("   СИСТЕМНЫЙ ПОДХОД ИЛИ НИЧЕГО")
	fmt.Println()
	fmt.Println("🔥 НЕ ОТВЕЧАЙ НА ТРОЛЛЕЙ — ОТВЕЧАЙ КОДОМ")
	fmt.Println("   НЕ ОПРАВДЫВАЙСЯ — КОММИТЬ")
	fmt.Println("   НЕ ОТВЛЕКАЙСЯ — УГЛУБЛЯЙСЯ")
}

// Вспомогательная функция для повторения строки
func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// Вспомогательная функция для Go 1.20 и ниже
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
