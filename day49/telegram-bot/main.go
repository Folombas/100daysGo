package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GameState хранит состояние игры для каждого пользователя
type GameState struct {
	SecretNumber int
	Attempts     int
	MaxAttempts  int
	IsPlaying    bool
}

var (
	games = make(map[int64]*GameState) // chatID -> GameState
	mutex = &sync.Mutex{}
	bot   *tgbotapi.BotAPI
)

func main() {
	// Загружаем конфигурацию
	config := LoadConfig()

	var err error
	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		// Инициализация игры для пользователя, если её ещё нет
		mutex.Lock()
		if games[chatID] == nil {
			games[chatID] = &GameState{MaxAttempts: 10}
		}
		game := games[chatID]
		mutex.Unlock()

		// Обработка команд
		if text == "/start" {
			sendMessage(chatID, "🎯 Добро пожаловать в игру 'Угадай число'!\n\nЯ загадал число от 1 до 100. Попробуй угадать его за 10 попыток!\n\nНапиши число от 1 до 100 чтобы начать.")
			newGame(game)
		} else if text == "/help" {
			sendMessage(chatID, "❓ Правила игры:\n- Я загадываю число от 1 до 100\n- У тебя есть 10 попыток, чтобы угадать его\n- После каждой попытки я скажу, больше или меньше загаданное число\n- Для начала игры напиши /start")
		} else {
			// Попытка угадать число
			if guess, err := strconv.Atoi(text); err == nil {
				handleGuess(chatID, game, guess)
			} else if !game.IsPlaying {
				sendMessage(chatID, "Напиши /start чтобы начать новую игру!")
			} else {
				sendMessage(chatID, "Пожалуйста, введи число от 1 до 100")
			}
		}
	}
}

func newGame(game *GameState) {
	rand.Seed(time.Now().UnixNano())
	game.SecretNumber = rand.Intn(100) + 1
	game.Attempts = 0
	game.IsPlaying = true
}

func handleGuess(chatID int64, game *GameState, guess int) {
	if !game.IsPlaying {
		sendMessage(chatID, "Напиши /start чтобы начать новую игру!")
		return
	}

	if guess < 1 || guess > 100 {
		sendMessage(chatID, "Пожалуйста, введи число от 1 до 100")
		return
	}

	game.Attempts++

	if guess < game.SecretNumber {
		sendMessage(chatID, fmt.Sprintf("⬆️ Загаданное число больше чем %d\n\nПопытка %d/%d", guess, game.Attempts, game.MaxAttempts))
	} else if guess > game.SecretNumber {
		sendMessage(chatID, fmt.Sprintf("⬇️ Загаданное число меньше чем %d\n\nПопытка %d/%d", guess, game.Attempts, game.MaxAttempts))
	} else {
		sendMessage(chatID, fmt.Sprintf("🎉 Поздравляю! Ты угадал число %d за %d попыток!\n\nНапиши /start чтобы сыграть ещё раз.", game.SecretNumber, game.Attempts))
		game.IsPlaying = false
		return
	}

	if game.Attempts >= game.MaxAttempts {
		sendMessage(chatID, fmt.Sprintf("😞 К сожалению, ты исчерпал все попытки. Загаданное число было: %d\n\nНапиши /start чтобы сыграть ещё раз.", game.SecretNumber))
		game.IsPlaying = false
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
