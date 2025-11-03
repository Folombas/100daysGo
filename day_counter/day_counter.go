package main

import (
	"fmt"
	"math"
	"time"
)

type ChallengeStats struct {
	TotalDays         int
	DaysCompleted     int
	DaysRemaining     int
	ProgressPercent   float64
	CurrentStreak     int
	LongestStreak     int
	LastStudyDate     time.Time
	ProductivityScore int
}

type PersonalGrowth struct {
	GamingAvoided       int
	AdultContentAvoided int
	StudyHours          float64
	SkillsLearned       []string
	Mood                string
}

func main() {
	// НАЧАЛО НОВОГО ЧЕЛЛЕНДЖА - 3 ноября 2025
	startDate := time.Date(2025, time.November, 3, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()

	// Расчет дней (zero-based)
	daysPassed := int(currentDate.Sub(startDate).Hours() / 24)
	currentDay := daysPassed // Day0, Day1, etc.

	// Статистика челленджа
	stats := calculateStats(startDate, currentDate, daysPassed)

	// Личный рост
	growth := PersonalGrowth{
		GamingAvoided:       daysPassed * 2, // Примерная статистика
		AdultContentAvoided: daysPassed * 3,
		StudyHours:          float64(daysPassed) * 1.5,
		SkillsLearned:       []string{"Go basics", "Functions", "Packages", "Concurrency"},
		Mood:                getMood(daysPassed),
	}

	fmt.Println("🎯 100daysGo: ПЕРЕЗАГРУЗКА")
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("👤 Участник: Гоша, 37 лет | СДВГ+ОКР\n")
	fmt.Printf("📅 Старт: %s\n", startDate.Format("02.01.2006"))
	fmt.Printf("📊 Сегодня: %s\n", currentDate.Format("02.01.2006"))
	fmt.Printf("🎮 Параллельный челлендж: Ноябрь-Небричабрь ✅\n")
	fmt.Println()

	// Основная информация о дне
	if daysPassed < 0 {
		fmt.Printf("⏳ До начала челленджа: %d дней\n", int(math.Abs(float64(daysPassed))))
		fmt.Printf("🎯 Стартуем: %s\n", startDate.Format("02.01.2006"))
	} else {
		fmt.Printf("🔥 ДЕНЬ ЧЕЛЛЕНДЖА: Day%d\n", currentDay)
		fmt.Printf("📈 Прогресс: %d/%d дней (%.1f%%)\n",
			stats.DaysCompleted, stats.TotalDays, stats.ProgressPercent)
		fmt.Printf("⏱️  Осталось дней: %d\n", stats.DaysRemaining)
		fmt.Printf("🔥 Текущая серия: %d дней\n", stats.CurrentStreak)
		fmt.Println()

		// Мотивационное сообщение
		printDailyMessage(currentDay, stats, growth)
	}

	fmt.Println()
	fmt.Println("💪 СЕГОДНЯШНИЕ ЦЕЛИ:")
	fmt.Println("   • Изучить новую тему Go")
	fmt.Println("   • Написать код и сделать коммит")
	fmt.Println("   • Избегать цифровых наркотиков")
	fmt.Println("   • Сделать шаг к финансовой свободе")

	fmt.Println()
	fmt.Println("🌟 СТАТИСТИКА ЛИЧНОГО РОСТА:")
	fmt.Printf("   🎮 Игр избежано: ~%d сессий\n", growth.GamingAvoided)
	fmt.Printf("   🔞 Контента избежано: ~%d раз\n", growth.AdultContentAvoided)
	fmt.Printf("   📚 Часов изучения: ~%.1f часов\n", growth.StudyHours)
	fmt.Printf("   😊 Настроение: %s\n", growth.Mood)

	fmt.Println()
	fmt.Println("🎯 ДОЛГОСРОЧНАЯ ЦЕЛЬ:")
	fmt.Printf("   💰 Устроиться Go-разработчиком до: %s\n",
		startDate.Add(100*24*time.Hour).Format("02.01.2006"))
	fmt.Println("   🏠 Перестать беспокоить маму о деньгах")
	fmt.Println("   🚀 Начать карьеру в IT")

	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("💡 Помни: Каждый день кода - это шаг")
	fmt.Println("   от цифровой зависимости к цифровой свободе!")
}

func calculateStats(startDate, currentDate time.Time, daysPassed int) ChallengeStats {
	totalDays := 100
	daysCompleted := daysPassed + 1
	if daysPassed < 0 {
		daysCompleted = 0
	}

	progressPercent := float64(daysCompleted) / float64(totalDays) * 100
	daysRemaining := totalDays - daysCompleted
	if daysRemaining < 0 {
		daysRemaining = 0
	}

	// Простая логика для серий (в реальном приложении нужно хранить историю)
	currentStreak := daysCompleted
	if daysPassed < 0 {
		currentStreak = 0
	}

	return ChallengeStats{
		TotalDays:         totalDays,
		DaysCompleted:     daysCompleted,
		DaysRemaining:     daysRemaining,
		ProgressPercent:   progressPercent,
		CurrentStreak:     currentStreak,
		LongestStreak:     currentStreak,
		LastStudyDate:     currentDate,
		ProductivityScore: daysCompleted * 10,
	}
}

func getMood(daysPassed int) string {
	if daysPassed < 0 {
		return "Ожидание старта 🎯"
	}

	moods := []string{
		"Энтузиазм старта 🚀",
		"Формирование привычки 💪",
		"Стабильный прогресс 📈",
		"Преодоление трудностей 🏔️",
		"Уверенность в себе 😎",
	}

	index := daysPassed / 20
	if index >= len(moods) {
		index = len(moods) - 1
	}
	return moods[index]
}

func printDailyMessage(day int, stats ChallengeStats, growth PersonalGrowth) {
	messages := map[int]string{
		0:   "🎉 СТАРТ! Ты начал путь к свободе от зависимостей и бедности!",
		1:   "🔥 Первый шаг сделан! Помни: игры и сериалы - это цифровые наркотики.",
		7:   "📅 Неделя без цифровых наркотиков! Ты становишься сильнее.",
		14:  "💪 Две недели! Мама уже заметила твоё упорство?",
		30:  "🎯 Месяц пути! Ты уже знаешь больше Go, чем 80% 'гуру' из ютуба.",
		50:  "🚀 Полпути! Представь: через 50 дней ты сможешь идти на собеседование.",
		75:  "🌟 75 дней! Ты уже не тот человек, что начинал этот путь.",
		99:  "🏁 Завтра - 100 дней! Готовь резюме, ты стал разработчиком.",
		100: "🎊 ФИНИШ! Теперь ты Go-разработчик. Время менять жизнь!",
	}

	if msg, exists := messages[day]; exists {
		fmt.Printf("💫 %s\n", msg)
	} else {
		// Случайные мотивационные сообщения
		dailyMessages := []string{
			"Каждая строка кода - это кирпичик в фундаменте твоего нового будущего.",
			"Сегодня ты выбрал код вместо игр. Завтра выберешь карьеру вместо бедности.",
			"ОКР и СДВГ - не оправдания, а особенности. Используй их как суперсилу!",
			"Мама будет гордиться, когда увидит твою первую зарплату разработчика.",
			"Цифровые наркотики украли твоё прошлое. Go вернёт тебе будущее.",
			"37 лет - идеальный возраст для перезагрузки. Опыт жизни + мудрость = успех.",
		}

		messageIndex := day % len(dailyMessages)
		fmt.Printf("💡 %s\n", dailyMessages[messageIndex])
	}

	// Специальные сообщения для ключевых моментов
	if day == 0 {
		fmt.Println()
		fmt.Println("🎯 СЕГОДНЯШНИЙ ПЛАН:")
		fmt.Println("   1. Настроить окружение разработки")
		fmt.Println("   2. Создать первый модуль Go")
		fmt.Println("   3. Написать коммит с описанием твоих целей")
		fmt.Println("   4. Гордиться собой - ты начал!")
	}

	// Прогресс по финансовой цели
	if day > 0 && day%10 == 0 {
		potentialSalary := 80000 + (day * 1000) // Рублей в месяц
		fmt.Printf("💰 Через %d дней ты сможешь зарабатывать ~%d руб/мес\n",
			stats.DaysRemaining, potentialSalary)
	}
}
