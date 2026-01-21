package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==================== КАСТОМНЫЕ ОШИБКИ НА ОСНОВЕ ИНТЕРФЕЙСА error ====================
type FinancialError struct {
	Operation string
	Amount    float64
	Reason    string
}

func (e FinancialError) Error() string {
	return fmt.Sprintf("💰 Финансовая ошибка: %s на %.2f руб. Причина: %s",
		e.Operation, e.Amount, e.Reason)
}

type TransportError struct {
	Route      string
	Vehicle    string
	Problem    string
	WastedTime time.Duration
}

func (e TransportError) Error() string {
	return fmt.Sprintf("🚌 Транспортная ошибка: %s (%s). Проблема: %s. Потеряно: %v",
		e.Route, e.Vehicle, e.Problem, e.WastedTime)
}

type MotivationError struct {
	Distraction string
	HoursLost   int
}

func (e MotivationError) Error() string {
	return fmt.Sprintf("🎯 Ошибка мотивации: отвлекся на %s, потерял %d часов",
		e.Distraction, e.HoursLost)
}

// ==================== МОТИВАЦИОННЫЕ ФРАЗЫ ====================
var motivationPhrases = []string{
	"💪 Холодный душ утром — горячий код вечером!",
	"🚀 Автобус 1346 везет в Химки, Go везет к 200К!",
	"🎯 700 рублей сегодня — 200К завтра как Go-разработчик!",
	"🔥 Никаких видеоигр! Только интерфейс error и коммиты!",
	"💡 Финансовый удар — мотивация учить Go ударными темпами!",
	"🌟 Мега-Химки — временно, Go-экосистема — навсегда!",
	"📈 2К маме сегодня, 200К себе завтра!",
	"🎮 Фильмы и сериалы подождут, горят дедлайны по Go!",
	"🏃‍♂️ Беготня по Москве — тренировка для мозга перед кодом!",
	"🚀 Тройка в автобусе, Go в карьере — билет к успеху!",
}

// Генератор мотивации
func getMotivation() string {
	return motivationPhrases[rand.Intn(len(motivationPhrases))]
}

// ==================== СИМУЛЯЦИЯ ДНЯ ====================
// Симуляция утра
func morningRoutine() error {
	fmt.Println("⏰ 10:30 - Подъем (горячей воды нет)")
	fmt.Println("🧊 Умываюсь холодной водой... бррр")

	if rand.Intn(100) < 30 {
		return fmt.Errorf("❌ Утренняя ошибка: сломался будильник, проспал")
	}

	fmt.Println("☕ Завтрак, чай")
	return nil
}

// Симуляция заказа
func executeOrder(orderName string, route string, basePrice float64) (float64, error) {
	fmt.Printf("\n📦 Заказ: %s\n", orderName)
	fmt.Printf("   Маршрут: %s\n", route)

	// 20% шанс отмены заказа
	if rand.Intn(100) < 20 {
		return 0, errors.New("заказ отменен клиентом")
	}

	// Динамическое ценообразование
	priceMultiplier := 0.5 + rand.Float64() // 0.5-1.5
	finalPrice := basePrice * priceMultiplier

	// 15% шанс транспортной проблемы
	if rand.Intn(100) < 15 {
		wastedTime := time.Duration(15+rand.Intn(45)) * time.Minute
		return 0, TransportError{
			Route:      route,
			Vehicle:    "автобус 1346",
			Problem:    "пробки на Ленинградке",
			WastedTime: wastedTime,
		}
	}

	return finalPrice, nil
}

// Симуляция финансовых операций
func handleFinances(earned float64, expenses float64) error {
	fmt.Printf("\n💰 Финансы дня:\n")
	fmt.Printf("   Заработано: %.2f руб\n", earned)
	fmt.Printf("   Расходы: %.2f руб\n", expenses)

	if earned < expenses {
		return FinancialError{
			Operation: "перевод маме",
			Amount:    2000.0,
			Reason:    "заработал меньше, чем отдал",
		}
	}

	balance := earned - expenses
	fmt.Printf("   Баланс: +%.2f руб\n", balance)

	if balance < 1000 {
		return fmt.Errorf("⚠️ Маленький баланс: всего %.2f руб", balance)
	}

	return nil
}

// Проверка мотивации
func checkMotivation(hoursProgramming int) error {
	distractions := []string{"YouTube", "Игры", "Соцсети", "Монтаж видео", "Серии"}

	if hoursProgramming < 2 {
		return MotivationError{
			Distraction: distractions[rand.Intn(len(distractions))],
			HoursLost:   2 - hoursProgramming,
		}
	}

	return nil
}

// Генератор целей
func generateGoals() []string {
	return []string{
		"📚 Выучить все интерфейсы Go standard library",
		"🏢 Устроиться Junior Go Developer в Tinkoff/Sber",
		"💼 Зарабатывать 200К к концу года",
		"🚫 Не пропускать ни одного дня 100DaysGo",
		"🎯 Сделать 100 коммитов на GitHub",
		"🏝️ Смонтировать Филиппины после первой зарплаты",
		"🔥 Пройти 3 собеседования в месяц",
		"🚀 Внести contribution в open-source проект",
	}
}

// ==================== ОБРАБОТЧИКИ ОШИБОК ====================
func handleError(err error) {
	fmt.Println("\n🔄 Обрабатываю ошибку...")

	switch e := err.(type) {
	case FinancialError:
		fmt.Printf("   🏦 %v\n", e)
		fmt.Printf("   💡 Решение: Учить Go интенсивнее!\n")
		fmt.Printf("   🎯 Мотивация: %s\n", getMotivation())

	case TransportError:
		fmt.Printf("   🚍 %v\n", e)
		fmt.Printf("   💡 В автобусе можно читать документацию Go!\n")
		fmt.Printf("   🎯 Мотивация: %s\n", getMotivation())

	case MotivationError:
		fmt.Printf("   🎮 %v\n", e)
		fmt.Printf("   💡 Блокирую %s на 24 часа\n", e.Distraction)
		fmt.Printf("   🎯 Мотивация: %s\n", getMotivation())

	default:
		fmt.Printf("   ⚠️ Стандартная ошибка: %v\n", err)
		fmt.Printf("   🎯 Мотивация: %s\n", getMotivation())
	}
}

// ==================== MAIN ====================
func main() {
	rand.Seed(time.Now().UnixNano())

	separator := strings.Repeat("=", 70)

	fmt.Println(separator)
	fmt.Println("🌅 ДЕНЬ 78: ИНТЕРФЕЙС ERROR - ФИНАНСОВАЯ МОТИВАЦИЯ")
	fmt.Println(separator)

	// Утро
	fmt.Println("\n📅 УТРЕННИЙ РИТУАЛ:")
	if err := morningRoutine(); err != nil {
		handleError(err)
	}

	// Рабочий день
	fmt.Println("\n🚀 РАБОЧИЙ ДЕНЬ КУРЬЕРА:")

	// Заказ 1: Химки → Москва
	price1, err1 := executeOrder(
		"Документы в центр Москвы",
		"Совхозная → Центр (автобус 1346)",
		700.0,
	)

	if err1 != nil {
		handleError(err1)
		price1 = 0
	} else {
		fmt.Printf("   ✅ Выполнен за %.2f руб\n", price1)
	}

	// Заказ 2: Баррикадная → Химки
	price2, err2 := executeOrder(
		"Документы в Новые Химки",
		"Баррикадная → МЕГА-Химки (автобус 359)",
		600.0,
	)

	if err2 != nil {
		handleError(err2)
		price2 = 0
	} else {
		fmt.Printf("   ✅ Выполнен за %.2f руб\n", price2)
	}

	// Финансы
	earned := price1 + price2
	expenses := 2000.0 // перевод маме

	fmt.Println("\n💸 ФИНАНСОВЫЙ ОТЧЕТ:")
	if err := handleFinances(earned, expenses); err != nil {
		handleError(err)
	}

	// Вечер: программирование
	fmt.Println("\n🌙 ВЕЧЕР ПРОГРАММИСТА:")
	fmt.Println("   🚿 Горячий душ после холодного дня")
	fmt.Println("   🍽️ Ужин, чай с мамой")
	fmt.Println("   💻 21:00 - Сажусь за Go")

	// Симуляция программирования
	programmingHours := 2 + rand.Intn(3) // 2-4 часа

	if err := checkMotivation(programmingHours); err != nil {
		handleError(err)
	} else {
		fmt.Printf("   ✅ Программировал %d часа! Интерфейс error изучен!\n", programmingHours)
	}

	// Генерация целей
	fmt.Println("\n🎯 ЦЕЛИ НА БУДУЩЕЕ:")
	goals := generateGoals()
	for i, goal := range goals {
		fmt.Printf("   %d. %s\n", i+1, goal)
	}

	// Геймификация
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ДНЯ 78:")
	score := 0
	if earned > 0 {
		score += 10
	}
	if programmingHours >= 2 {
		score += 20
	}
	if err1 == nil {
		score += 5
	}
	if err2 == nil {
		score += 5
	}

	fmt.Printf("   🏆 Очки продуктивности: %d/40\n", score)
	fmt.Printf("   🎯 Уровень мотивации: %d%%\n", (score*100)/40)
	fmt.Printf("   💰 До цели 200К: %.1f%%\n", (earned/200000)*100)

	// Disclaimer
	fmt.Println("\n" + separator)
	fmt.Println("📢 DISCLAIMER:")
	fmt.Println("   Все персонажи и события вымышлены.")
	fmt.Println("   Любые совпадения с реальными людьми случайны.")
	fmt.Println("   История создана для мотивации изучения Go.")
	fmt.Println("   © Daily Code Life Story - художественный вымысел.")
	fmt.Println(separator)

	// Финальная мотивация
	fmt.Println("\n🚀 ЗАВТРА ДЕНЬ 79! ПРОДОЛЖАЕМ ИЗУЧАТЬ ERROR WRAPPING!")
	fmt.Println("   Помни: каждый interface{} приближает к работе мечты!")
}
