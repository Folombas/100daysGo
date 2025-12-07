package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

// Структура клиента с инкапсулированными полями
type Client struct {
	Name       string
	Attitude   string
	Comment    string
	HatredLevel int
}

// Структура посылки с методами
type Package struct {
	WeightKG float64
	Content  string
	Fragile  bool
	Delivered bool
}

// Метод для получения статуса посылки
func (p Package) GetStatus() string {
	if p.Delivered {
		return "✅ Доставлено"
	}
	return "📦 В пути"
}

// Структура для отслеживания прогресса Гоши
type GoshaProgress struct {
	Name            string
	Age             int
	Confidence      int
	Willpower       int
	CodeLines       int
	DaysLearning    int
	Savings         float64
	StructsLearned  int
	HatredAbsorbed  int
}

// Метод для отображения прогресса
func (gp GoshaProgress) DisplayProgress() {
	fmt.Printf("🔥 %s - СТРОИТЕЛЬ СТРУКТУР 🔥\n", gp.Name)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("🎂 Возраст: %d лет\n", gp.Age)
	fmt.Printf("💪 Уровень уверенности: %d%%\n", gp.Confidence)
	fmt.Printf("🛡️ Уровень воли: %d%%\n", gp.Willpower)
	fmt.Printf("📅 Дней обучения: %d\n", gp.DaysLearning)
	fmt.Printf("💻 Написано строк кода: %d\n", gp.CodeLines)
	fmt.Printf("💰 Накоплено: %.0f ₽\n", gp.Savings)
	fmt.Printf("🏗️ Изучено структур: %d\n", gp.StructsLearned)
	fmt.Printf("😠 Поглощено ненависти: %d единиц\n", gp.HatredAbsorbed)
	fmt.Println(strings.Repeat("=", 60))
}

// Структура для темы дня
type DailyTheme struct {
	Date  string
	Topic string
	Day   int
}

// Структура для жизненного урока
type LifeLesson struct {
	Title       string
	Description string
	CodeAnalogy string
}

func main() {
	// Генерация случайных данных
	r := rand.New(rand.NewPCG(
		uint64(time.Now().UnixNano()),
		uint64(time.Now().UnixNano()>>32),
	))

	// Инициализация прогресса Гоши
	gosha := GoshaProgress{
		Name:           "Гоша",
		Age:            38,
		Confidence:     48,
		Willpower:      60,
		CodeLines:      1280,
		DaysLearning:   34,
		Savings:        1450.75,
		StructsLearned: 15,
		HatredAbsorbed: 42,
	}

	// Создание списка клиентов
	clients := []Client{
		{"Мадам в норковой шубке", "Грубая", "Петрович, смотри, этот неудачник даже по воскресеньям с коробками таскается, ахах", 9},
		{"Лысый Петрович", "Пренебрежительный", "Да у меня дворник больше зарабатывает!", 7},
		{"Элитный клиент", "Снисходительный", "Вы бы лучше кодили вместо этого...", 5},
	}

	// Создание посылок
	packages := []Package{
		{WeightKG: 16.5, Content: "Букет с хрупкими цветами", Fragile: true, Delivered: true},
		{WeightKG: 22.0, Content: "Техника в коробке 'до 10 кг'", Fragile: false, Delivered: false},
		{WeightKG: 8.3, Content: "Подарок на день рождения", Fragile: true, Delivered: true},
	}

	// Тема дня
	theme := DailyTheme{
		Date:  "7 декабря 2025",
		Topic: "Structs",
		Day:   34,
	}

	// Жизненный урок
	lesson := LifeLesson{
		Title:       "Структуры vs Хаос",
		Description: "Когда мир пытается сломать тебя грубостью и насмешками — создай прочную структуру своей жизни.",
		CodeAnalogy: "Как структуры организуют данные в Go, так и ты должен организовать свою жизнь в строгую архитектуру успеха.",
	}

	// Отображение прогресса
	gosha.DisplayProgress()

	// Отображение темы дня
	fmt.Printf("📅 Тема дня: %s\n", theme.Topic)
	fmt.Printf("📆 Дата: %s\n", theme.Date)
	fmt.Printf("🔢 День челленджа: %d\n", theme.Day)
	fmt.Println(strings.Repeat("-", 40))

	// Отображение клиентов
	fmt.Println("😤 КЛИЕНТЫ СЕГОДНЯ (ВОСКРЕСЕНЬЕ):")
	for i, client := range clients {
		fmt.Printf("%d. %s: \"%s\"\n", i+1, client.Name, client.Comment)
		fmt.Printf("   💢 Уровень грубости: %d/10\n", client.HatredLevel)
	}
	fmt.Println(strings.Repeat("-", 40))

	// Отображение посылок
	fmt.Println("📦 ПОСЫЛКИ ДЛЯ ДОСТАВКИ:")
	for i, pkg := range packages {
		status := pkg.GetStatus()
		fragileStatus := ""
		if pkg.Fragile {
			fragileStatus = "⚠️ Хрупкое!"
		}
		fmt.Printf("%d. [%s] %.1f кг | %s %s\n",
			i+1, status, pkg.WeightKG, pkg.Content, fragileStatus)
	}
	fmt.Println(strings.Repeat("=", 60))

	// Отображение урока
	fmt.Println("💡 ЖИЗНЕННЫЙ УРОК ДНЯ:")
	fmt.Printf("🌟 %s\n", lesson.Title)
	fmt.Println(lesson.Description)
	fmt.Printf("⚙️ Аналогия в коде: %s\n", lesson.CodeAnalogy)
	fmt.Println(strings.Repeat("-", 40))

	// Мотивационное сообщение
	motivations := []string{
		"Твоя жизнь — это структура. Определи поля правильно!",
		"Каждая насмешка — это просто шум в канале. Фильтруй его!",
		"Структуры не разрушаются под давлением — ни в коде, ни в жизни!",
		"Твой код сегодня — это твой будущий офис с видом на город!",
		"Из каждой грубости собирай данные для своей мотивации!",
		"Структурируй свою ненависть в энергию для обучения!",
		"Ты не курьер. Ты архитектор своей судьбы!",
	}
	motivation := motivations[r.IntN(len(motivations))]

	fmt.Println("⚡ МОТИВАЦИОННЫЙ ПУЛЬС:")
	fmt.Printf("💥 \"%s\"\n", motivation)
	fmt.Println(strings.Repeat("=", 60))

	// Завершение
	fmt.Println("🚀 СЛЕДУЮЩАЯ ЦЕЛЬ: 75 ДНЕЙ И КОРПОРАТИВНЫЙ ОФИС В ЦЕНТРЕ МОСКВЫ!")
	fmt.Println("💻 КОД СЕГОДНЯ: СТРУКТУРЫ БУДУЩЕГО. ЖИЗНЬ ЗАВТРА: СТРУКТУРА УСПЕХА!")
}
