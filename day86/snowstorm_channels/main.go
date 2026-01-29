package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// Курьерский заказ
type DeliveryOrder struct {
	ID        int
	From      string
	To        string
	Distance  int    // в км
	Price     int    // в рублях
	Status    string // "ожидает", "в пути", "доставлен", "отменен"
	Danger    int    // уровень сложности (0-100)
}

// Канал для связи между курьерами и диспетчером
type CourierChannel struct {
	ID        int
	Name      string
	Channel   chan DeliveryOrder
	Active    bool
	XP        int  // опыт курьера
	Health    int  // здоровье (0-100)
	Completed int  // выполнено заказов
}

func main() {
	fmt.Println("❄️ День 86: КАНАЛЫ В СНЕЖНОЙ МЕТЕЛИ ❄️")
	fmt.Println("История Гоши о том, как каналы спасают в метель")

	rand.Seed(time.Now().UnixNano())

	// Создаем каналы для разных типов заказов
	standardOrders := make(chan DeliveryOrder, 5)    // буферизованный канал
	urgentOrders := make(chan DeliveryOrder)         // небуферизованный канал
	emergencyOrders := make(chan DeliveryOrder, 3)   // буферизованный для срочных
	weatherAlerts := make(chan string, 10)           // канал для оповещений о погоде

	// WaitGroup для отслеживания горутин
	var wg sync.WaitGroup

	// Канал для результатов доставки
	results := make(chan string, 15)

	// Запускаем метель (постоянный источник проблем)
	wg.Add(1)
	go func() {
		defer wg.Done()
		alerts := []string{
			"⚠️ МЕТЕЛЬ: Видимость упала до 50 метров!",
			"❄️ СНЕГОПАД: Дороги заметает!",
			"🌬️ ВЕТЕР: Порывы до 15 м/с!",
			"☃️ СУГРОБЫ: Высота снега 40 см!",
			"🚨 ОПАСНО: Гололед на дорогах!",
		}
		for i := 0; i < 8; i++ {
			time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)
			alert := alerts[rand.Intn(len(alerts))]
			weatherAlerts <- alert
			results <- fmt.Sprintf("🌡️  ПОГОДА: %s", alert)
		}
		close(weatherAlerts)
	}()

	// Генератор заказов
	wg.Add(1)
	go func() {
		defer wg.Done()
		orders := []DeliveryOrder{
			{1, "Склад на улице", "Новопеределкино", 25, 950, "ожидает", 30},
			{2, "Центр", "Север Москвы", 15, 750, "ожидает", 20},
			{3, "Юг", "Запад", 20, 850, "ожидает", 40},
			{4, "Аэропорт", "Сити", 30, 1200, "ожидает", 60},
			{5, "Восток", "Северо-Восток", 10, 500, "ожидает", 10},
			{6, "Патрики", "Останкино", 8, 600, "ожидает", 25},
			{7, "МГУ", "Физтех", 12, 700, "ожидает", 35},
			{8, "ИКЕА", "ТЦ", 5, 400, "ожидает", 5},
		}

		for _, order := range orders {
			time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)

			// Выбираем канал в зависимости от цены и опасности
			if order.Price > 1000 || order.Danger > 50 {
				urgentOrders <- order
				results <- fmt.Sprintf("🚨 СРОЧНЫЙ ЗАКАЗ: %s → %s (%d руб.)",
					order.From, order.To, order.Price)
			} else if order.Danger > 30 {
				emergencyOrders <- order
				results <- fmt.Sprintf("⚠️  СЛОЖНЫЙ ЗАКАЗ: %s → %s", order.From, order.To)
			} else {
				standardOrders <- order
				results <- fmt.Sprintf("📦 СТАНДАРТНЫЙ ЗАКАЗ: %s → %s", order.From, order.To)
			}
		}

		close(standardOrders)
		close(urgentOrders)
		close(emergencyOrders)
	}()

	// Курьер Гоша (обрабатывает заказы из нескольких каналов)
	wg.Add(1)
	go func() {
		defer wg.Done()
		courier := CourierChannel{
			ID:      1,
			Name:    "Гоша-Курьер",
			Channel: make(chan DeliveryOrder, 2),
			Active:  true,
			XP:      100,
			Health:  80,
		}

		results <- fmt.Sprintf("🚶‍♂️ КУРЬЕР %s ВЫХОДИТ НА МАРШРУТ (Здоровье: %d%%, Опыт: %d)",
			courier.Name, courier.Health, courier.XP)

		deliveryCount := 0

		for courier.Health > 0 && deliveryCount < 8 {
			select {
			case order, ok := <-standardOrders:
				if !ok {
					standardOrders = nil
				} else {
					time.Sleep(time.Duration(order.Distance*100) * time.Millisecond)
					if rand.Intn(100) > order.Danger {
						courier.Completed++
						courier.XP += 10
						results <- fmt.Sprintf("✅ %s доставил заказ #%d (%s → %s) +10 XP",
							courier.Name, order.ID, order.From, order.To)
					} else {
						courier.Health -= 15
						results <- fmt.Sprintf("❌ %s не справился с заказом #%d (метель!) -15%% здоровья",
							courier.Name, order.ID)
					}
					deliveryCount++
				}

			case order, ok := <-urgentOrders:
				if !ok {
					urgentOrders = nil
				} else {
					time.Sleep(time.Duration(order.Distance*50) * time.Millisecond)
					if rand.Intn(100) > order.Danger/2 {
						courier.Completed++
						courier.XP += 25
						courier.Health -= 5
						results <- fmt.Sprintf("🚀 %s ВЫПОЛНИЛ СРОЧНЫЙ заказ #%d +25 XP",
							courier.Name, order.ID)
					} else {
						courier.Health -= 25
						results <- fmt.Sprintf("💥 %s ПРОВАЛИЛ срочный заказ #%d -25%% здоровья",
							courier.Name, order.ID)
					}
					deliveryCount++
				}

			case alert := <-weatherAlerts:
				courier.Health -= 5
				results <- fmt.Sprintf("🌨️  %s пострадал от: %s -5%% здоровья",
					courier.Name, alert)

			case <-time.After(2 * time.Second):
				results <- fmt.Sprintf("⏰ %s ждет новые заказы...", courier.Name)

			case order, ok := <-emergencyOrders:
				if !ok {
					emergencyOrders = nil
				} else {
					time.Sleep(time.Duration(order.Distance*150) * time.Millisecond)
					if rand.Intn(100) > order.Danger {
						courier.Completed++
						courier.XP += 50
						courier.Health -= 10
						results <- fmt.Sprintf("🏆 %s ГЕРОИЧЕСКИ выполнил сложный заказ #%d +50 XP",
							courier.Name, order.ID)
					} else {
						courier.Health -= 30
						results <- fmt.Sprintf("💀 %s не смог пробиться через сугробы к заказу #%d -30%% здоровья",
							courier.Name, order.ID)
					}
					deliveryCount++
				}
			}

			// Проверяем здоровье
			if courier.Health <= 0 {
				results <- fmt.Sprintf("🏥 %s УСТАЛ И ВЕРНУЛСЯ ДОМОЙ. Нужно учить Go!", courier.Name)
				break
			}
		}

		// Итоги работы курьера
		results <- fmt.Sprintf("\n📊 ИТОГИ РАБОТЫ %s:", courier.Name)
		results <- fmt.Sprintf("   Доставлено заказов: %d", courier.Completed)
		results <- fmt.Sprintf("   Заработано опыта: %d XP", courier.XP)
		results <- fmt.Sprintf("   Остаток здоровья: %d%%", courier.Health)
		if courier.XP >= 200 {
			results <- fmt.Sprintf("   🎉 ОТЛИЧНО! Можно купить горячий чай и учить Go!")
		} else {
			results <- fmt.Sprintf("   💪 Нужно больше стараться! Время учить каналы в Go!")
		}
	}()

	// Выводим результаты в реальном времени
	go func() {
		wg.Wait()
		close(results)
	}()

	// Читаем и выводим результаты
	fmt.Println("📡 НАЧАЛО РАБОТЫ СИСТЕМЫ КАНАЛОВ:")
	fmt.Println("==================================")
	for result := range results {
		fmt.Println(result)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎯 ВЫВОДЫ ДНЯ 86:")
	fmt.Println("1. Каналы — это как маршруты курьера в метель")
	fmt.Println("2. Буферизованные каналы — заказы могут ждать")
	fmt.Println("3. Небуферизованные каналы — мгновенная передача")
	fmt.Println("4. Select — выбор между разными каналами")
	fmt.Println("5. Таймауты — защита от вечного ожидания")
	fmt.Println("6. Закрытие каналов — конец рабочего дня")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println("\n💡 МОРАЛЬ: Каналы в Go синхронизируют горутины,")
	fmt.Println("   как диспетчер синхронизирует курьеров в метель.")
	fmt.Println("   Учись эффективно использовать каналы!")
}

// Дополнительная функция для форматирования строк
var strings = struct {
	Repeat func(string, int) string
}{
	Repeat: func(s string, count int) string {
		var result string
		for i := 0; i < count; i++ {
			result += s
		}
		return result
	},
}
