package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
	"unicode"
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
	HeavyPackages     int
	TodaysWeight      int
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
		Confidence:        40.0,
		Willpower:         52.0,
		DaysSinceChange:   32,
		StudyHours:        5.5,
		MoneySaved:        1300.0,
		AbandonedGames:    15,
		AbandonedSeries:   7,
		HeavyPackages:     3,
		TodaysWeight:      15,
	}
}

// Функция для получения текущей темы дня
func getCurrentTheme() DailyTheme {
	return DailyTheme{
		Date:   "5 декабря 2025",
		Topic:  "Strings",
		Day:    32,
	}
}

// Функция для генерации мотивационного сообщения
func getMotivationalMessage(r *rand.Rand) string {
	messages := []string{
		"Каждая строка кода — это шаг к свободе от коробок!",
		"Твои строки кода сильнее любых тяжелых посылок!",
		"Когда мир пытается сломать тебя — конкатенируй свою силу!",
		"Твой путь — от строк 'курьер' к строкам 'разработчик'!",
		"Строка кода сегодня — стабильная зарплата завтра!",
		"Твои руки созданы не для коробок, а для клавиатуры!",
		"Каждый символ в коде — это гвоздь в гроб твоей старой жизни!",
		"Твои строки кода греют душу лучше любого торгового центра!",
		"Когда тебя обманывают с весом — обманывай системы своими строками!",
		"Ты не таскаешь коробки — ты собираешь строки в программу своей жизни!",
	}

	return messages[r.IntN(len(messages))]
}

// Функция для преобразования текста в заголовок (Title Case)
func toTitleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runeWord := []rune(word)
		if len(runeWord) > 0 {
			runeWord[0] = unicode.ToUpper(runeWord[0])
		}
		words[i] = string(runeWord)
	}
	return strings.Join(words, " ")
}

// Функция для отображения статуса Гоши
func (g Gosha) displayStatus() {
	fmt.Println("🔥🔥🔥 ГОША - ПРОГРАММИСТ-ПЕРЕРОЖДЕННЫЙ 🔥🔥🔥")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("👤 Имя: %s\n", g.Name)
	fmt.Printf("🎂 Возраст: %d лет\n", g.Age)
	fmt.Printf("🎤 Бывшая карьера: %s\n", g.PreviousCareer)
	fmt.Printf("📦 Текущая работа: %s\n", g.CurrentJob)
	fmt.Printf("🎯 Цель: %s\n", g.ProgrammingGoal)
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("💪 Уверенность: %.0f%%\n", g.Confidence)
	fmt.Printf("🛡️ Воля: %.0f%%\n", g.Willpower)
	fmt.Printf("📅 Дней с изменений: %d\n", g.DaysSinceChange)
	fmt.Printf("💻 Часов обучения: %.1f\n", g.StudyHours)
	fmt.Printf("💰 Сэкономлено: %.0f ₽\n", g.MoneySaved)
	fmt.Printf("🎮 Удалено игр: %d\n", g.AbandonedGames)
	fmt.Printf("📺 Удалено сериалов: %d\n", g.AbandonedSeries)
	fmt.Printf("📦 Тяжелых посылок сегодня: %d (реальный вес: %d кг!)\n", g.HeavyPackages, g.TodaysWeight)
	fmt.Println(strings.Repeat("=", 60))
}

// Функция для отображения темы дня
func (t DailyTheme) displayTheme() {
	fmt.Printf("📅 Тема дня: %s\n", t.Topic)
	fmt.Printf("📆 Дата: %s\n", t.Date)
	fmt.Printf("🔢 День челленджа: %d\n", t.Day)
	fmt.Println(strings.Repeat("-", 40))
}

// Функция для отображения мотивации
func displayMotivation(r *rand.Rand) {
	message := getMotivationalMessage(r)
	decorated := fmt.Sprintf("✨ %s ✨", message)

	fmt.Println("💬 МОТИВАЦИЯ ДНЯ:")
	fmt.Println(decorated)
	fmt.Println(strings.Repeat("-", 40))
}

// Функция для отображения истории
func displayHistory() {
	fmt.Println("📖 ИСТОРИЯ ГОШИ СЕГОДНЯ:")
	fmt.Println("Сегодня я взял заказ, в котором написано 'до 10 кг'.")
	fmt.Println("А там на самом деле по ощущениям было 15 кг, не меньше.")
	fmt.Println("Я чуть не умер, пока тащил эту тяжеленную коробку по улицам и в метро.")
	fmt.Println("Тётка сказала: 'Там в коробке ничего хрупкого нету, не бойтесь!'")
	fmt.Println("Как смешно! А как вашу коробку тащить?")
	fmt.Println("Хорошо, что я такой здоровый мощный бык и смог её кое-как дотащить!")
	fmt.Println("Но я понял одно — надо далее учить язык программирования Go")
	fmt.Println("и заканчивать с этой ишачьей курьерской работой!")
	fmt.Println(strings.Repeat("-", 40))
}

// Функция для отображения урока по строкам
func displayStringLesson() {
	fmt.Println("🔧 STRINGS В GO - ОСНОВНЫЕ ОПЕРАЦИИ:")
	fmt.Println("- Строки в Go — это неизменяемые последовательности байтов")
	fmt.Println("- Конкатенация: s := 'Hello' + ' World'")
	fmt.Println("- Поиск подстроки: strings.Contains(s, 'World')")
	fmt.Println("- Замена: strings.Replace(s, 'old', 'new', -1)")
	fmt.Println("- Разделение: strings.Split(s, ' ')")
	fmt.Println("- Преобразование регистра: strings.ToUpper(s)")
	fmt.Println("- Работа с Unicode: unicode.ToUpper(rune)")
	fmt.Println(strings.Repeat("=", 60))

	// Демонстрация работы со строками
	original := "курьер"
	transformed := strings.ReplaceAll(strings.ToUpper(original), "КУРЬЕР", "GO-РАЗРАБОТЧИК")

	fmt.Printf("🎯 Пример трансформации:\n")
	fmt.Printf("  Было: \"%s\"\n", original)
	fmt.Printf("  Стало: \"%s\"\n", transformed)
	fmt.Printf("  Как твоя жизнь: из %s в %s!\n", original, transformed)
	fmt.Println(strings.Repeat("-", 40))
}

// Функция для отображения жизненного урока
func displayLifeLesson() {
	lesson := "Когда тебя обманывают с весом посылки — помни: "
	lesson += "твоя ценность не в том, сколько ты можешь унести на плечах, "
	lesson += "а в том, какие строки кода ты можешь написать."

	formatted := toTitleCase(lesson)

	fmt.Println("💭 ЖИЗНЕННЫЙ УРОК ДНЯ:")
	fmt.Println(formatted)
	fmt.Println(strings.Repeat("-", 40))
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
	displayStringLesson()
	displayLifeLesson()

	// Завершение программы
	fmt.Println("🚀 ПРОДОЛЖАЙ СВОИ ЗАНЯТИЯ, ГОША! ТВОИ СТРОКИ КОДА СИЛЬНЕЕ ЛЮБЫХ КОРОБОК!")
	fmt.Println("💻 ЗАВТРА ТЫ БУДЕШЬ НЕ ТАСКАТЬ ПОСЫЛКИ, А ПИСАТЬ ПРОГРАММЫ, КОТОРЫЕ ИХ РАЗВОЗЯТ!")
}

