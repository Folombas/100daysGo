package main

import (
	"fmt"
	"time"
)

type ChallengeTracker struct {
	startDate    time.Time
	botStartTime time.Time
}

func NewChallengeTracker(botStartTime time.Time) *ChallengeTracker {
	// Дата начала челленджа - 25.10.2025
	startDate := time.Date(2025, 10, 25, 0, 0, 0, 0, time.UTC)
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
	return days + 1 // +1 потому что первый день считается как день 1
}

func (ct *ChallengeTracker) GetProgressMessage() string {
	currentDay := ct.GetCurrentDay()
	progress := float64(currentDay) / 100.0 * 100

	message := fmt.Sprintf("🎯 *100daysGo Перезагрузка*\n\n", currentDay)
	message += fmt.Sprintf("📅 *Текущий день:* %d из 100\n", currentDay)
	message += fmt.Sprintf("📊 *Прогресс:* %.1f%%\n\n", progress)

	// Визуальный прогресс-бар
	bar := ct.getProgressBar(currentDay)
	message += fmt.Sprintf("`%s`\n\n", bar)

	message += fmt.Sprintf("⏰ *Время запуска бота:* %s\n", ct.botStartTime.Format("02.01.2006 15:04:05"))
	message += fmt.Sprintf("🕒 *Текущее время:* %s", time.Now().Format("02.01.2006 15:04:05"))

	return message
}

func (ct *ChallengeTracker) getProgressBar(currentDay int) string {
	width := 20
	completed := (currentDay * width) / 100
	if completed > width {
		completed = width
	}

	bar := ""
	for i := 0; i < width; i++ {
		if i < completed {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}
