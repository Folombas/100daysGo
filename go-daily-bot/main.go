package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Получаем токен бота из переменных окружения
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Panic("❌ BOT_TOKEN not set in environment variables")
	}

	// Инициализируем бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("🤖 Authorized on account %s", bot.Self.UserName)

	// Время запуска бота
	startTime := time.Now()

	// Инициализируем трекер челленджа
	tracker := NewChallengeTracker(startTime)

	// Настраиваем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("🚀 Бот запущен и готов к работе!")

	// Обрабатываем сообщения
	for update := range updates {
		if update.Message == nil {
			continue
		}

		HandleMessage(bot, update.Message, tracker)
	}
}
