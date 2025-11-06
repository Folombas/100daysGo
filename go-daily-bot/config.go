package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken    string
	ChallengeStart string
	AdminID     int64
	DebugMode   bool
}

func LoadConfig() (*Config, error) {
	// Загружаем .env файл (если существует)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using environment variables")
	}

	// Получаем конфигурацию из переменных окружения
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("❌ BOT_TOKEN is required")
	}

	challengeStart := os.Getenv("CHALLENGE_START")
	if challengeStart == "" {
		challengeStart = "2025-10-25" // Дата по умолчанию
	}

	adminID, _ := strconv.ParseInt(os.Getenv("ADMIN_ID"), 10, 64)
	debugMode := os.Getenv("DEBUG") == "true"

	return &Config{
		BotToken:      botToken,
		ChallengeStart: challengeStart,
		AdminID:       adminID,
		DebugMode:     debugMode,
	}, nil
}
