package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// Конфигурация
const (
	DATE_FORMAT = "02.01.2006"
	BAR_WIDTH   = 30
	BOX_WIDTH   = 44 // Общая ширина рамки (включая границы)
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
		"Go365":     {"Go365", "2026-01-01", 365},
		"100daysGo": {"100daysGo", "2025-11-03", 100},
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

// Вычисляет видимую ширину строки с учетом кириллицы (2 символа на букву)
func visibleWidth(s string) int {
	width := 0
	for _, r := range s {
		if r <= 127 {
			width++ // ASCII символы
		} else if unicode.Is(unicode.Cyrillic, r) {
			width += 2 // Кириллица
		} else {
			width++ // Остальные символы (эмодзи, пунктуация)
		}
	}
	return width
}

// Создает строку с выравниванием по заданной ширине
func padToWidth(s string, width int) string {
	visible := visibleWidth(s)
	if visible >= width {
		return s
	}
	return s + strings.Repeat(" ", width-visible)
}

// ========== ВЫВОД ==========

func printHeader(date time.Time, day int) {
	fmt.Printf("\n🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO\n")
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("📅 %s | 🔥 День %d абсолютного фокуса\n", date.Format(DATE_FORMAT), day)
	fmt.Printf("🧠 Уровень концентрации: %s\n\n", getFocusLevel(day))
}

func printChallengesProgress(challenges map[string]Challenge, today time.Time, go365Day int) {
	fmt.Println("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ")
	fmt.Println()

	for name, ch := range challenges {
		days := calcChallengeDays(name, ch, today, go365Day)
		percent := days * 100 / ch.TotalDays
		if percent > 100 {
			percent = 100
		}
		level := min(days/10+1, 10)

		fmt.Printf("%s: День %d | Ур.%d\n", name, days, level)
		printProgressBar(percent)
		fmt.Println()
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

func printTopicBox(topic *DailyTopic, day int) {
	fmt.Println("📚 ТЕМА ДНЯ")
	printBoxTop()

	// Заголовок темы (может быть многострочным)
	titleLines := splitToLines(topic.Title, BOX_WIDTH-4)
	for _, line := range titleLines {
		fmt.Printf("│ %-40s │\n", padToWidth(line, BOX_WIDTH-4))
	}

	printBoxSeparator()

	// Уровень понимания
	understanding := (day % 10) + 1
	emoji := "🟢"
	switch {
	case understanding <= 3:
		emoji = "🔴"
	case understanding <= 7:
		emoji = "🟡"
	}

	fmt.Printf("│ %s │\n", padToWidth(fmt.Sprintf("%s Уровень понимания: %d/10", emoji, understanding), BOX_WIDTH-4))
	fmt.Printf("│ %s │\n", padToWidth(fmt.Sprintf("Цель: %d+ строк кода", topic.MinLines), BOX_WIDTH-4))

	printBoxSeparator()

	// Задачи
	for i, task := range topic.Tasks {
		if i < 3 {
			fmt.Printf("│ %s │\n", padToWidth(fmt.Sprintf("  • %s", task), BOX_WIDTH-4))
		}
	}

	if len(topic.Tasks) > 3 {
		fmt.Printf("│ %s │\n", padToWidth(fmt.Sprintf("  • ...и ещё %d задач", len(topic.Tasks)-3), BOX_WIDTH-4))
	}

	printBoxBottom()
	fmt.Println()
}

func printFocusManifesto() {
	fmt.Println("📜 МАНИФЕСТ ФОКУСА")
	printBoxTop()

	items := []string{
		"БАРЫ/КЛУБЫ                  → ❌ НЕТ",
		"ФИЛЬМЫ/СЕРИАЛЫ              → ❌ НЕТ",
		"ВИДЕОИГРЫ                   → ❌ НЕТ",
		"SCROLL                      → ❌ НЕТ",
		"ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ → ❌ НЕТ",
	}

	for _, item := range items {
		fmt.Printf("│ %s │\n", padToWidth(item, BOX_WIDTH-4))
	}

	printBoxBottom()
	fmt.Println("   ✅ РАЗРЕШЕНО: GO + КОД + ДОКУМЕНТАЦИЯ")
	fmt.Println()
}

func printAllowedActivities() {
	fmt.Println("🎯 ФОКУС НА РАЗВИТИИ")
	printBoxTop()

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

	for _, item := range items {
		fmt.Printf("│ %s │\n", padToWidth(item, BOX_WIDTH-4))
	}

	printBoxBottom()
	fmt.Println("   🔥 НАПИСАТЬ КОД ЛУЧШЕ, ЧЕМ НАПИСАТЬ ОПРАВДАНИЯ ТРОЛЛЯМ")
	fmt.Println()
}

// Вспомогательные функции для рисования рамок
func printBoxTop() {
	fmt.Printf("┌%s┐\n", strings.Repeat("─", BOX_WIDTH-2))
}

func printBoxBottom() {
	fmt.Printf("└%s┘\n", strings.Repeat("─", BOX_WIDTH-2))
}

func printBoxSeparator() {
	fmt.Printf("├%s┤\n", strings.Repeat("─", BOX_WIDTH-2))
}

// Разделяет длинную строку на несколько строк
func splitToLines(text string, maxWidth int) []string {
	var lines []string
	words := strings.Fields(text)

	if len(words) == 0 {
		return []string{""}
	}

	currentLine := words[0]

	for _, word := range words[1:] {
		if visibleWidth(currentLine+" "+word) <= maxWidth {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func printFooter(day int) {
	fmt.Println(strings.Repeat("═", 50))

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
