package main

import (
	"fmt"
	"time"
)

type ChallengeTracker struct {
	startDate    time.Time
	botStartTime time.Time
}

func NewChallengeTracker(botStartTime time.Time, startDateStr string) *ChallengeTracker {
	// Парсим дату начала из конфигурации
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		// Если ошибка парсинга, используем дату по умолчанию
		startDate = time.Date(2024, 10, 25, 0, 0, 0, 0, time.UTC)
	}

	return &ChallengeTracker{
		startDate:    startDate,
		botStartTime: botStartTime,
	}
}

func (ct *ChallengeTracker) GetCurrentDay() int {
	now := time.Now().UTC()
	days := int(now.Sub(ct.startDate).Hours() / 24)

	if days < 1 {
		return 1 // Минимальный день - первый
	}
	if days > 100 {
		return 100 // Максимальный день - сотый
	}
	return days
}

func (ct *ChallengeTracker) GetProgressMessage() string {
	currentDay := ct.GetCurrentDay()
	progress := float64(currentDay) / 100.0 * 100

	message := "🎯 *100daysGo Перезагрузка*\n\n"
	message += fmt.Sprintf("📅 *Текущий день:* %d из 100\n", currentDay)
	message += fmt.Sprintf("📊 *Прогресс:* %.1f%%\n\n", progress)

	// Визуальный прогресс-бар (реальные данные)
	bar := ct.getProgressBar(currentDay)
	message += fmt.Sprintf("`%s`\n\n", bar)

	message += fmt.Sprintf("⏰ *Запуск бота:* %s\n", ct.botStartTime.Format("15:04:05"))
	message += fmt.Sprintf("📅 *Текущая дата:* %s", time.Now().Format("02.01.2006"))

	return message
}

func (ct *ChallengeTracker) getProgressBar(currentDay int) string {
	width := 20
	// Реальное количество заполненных символов (каждый = 5 дней)
	completed := (currentDay * width) / 100
	if completed > width {
		completed = width
	}

	bar := ""
	for i := 0; i < width; i++ {
		if i < completed {
			bar += "█" // Заполненные дни
		} else {
			bar += "░" // Оставшиеся дни
		}
	}
	return bar
}
