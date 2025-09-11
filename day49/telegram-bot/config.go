package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config структура для хранения конфигурации
type Config struct {
    BotToken string
}

// LoadConfig загружает конфигурацию
func LoadConfig() *Config {
    // Загружаем переменные из .env файла
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: .env file not found: %v", err)
    }

    // Получаем токен из переменных окружения
    token := os.Getenv("BOT_TOKEN")
    if token == "" {
        log.Fatal("Bot token not found. Set BOT_TOKEN environment variable")
    }

    return &Config{BotToken: token}
}
