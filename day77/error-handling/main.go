package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ==================== ТИПЫ ОШИБОК ====================
type DeliveryError struct {
	Action  string
	Message string
	Code    int
}

func (e DeliveryError) Error() string {
	return fmt.Sprintf("🚫 Ошибка доставки [%d]: %s - %s", e.Code, e.Action, e.Message)
}

type BalanceError struct {
	ProgrammingTime time.Duration
	EditingTime     time.Duration
}

func (e BalanceError) Error() string {
	return fmt.Sprintf("⚖️ Дисбаланс: программирование=%v, монтаж=%v",
		e.ProgrammingTime, e.EditingTime)
}

// ==================== МОТИВАЦИОННЫЕ ФРАЗЫ ====================
var motivationalPhrases = []string{
	"💡 Каждая обработанная ошибка делает тебя сильнее в Go!",
	"🚀 Ошибки — это ступеньки к мастерству, а не стены!",
	"🎯 Сегодня ты обработал ошибку, завтра напишешь микросервис!",
	"🔥 Не отвлекайся на клубы/бары — твой код ждет тебя!",
	"💪 Видеомонтаж подождет, сначала стань гуру Go!",
	"🌟 Курьерка — временно, программирование — навсегда!",
	"📈 Каждая строчка кода приближает к работе мечты!",
	"🎮 Баланс важен, но сегодня побеждает Go!",
	"🏝️ После первого года работы Go-разработчиком Филиппины ждут с новым тревел-влогом!",
	"🚀 Go вперед — эко-система Go ждет своего айти-героя!",
}

func getMotivation() string {
	index := int(time.Now().Unix()) % len(motivationalPhrases)
	return motivationalPhrases[index]
}

// ==================== СИМУЛЯЦИЯ ДНЯ ГОШИ ====================
func acceptOrder(orderID string) (string, error) {
	// Симуляция: 30% заказов отменяются
	if time.Now().Nanosecond()%10 < 3 {
		return "", DeliveryError{
			Action:  "Принятие заказа " + orderID,
			Message: "Заказ отменен клиентом",
			Code:    401,
		}
	}
	return fmt.Sprintf("Заказ %s принят! 🎉", orderID), nil
}

func getCompensation(orderID string) (float64, error) {
	// Симуляция: компенсация только в 10% случаев
	if time.Now().Nanosecond()%10 == 0 {
		return 50.0, nil
	}
	return 0.0, errors.New("компенсация не предоставлена")
}

func installCapCut(isTired bool) error {
	if isTired {
		return errors.New("устал сегодня, установлю завтра")
	}
	fmt.Println("🎬 CapCut установлен! Готов к монтажу отпуска на Филиппинах лето-2019 🏝️🌴!")
	return nil
}

func dailyBalance(programmingHours, editingHours int) error {
	if programmingHours == 0 {
		return BalanceError{
			ProgrammingTime: 0,
			EditingTime:     time.Duration(editingHours) * time.Hour,
		}
	}

	if editingHours > programmingHours*2 {
		return fmt.Errorf("⚠️ Слишком много монтажа! Go в опасности!")
	}

	fmt.Printf("✅ Отличный баланс! Код: %dч, Монтаж: %dч\n",
		programmingHours, editingHours)
	return nil
}

// ==================== ОБРАБОТЧИКИ ОШИБОК ====================
func handleDeliveryError(err error) {
	if de, ok := err.(DeliveryError); ok {
		fmt.Printf("📦 Обрабатываем ошибку доставки: %v\n", de.Message)
		fmt.Println("   Извлекаем урок: нужно учить Go!")
	} else {
		fmt.Printf("❌ Неизвестная ошибка: %v\n", err)
	}
}

func tryRecover() {
	if r := recover(); r != nil {
		fmt.Println("🚑 Критическая ошибка! Восстанавливаемся...")
		fmt.Println("   Мотивация:", getMotivation())
	}
}

// ==================== MAIN ====================
func main() {
	defer tryRecover()

	separator := strings.Repeat("=", 50)

	fmt.Println(separator)
	fmt.Println("🌅 ДЕНЬ 77: ОСНОВЫ ОБРАБОТКИ ОШИБОК В GO")
	fmt.Println(separator)

	// 1. Симуляция утра
	fmt.Println("\n⏰ 8:00 - Подъем... эх, поспать бы еще")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("😴 9:30 - Фактический подъем")

	// 2. Принятие заказа с обработкой ошибки
	fmt.Println("\n🛵 ПРИНИМАЮ ЗАКАЗ ИЗ ДОМА...")
	result, err := acceptOrder("MSC-2024-77")
	if err != nil {
		handleDeliveryError(err)
		fmt.Println(getMotivation())
	} else {
		fmt.Println(result)
	}

	// 3. Попытка получить компенсацию
	fmt.Println("\n💰 ПЫТАЮСЬ ПОЛУЧИТЬ КОМПЕНСАЦИЮ...")
	comp, err := getCompensation("MSC-2024-77")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		fmt.Println("💡 Вывод: нужно учить Go, чтобы не зависеть от курьерки!")
		fmt.Println(getMotivation())
	} else {
		fmt.Printf("✅ Получена компенсация: %.2f руб\n", comp)
	}

	// 4. Установка CapCut
	fmt.Println("\n🎬 УСТАНОВКА CAPCUT ДЛЯ МОНТАЖА...")
	err = installCapCut(true) // устал = true
	if err != nil {
		fmt.Printf("⏸️ %v\n", err)
		fmt.Println(getMotivation())
	}

	// 5. Балансировка обучения (с panic/recover демонстрацией)
	fmt.Println("\n⚖️ БАЛАНС ОБУЧЕНИЯ НА ДЕНЬ 77...")

	// Демонстрация паники
	fmt.Println("💥 Симуляция критической ситуации...")
	go func() {
		time.Sleep(100 * time.Millisecond)
		panic("🚨 Критическая ошибка: Дисбаланс зашкаливает!")
	}()
	time.Sleep(200 * time.Millisecond)

	// Нормальная балансировка
	err = dailyBalance(3, 1) // 3ч программирования, 1ч монтажа
	if err != nil {
		fmt.Printf("⚠️ %v\n", err)
		fmt.Println("📝 План на завтра: 4ч Go, 0.5ч монтаж")
	}

	// 6. Вывод итогов дня
	fmt.Println("\n" + separator)
	fmt.Println("📊 ИТОГИ ДНЯ 77:")
	fmt.Println("   ✓ Изучены основы обработки ошибок")
	fmt.Println("   ✓ Практика: errors.New, error интерфейс")
	fmt.Println("   ✓ Практика: кастомные типы ошибок")
	fmt.Println("   ✓ Практика: panic/recover для критических ошибок")
	fmt.Println("   ✗ CapCut не установлен (отложено)")
	fmt.Println("   ✓ Баланс найден!")

	fmt.Println("\n" + separator)
	fmt.Println("📢 DISCLAIMER:")
	fmt.Println("   Все персонажи и события в историях Гоши вымышлены.")
	fmt.Println("   Любые совпадения с реальными людьми случайны.")
	fmt.Println("   Истории созданы для мотивации изучения Go.")
	fmt.Println("   © Daily Code Life IT Story - художественный вымысел.")
	fmt.Println(separator)

	fmt.Println("\n🎯 КОММИЧЬСЯ И УЧИСЬ ДАЛЬШЕ! DAY78 ЖДЕТ!")
}
