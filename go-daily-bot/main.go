package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Загружаем конфигурацию
	config, err := LoadConfig()
	if err != nil {
		log.Panicf("❌ Failed to load config: %v", err)
	}

	// Инициализируем бота
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.DebugMode
	log.Printf("🤖 Authorized on account %s", bot.Self.UserName)

	// Время запуска бота
	startTime := time.Now()

	// Инициализируем трекер челленджа
	tracker := NewChallengeTracker(startTime, config.ChallengeStart)

	sysInfo := NewSystemInfo(startTime)

	// Настраиваем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("🚀 Бот запущен и готов к работе!")
	log.Printf("📅 Дата начала челленджа: %s", config.ChallengeStart)

	// Обрабатываем сообщения
	for update := range updates {
		if update.Message == nil {
			continue
		}

		HandleMessage(bot, update.Message, tracker, sysInfo, config)
	}
}
