package main

import (
	"fmt"
	"time"
)

type Challenge struct {
	Name      string
	StartDate string
	TotalDays int
}

type Progress struct {
	DaysCount int
	Level     int
	TotalXP   int
}

func main() {
	challenges := map[string]Challenge{
		"100daysGo": {"100daysGo", "2025-11-03", 100},
		"Go365":     {"Go365", "2026-01-01", 365},
	}

	currentDate := time.Now()
	fmt.Printf("\n🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO\n")
	fmt.Println("════════════════════════════════════")

	// Рассчитываем день от начала года (1 января = день 1)
	jan1 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	go365Day := int(currentDate.Sub(jan1).Hours()/24) + 1
	if go365Day < 1 {
		go365Day = 1
	}

	fmt.Printf("📅 %s | 🔥 День %d абсолютного фокуса\n",
		currentDate.Format("02.01.2006"), go365Day)
	fmt.Printf("🧠 Уровень концентрации: %s\n\n", getFocusLevel(go365Day))

	// Прогресс по челленджам
	fmt.Println("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ")
	for name, challenge := range challenges {
		days := calculateDays(challenge.StartDate, currentDate)
		percent := days * 100 / challenge.TotalDays
		if percent > 100 {
			percent = 100
		}

		// Для Go365 используем исправленный расчёт
		if name == "Go365" {
			days = go365Day
			percent = days * 100 / 365
		}

		level := days/10 + 1
		if level > 10 {
			level = 10
		}

		fmt.Printf("\n▸ %s: День %d | Ур.%d\n", name, days, level)
		printProgressBar(percent)
	}

	printManifesto()
	printDailyTopic("Generics in Go Programming Language: Type Inference", go365Day)
	printFooter(go365Day)
}

func calculateDays(startDate string, currentDate time.Time) int {
	start, _ := time.Parse("2006-01-02", startDate)
	days := int(currentDate.Sub(start).Hours() / 24)
	if days < 0 {
		days = 0
	}
	return days + 1 // +1 чтобы первый день был 1, а не 0
}

func getFocusLevel(days int) string {
	switch {
	case days >= 30:
		return "🚀 КОСМИЧЕСКАЯ КОНЦЕНТРАЦИЯ"
	case days >= 20:
		return "🔥 ПОЛНЫЙ ФОКУС"
	case days >= 10:
		return "⚡ ВЫСОКАЯ СКОРОСТЬ"
	default:
		return "🌱 НАЧАЛО ПУТИ"
	}
}

func printProgressBar(percent int) {
	width := 30
	filled := percent * width / 100
	empty := width - filled

	fmt.Print("   [")
	for i := 0; i < filled; i++ {
		fmt.Print("█")
	}
	for i := 0; i < empty; i++ {
		fmt.Print("░")
	}
	fmt.Printf("] %d%%\n", percent)
}

func printManifesto() {
	fmt.Println("\n📜 МАНИФЕСТ ФОКУСА")
	fmt.Println("   ┌─────────────────────────────────────────┐")

	items := []string{
		"БАРЫ/КЛУБЫ      → ❌ НЕТ",
		"ФИЛЬМЫ/СЕРИАЛЫ → ❌ НЕТ",
		"ВИДЕОИГРЫ      → ❌ НЕТ",
		"SCROLL         → ❌ НЕТ",
		"ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ → ❌ НЕТ",
	}

	for _, item := range items {
		fmt.Printf("   │ %-40s │\n", item)
	}
	fmt.Println("   └─────────────────────────────────────────┘")
	fmt.Println("   ✅ РАЗРЕШЕНО: GO + КОД + ДОКУМЕНТАЦИЯ")
}

func printDailyTopic(topic string, day int) {
	fmt.Printf("\n📚 ТЕМА ДНЯ: %s\n", topic)

	tasks := []string{
		"• 100+ строк кода по теме",
		"• Прочитать документацию",
		"• Написать примеры",
		"• Создать конспект",
		"• Поделиться инсайтами",
	}

	for _, task := range tasks {
		fmt.Printf("   %s\n", task)
	}

	// Уровень понимания зависит от дня
	understanding := (day % 10) + 1
	fmt.Printf("\n   🎯 Уровень понимания: %d/10\n", understanding)
}

func printFooter(day int) {
	fmt.Println("\n════════════════════════════════════")

	quotes := []string{
		"«Каждый день кода — шаг к свободе»",
		"«Распыление создаёт дилетантов. Фокус — мастеров»",
		"«Мои тусовки — это коммиты. Мои друзья — это горутины»",
		"«1 час кода стоит 10 часов скроллинга»",
	}

	quoteIndex := day % len(quotes)
	fmt.Printf("💬 %s\n", quotes[quoteIndex])
	fmt.Println("\n🚀 GO ИЛИ НИЧЕГО. КОД ИЛИ НИЧЕГО.")
}
