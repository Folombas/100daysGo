package main

import (
	"fmt"
	"runtime"
	"time"
)

type SystemInfo struct {
	OS           string
	GoVersion    string
	Architecture string
	NumCPU       int
	StartTime    time.Time
}

func NewSystemInfo(startTime time.Time) *SystemInfo {
	return &SystemInfo{
		OS:           runtime.GOOS,
		GoVersion:    runtime.Version(),
		Architecture: runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
		StartTime:    startTime,
	}
}

func (si *SystemInfo) GetSystemMessage() string {
	uptime := time.Since(si.StartTime)

	message := "💻 *Информация о системе:*\n\n"
	message += fmt.Sprintf("⚙️  *ОС:* %s\n", si.OS)
	message += fmt.Sprintf("🚀 *Архитектура:* %s\n", si.Architecture)
	message += fmt.Sprintf("🔢 *Процессоры:* %d\n", si.NumCPU)
	message += fmt.Sprintf("🐹 *Версия Go:* %s\n", si.GoVersion)
	message += fmt.Sprintf("⏰ *Аптайм:* %s\n", si.formatUptime(uptime))
	message += fmt.Sprintf("🕒 *Время запуска бота:* %s", si.StartTime.Format("02.01.2006 15:04:05"))

	return message
}

func (si *SystemInfo) formatUptime(uptime time.Duration) string {
	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%d ч %d мин %d сек", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d мин %d сек", minutes, seconds)
	} else {
		return fmt.Sprintf("%d сек", seconds)
	}
}
