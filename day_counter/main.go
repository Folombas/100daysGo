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

func main() {
	challenges := map[string]Challenge{
		"100daysGo": {"100daysGo", "2025-11-03", 100},
		"Go365":     {"Go365", "2026-01-01", 365},
	}

	currentDate := time.Now()
	jan1 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	go365Day := calculateGo365Day(currentDate, jan1)

	printHeader(currentDate, go365Day)
	printChallengesProgress(challenges, currentDate, go365Day)
	printFocusManifesto()
	printAllowedActivities()
	printDailyTopic("Generics in Go Programming Language: Type Inference", go365Day)
	printFooter(go365Day)
}

func calculateGo365Day(currentDate, jan1 time.Time) int {
	days := int(currentDate.Sub(jan1).Hours()/24) + 1
	if days < 1 {
		days = 1
	}
	return days
}

func calculateDays(startDate string, currentDate time.Time) int {
	start, _ := time.Parse("2006-01-02", startDate)
	days := int(currentDate.Sub(start).Hours() / 24)
	if days < 0 {
		days = 0
	}
	return days + 1
}

func printHeader(date time.Time, go365Day int) {
	fmt.Printf("\n🚫 НИКАКИХ РАЗВЛЕЧЕНИЙ — ТОЛЬКО GO\n")
	fmt.Println("════════════════════════════════════")
	fmt.Printf("📅 %s | 🔥 День %d абсолютного фокуса\n",
		date.Format("02.01.2006"), go365Day)
	fmt.Printf("🧠 Уровень концентрации: %s\n\n", getFocusLevel(go365Day))
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

func printChallengesProgress(challenges map[string]Challenge, currentDate time.Time, go365Day int) {
	fmt.Println("📊 ПРОГРЕСС ЧЕЛЛЕНДЖЕЙ")

	for name, challenge := range challenges {
		days, percent, level := calculateChallengeProgress(name, challenge, currentDate, go365Day)
		fmt.Printf("\n▸ %s: День %d | Ур.%d\n", name, days, level)
		printProgressBar(percent)
	}
}

func calculateChallengeProgress(name string, challenge Challenge, currentDate time.Time, go365Day int) (days, percent, level int) {
	if name == "Go365" {
		days = go365Day
	} else {
		days = calculateDays(challenge.StartDate, currentDate)
	}

	percent = days * 100 / challenge.TotalDays
	if percent > 100 {
		percent = 100
	}

	level = days/10 + 1
	if level > 10 {
		level = 10
	}

	return days, percent, level
}

func printProgressBar(percent int) {
	const width = 30
	filled := percent * width / 100

	fmt.Print("   [")
	for i := 0; i < width; i++ {
		if i < filled {
			fmt.Print("█")
		} else {
			fmt.Print("░")
		}
	}
	fmt.Printf("] %d%%\n", percent)
}

func printFocusManifesto() {
	fmt.Println("\n📜 МАНИФЕСТ ФОКУСА")
	printBoxedItems([]string{
		"БАРЫ/КЛУБЫ      → ❌ НЕТ",
		"ФИЛЬМЫ/СЕРИАЛЫ → ❌ НЕТ",
		"ВИДЕОИГРЫ      → ❌ НЕТ",
		"SCROLL         → ❌ НЕТ",
		"ПУСТЫЕ РАЗГОВОРЫ С ТРОЛЛЯМИ → ❌ НЕТ",
	})
	fmt.Println("   ✅ РАЗРЕШЕНО: GO + КОД + ДОКУМЕНТАЦИЯ")
}

func printAllowedActivities() {
	fmt.Println("\n🎯 ФОКУС НА РАЗВИТИИ")
	printBoxedItems([]string{
		"ПИСАТЬ КОД                     → ✅ ДА",
		"ОСНОВЫ LINUX                   → ✅ ДА",
		"TERMINAL                       → ✅ ДА",
		"DOCKER                         → ✅ ДА",
		"АЛГОРИТМЫ И СТРУКТУРЫ ДАННЫХ  → ✅ ДА",
		"ЧИТАТЬ ДОКУМЕНТАЦИЮ            → ✅ ДА",
		"СОЗДАВАТЬ ПРОЕКТЫ             → ✅ ДА",
		"РЕШАТЬ ЗАДАЧИ НА LEETCODE      → ✅ ДА",
		"ИЗУЧАТЬ АРХИТЕКТУРУ ПО         → ✅ ДА",
		"ПИСАТЬ ТЕСТЫ                   → ✅ ДА",
		"РАБОТАТЬ С GIT                 → ✅ ДА",
		"ЧИТАТЬ ЧУЖОЙ КОД               → ✅ ДА",
	})
	fmt.Println("   🔥 НАПИСАТЬ КОД ЛУЧШЕ, ЧЕМ НАПИСАТЬ ОПРАВДАНИЯ ТРОЛЛЯМ")
}

func printBoxedItems(items []string) {
	fmt.Println("   ┌─────────────────────────────────────────┐")
	for _, item := range items {
		fmt.Printf("   │ %-40s │\n", item)
	}
	fmt.Println("   └─────────────────────────────────────────┘")
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
		"«Тролли кормятся вниманием. Лишай их питания — пиши код»",
		"«Лучший ответ троллю — твой следующий коммит»",
	}

	fmt.Printf("💬 %s\n", quotes[day%len(quotes)])
	fmt.Println("\n🚀 GO ИЛИ НИЧЕГО. КОД ИЛИ НИЧЕГО.")
	fmt.Println("   🔥 НЕ ОТВЕЧАЙ НА ТРОЛЛЕЙ — ОТВЕЧАЙ КОДОМ")
}
