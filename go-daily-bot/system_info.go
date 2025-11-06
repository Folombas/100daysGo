package main

import (
	"fmt"
	"runtime"
	"time"
)

type SystemInfo struct {
	OS          string
	GoVersion   string
	Architecture string
	NumCPU      int
}

func NewSystemInfo() *SystemInfo {
	return &SystemInfo{
		OS:          runtime.GOOS,
		GoVersion:   runtime.Version(),
		Architecture: runtime.GOARCH,
		NumCPU:      runtime.NumCPU(),
	}
}

func (si *SystemInfo) GetSystemMessage() string {
	message := "💻 *Информация о системе:*\n\n"
	message += fmt.Sprintf("⚙️  *ОС:* %s\n", si.OS)
	message += fmt.Sprintf("🚀 *Архитектура:* %s\n", si.Architecture)
	message += fmt.Sprintf("🔢 *Процессоры:* %d\n", si.NumCPU)
	message += fmt.Sprintf("🐹 *Версия Go:* %s\n", si.GoVersion)
	message += fmt.Sprintf("⏰ *Аптайм:* %s", si.getUptime())

	return message
}

func (si *SystemInfo) getUptime() string {
	// Для демонстрации - случайное время
	return "2 часа 15 минут"
}
