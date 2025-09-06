package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Получаем токен бота из переменных окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен в .env файле")
	}

	// Настраиваем опции бота
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// Создаем экземпляр бота
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Регистрируем обработчики команд
	bot.Handle("/start", handleStart)
	bot.Handle("/help", handleHelp)
	bot.Handle("/go", handleGoFact)
	bot.Handle("/code", handleCodeExample)
	bot.Handle("/links", handleLinks)
	bot.Handle("/weather", handleWeather)
	
	// Обработчик текстовых сообщений
	bot.Handle(telebot.OnText, handleText)

	// Запускаем бота
	fmt.Println("🤖 Бот запущен и готов к работе...")
	bot.Start()
}