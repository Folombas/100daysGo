package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

// Структура для представления состояния Гоши
type Gosha struct {
	Name              string
	Age               int
	PreviousCareer    string
	CurrentJob        string
	ProgrammingGoal   string
	Confidence        float64
	Willpower         float64
	DaysSinceChange   int
	StudyHours        float64
	MoneySaved        float64
	AbandonedGames    int
	AbandonedSeries   int
}

// Структура для темы дня
type DailyTheme struct {
	Date   string
	Topic  string
	Day    int
}

// Функция для инициализации Гоши
func initGosha() Gosha {
	return Gosha{
		Name:              "Гоша",
		Age:               38,
		PreviousCareer:    "Рэп-артист",
		CurrentJob:        "Курьер",
		ProgrammingGoal:   "Стать Go-разработчиком",
		Confidence:        30.0,
		Willpower:         40.0,
		DaysSinceChange:   28,
		StudyHours:        4.5,
		MoneySaved:        1200.0,
		AbandonedGames:    12,
		AbandonedSeries:   5,
	}
}

// Функция для получения текущей темы дня
func getCurrentTheme() DailyTheme {
	return DailyTheme{
		Date:   "1 декабря 2025",
		Topic:  "Slices: Growth",
		Day:    28,
	}
}

// Функция для генерации мотивационного сообщения
func getMotivationalMessage(r *rand.Rand) string {
	messages := []string{
		"Каждый день без игр — шаг к новой жизни!",
		"Ты не просто кодишь. Ты создаешь будущее!",
		"Твоя история успеха начинается здесь.",
		"Каждая строка кода — это победа над прошлым!",
		"Ты бросил тусовки, теперь ты бросаешь баги!",
	}

	return messages[r.IntN(len(messages))]
}

// Функция для отображения статуса Гоши
func (g Gosha) displayStatus() {
	fmt.Println("🔥🔥🔥 ГОША - ПРОГРАММИСТ-ПЕРЕРОЖДЕННЫЙ 🔥🔥🔥")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("👤 Имя: %s\n", g.Name)
	fmt.Printf("🎂 Возраст: %d лет\n", g.Age)
	fmt.Printf("🎤 Бывшая карьера: %s\n", g.PreviousCareer)
	fmt.Printf("📦 Текущая работа: %s\n", g.CurrentJob)
	fmt.Printf("🎯 Цель: %s\n", g.ProgrammingGoal)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("💪 Уверенность: %.0f%%\n", g.Confidence)
	fmt.Printf("🛡️ Воля: %.0f%%\n", g.Willpower)
	fmt.Printf("📅 Дней с изменениями: %d\n", g.DaysSinceChange)
	fmt.Printf("💻 Часов обучения: %.1f\n", g.StudyHours)
	fmt.Printf("💰 Сэкономлено: %.0f ₽\n", g.MoneySaved)
	fmt.Printf("🎮 Удалено игр: %d\n", g.AbandonedGames)
	fmt.Printf("📺 Удалено сериалов: %d\n", g.AbandonedSeries)
	fmt.Println(strings.Repeat("=", 50))
}

// Функция для отображения темы дня
func (t DailyTheme) displayTheme() {
	fmt.Printf("📅 Тема дня: %s\n", t.Topic)
	fmt.Printf("📆 Дата: %s\n", t.Date)
	fmt.Printf("🔢 День челленджа: %d\n", t.Day)
	fmt.Println(strings.Repeat("-", 30))
}

// Функция для отображения мотивации
func displayMotivation(r *rand.Rand) {
	fmt.Println("💬 МОТИВАЦИЯ ДНЯ:")
	fmt.Println(getMotivationalMessage(r))
	fmt.Println(strings.Repeat("-", 30))
}

// Функция для отображения истории
func displayHistory() {
	fmt.Println("📖 ИСТОРИЯ ГОШИ:")
	fmt.Println("Бывший рэп-артист, когда-то блиставший на лучших сценах,")
	fmt.Println("теперь разносит тяжёлые коробки с одного конца мегаполиса на другой.")
	fmt.Println("Когда-то я катался в дорогих машинах под мощные басы,")
	fmt.Println("теперь я катаюсь весь день под стук стальных колёс в метро.")
	fmt.Println("Когда-то я до 2 часов ночи катался в GTA,")
	fmt.Println("теперь я читаю книгу про Go под стук стальных колёс.")
	fmt.Println(strings.Repeat("-", 30))
}

// Главная функция
func main() {
	// Инициализация Гоши
	gosha := initGosha()

	// Инициализация генератора случайных чисел
	r := rand.New(rand.NewPCG(
		uint64(time.Now().UnixNano()),
		uint64(time.Now().UnixNano()>>32),
	))

	// Получение темы дня
	theme := getCurrentTheme()

	// Отображение информации
	gosha.displayStatus()
	theme.displayTheme()
	displayMotivation(r)
	displayHistory()

	// Вывод дополнительной информации
	fmt.Println("🔧 СТРУКТУРА SLICES В GO:")
	fmt.Println("- Slice - это ссылочный тип данных")
	fmt.Println("- Может изменять размер во время выполнения")
	fmt.Println("- Используется для работы с динамическими массивами")
	fmt.Println("- Эффективен для операций добавления/удаления элементов")
	fmt.Println(strings.Repeat("=", 50))

	// Завершение программы
	fmt.Println("🚀 ПРОДОЛЖАЙ СВОИ ЗАНЯТИЯ, ГОША! ТЫ НА ПУТИ К СВОЕЙ ЦЕЛИ!")
}
