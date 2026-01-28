// Day85: Снежная метель и горутины
// Отменённых заказик Гоши и зимняя хандра

package main

import (
	"fmt"
	"sync"
	"time"
)

// Trip представляет этап путешествия
type Trip struct {
	from     string
	to       string
	cost     int
	status   string
	duration time.Duration
}

// Функция-горутина для поездки
func rideTrain(trip Trip, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	fmt.Printf("🚂 Горутина стартовала: %s → %s\n", trip.from, trip.to)

	// Имитация времени поездки
	time.Sleep(trip.duration)

	if trip.status == "completed" {
		results <- fmt.Sprintf("✅ Достиг цели: %s → %s (стоимость: %d руб.)",
			trip.from, trip.to, trip.cost)
	} else {
		results <- fmt.Sprintf("❌ Отмена: %s → %s (возврат: %d руб.)",
			trip.from, trip.to, trip.cost)
	}
}

// Мониторинг прогресса (ещё одна горутина)
func progressMonitor(done chan bool) {
	fmt.Println("📊 Запущен мониторинг прогресса обучения...")

	for i := 1; i <= 10; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("   Прогресс: %d0%% | Дофамин: ↑\n", i)
	}

	done <- true
}

func main() {
	fmt.Println("=== День 85: ГОРУТИНЫ В МЕТЕЛЬ ===")
	fmt.Println("История Гоши в параллельных реальностях")

	// Поездки Гоши (конкурентные задачи)
	trips := []Trip{
		{"Химки", "Авиамоторная", 950, "completed", 2 * time.Second},
		{"Перово", "Сходня", 650, "cancelled", 1 * time.Second},
		{"Ховрино", "Дом", 0, "completed", 3 * time.Second},
	}

	var wg sync.WaitGroup
	results := make(chan string, len(trips))
	done := make(chan bool)

	// Запуск горутины-монитора
	go progressMonitor(done)

	// Запуск горутин-поездок
	fmt.Println("\n⚡ Запуск горутин-поездок...")
	for _, trip := range trips {
		wg.Add(1)
		go rideTrain(trip, &wg, results)
	}

	// Ожидание завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Сбор результатов
	fmt.Println("\n📨 Результаты поездок:")
	for result := range results {
		fmt.Println("  ", result)
	}

	// Ожидание монитора
	<-done

	fmt.Println("\n🎯 ИТОГ ДНЯ:")
	fmt.Println("Пройдено 3 сегмента пути")
	fmt.Println("Заработано: 0 руб. (заказ отменили)")
	fmt.Println("Получено опыта: +100 XP")
	fmt.Println("Выработано дофамина: ⬆⬆⬆")
	fmt.Println("\n💡 МОРАЛЬ: Горутины выполняются, даже если одна отменяется!")
	fmt.Println("   Продолжай учиться, даже если что-то не получается!")
}
