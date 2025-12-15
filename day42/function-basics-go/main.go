package main

import (
	"fmt"
	"strings"
	"time"
)

func calculateDevLevel(years int) string {
	switch {
	case years < 1:
		return "🌱 Intern: Учи синтаксис, пока мама варит рис"
	case years < 3:
		return "🚀 Junior: Твой первый коммит в продакшене!"
	case years < 5:
		return "🔥 Middle: Забыл, что такое 'курьерская сумка'"
	case years < 8:
		return "💎 Senior: Твои решения влияют на миллионы пользователей"
	default:
		return "👑 Team Lead: Нанимаешь бывших курьеров в команду"
	}
}

func isCourierDay(t time.Time) string {
	if t.Month() == time.December && t.Day() == 17 {
		return "🎁 СЕГОДНЯ ДЕНЬ КУРЬЕРА! Спасибо за 20 лет в профессии.\n   💡 Совет: Сегодня можно устроить 'последний рабочий день' курьера в своём воображении."
	}
	if t.Month() == time.December && t.Day() < 17 {
		days := 17 - t.Day()
		return fmt.Sprintf("⏳ До Дня Курьера (%s) осталось %d дней. Выдержи!", t.Format("02.01"), days)
	}
	return "📅 День Курьера: 17 декабря. Отметим через год — уже как разработчики!"
}

func dreamCostCalculator(monthlySalary int) (int, string) {
	cardPrice := 120000
	months := cardPrice / monthlySalary
	status := "💤 Спи спокойно"
	if monthlySalary < 50000 {
		status = "⚡ Учи Go! Каждый пропущенный урок = +1 день к ожиданию карты"
	} else if monthlySalary >= 200000 {
		status = "🎮 GeForce 5060 уже в корзине! Оформляй заказ после первого рабочего дня"
	}
	return months, status
}

func main() {
	now := time.Date(2025, time.December, 15, 18, 0, 0, 0, time.Local)

	fmt.Println("❄️  15 ДЕКАБРЯ 2025: ДЕНЬ КУРЬЕРА, КОТОРЫЙ МЕЧТАЕТ О СЕРВЕРАХ")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Printf("👨‍💻  ТВОЙ УРОВЕНЬ: %s\n", calculateDevLevel(0))
	fmt.Printf("📦  %s\n", isCourierDay(now))

	months, status := dreamCostCalculator(80000)
	fmt.Printf("\n💻  МЕЧТА (GeForce 5060):\n   Накопление: %d мес. | Статус: %s\n", months, status)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🔄  ТВОЙ ПЛАН:")
	fmt.Println("   func switchCareer() string {")
	fmt.Println("       deleteGamesAndSeries()")
	fmt.Println("       readBookDaily()")
	fmt.Println("       return \"Go-разработчик через 6 месяцев!\"")
	fmt.Println("   }")

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🔥  ГЛАВНОЕ ПРАВИЛО ДНЯ:")
	fmt.Println("   if hasDream {")
	fmt.Println("       ignoreColdStreets()")
	fmt.Println("       writeCodeEveryDay()")
	fmt.Println("   }")
}
