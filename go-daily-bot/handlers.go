package main

import (
	"log"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, tracker *ChallengeTracker, sysInfo *SystemInfo, config *Config) {
	log.Printf("👤 %s: %s", message.From.UserName, message.Text)

	var response string

	switch message.Text {
	case "/start", "/help":
		response = getWelcomeMessage()
	case "/progress", "/day":
		response = tracker.GetProgressMessage()
	case "/system", "/info":
		response = sysInfo.GetSystemMessage()
	case "/motivation":
		response = getMotivationMessage(tracker.GetCurrentDay())
	case "/config":
		// Только для администратора
		if message.From.ID == config.AdminID {
			response = getConfigInfo(config)
		} else {
			response = "❌ Доступ запрещён"
		}
	default:
		response = "🤔 Используй команды:\n/start - Начать\n/progress - Прогресс\n/system - Инфо о системе\n/motivation - Мотивация"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ParseMode = "Markdown"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("❌ Ошибка отправки: %v", err)
	}
}

func getConfigInfo(config *Config) string {
	return fmt.Sprintf(`⚙️ *Конфигурация бота:*

🤖 Режим отладки: %v
📅 Дата начала: %s
👤 Admin ID: %d
`, config.DebugMode, config.ChallengeStart, config.AdminID)
}

// Остальные функции без изменений...
func getWelcomeMessage() string {
	return `🚀 *Добро пожаловать в 100daysGo Перезагрузка!*

Я твой помощник в 100-дневном челлендже по изучению Go!

📋 *Доступные команды:*
/progress - Текущий день и прогресс
/system - Информация о системе
/motivation - Мотивационное сообщение
/help - Справка

Каждый день приближает тебя к цели! 💪`
}

func getMotivationMessage(day int) string {
	motivations := []string{
		"🔥 Ты делаешь то, о чем другие только мечтают!",
		"💪 Каждая строка кода - шаг к лучшей версии себя!",
		"🚀 Сегодняшние усилия - завтрашние навыки!",
		"🎯 Помни: эксперты когда-то тоже были новичками!",
		"🌟 Ты заменяешь временные удовольствия на вечные навыки!",
	}

	index := day % len(motivations)
	return fmt.Sprintf("📅 День %d:\n\n%s", day, motivations[index])
}
