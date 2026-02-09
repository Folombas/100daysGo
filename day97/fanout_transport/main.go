package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Passenger представляет пассажира с билетом
type Passenger struct {
	ID        int
	HasTicket bool
	Mood      string // депрессивный, нейтральный, ок
}

// TransportSimulator симулятор транспортных поездок
type TransportSimulator struct {
	Routes     []string
	Passengers []Passenger
	Results    chan string
	Wg         sync.WaitGroup
}

func main() {
	fmt.Println("🚌 ДЕНЬ 97: FAN-OUT В ТРАНСПОРТЕ ДЕПРЕССИИ 🚌")
	fmt.Println("==============================================")

	rand.Seed(time.Now().UnixNano())

	// Инициализируем симулятор
	sim := &TransportSimulator{
		Routes: []string{
			"Метро",
			"Автобус",
			"Трамвай",
			"Электричка в пригород",
			"Пешая прогулка в парк",
		},
		Results: make(chan string, 10),
	}

	// Создаем пассажиров (воркеров)
	sim.generatePassengers(8)

	fmt.Printf("🎫 Сгенерировано пассажиров: %d\n", len(sim.Passengers))
	fmt.Printf("🚏 Доступных маршрутов: %d\n", len(sim.Routes))
	fmt.Println("\n🔁 Запускаем FAN-OUT распределение по маршрутам...")
	fmt.Println("   (каждая горутина - это выход из дома)")

	// Запускаем обработку пассажиров с fan-out
	sim.processPassengers()

	fmt.Println("\n✅ Все пассажиры обработаны!")
	fmt.Println("📊 Собираем результаты (FAN-IN)...")

	// Fan-in: собираем результаты
	sim.collectResults()

	fmt.Println("\n" + getMotivationalQuote())
	fmt.Println("\n💪 ГОША, ТЫ СДЕЛАЛ ЭТО! ЕЩЁ ОДИН ДЕНЬ БЕЗ CAPCUT И ИГР!")
	fmt.Println("   Завтра купишь проездной и поедешь дальше.")
}

func (s *TransportSimulator) generatePassengers(count int) {
	for i := 1; i <= count; i++ {
		s.Passengers = append(s.Passengers, Passenger{
			ID:        i,
			HasTicket: rand.Float32() > 0.3, // 70% с билетами
			Mood:      []string{"депрессивный", "нейтральный", "ок"}[rand.Intn(3)],
		})
	}
}

func (s *TransportSimulator) processPassengers() {
	// Создаем канал для задач
	tasks := make(chan Passenger, len(s.Passengers))

	// Fan-out: запускаем пул воркеров
	workerCount := 3 // всего 3 силы воли осталось
	for i := 1; i <= workerCount; i++ {
		s.Wg.Add(1)
		go s.worker(i, tasks)
	}

	// Отправляем пассажиров в канал
	for _, p := range s.Passengers {
		tasks <- p
	}
	close(tasks)

	// Ждем завершения всех воркеров
	s.Wg.Wait()
	close(s.Results)
}

func (s *TransportSimulator) worker(id int, tasks <-chan Passenger) {
	defer s.Wg.Done()

	for passenger := range tasks {
		// Имитация обработки (поездки)
		route := s.Routes[rand.Intn(len(s.Routes))]
		duration := rand.Intn(5) + 1

		time.Sleep(time.Duration(duration) * 50 * time.Millisecond)

		result := fmt.Sprintf("👤 Пассажир %d (настроение: %s) → %s за %d мин",
			passenger.ID, passenger.Mood, route, duration)

		if !passenger.HasTicket {
			result += " ⚠️ БЕЗ БИЛЕТА!"
		} else {
			result += " ✅ Билет есть"
		}

		s.Results <- result
	}
}

func (s *TransportSimulator) collectResults() {
	for result := range s.Results {
		fmt.Printf("  %s\n", result)
	}
}

func getMotivationalQuote() string {
	quotes := []string{
		"🌟 КОДИРОВАНИЕ — ЭТО ПРОГУЛКА ДЛЯ МОЗГА, КОГДА ТЕЛО НЕ МОЖЕТ ВЫЙТИ",
		"🚀 КАЖДАЯ ГОРУТИНА — ШАГ ОТ ДЕПРЕССИИ",
		"💡 GO НЕ ТРЕБУЕТ УСТАНОВКИ CAPCUT, ЧТОБЫ НАЧАТЬ",
		"🔥 СЕГОДНЯ ТЫ ВЫБРАЛ go mod init ВМЕСТО reinstall CapCut",
		"🎯 ФАН-АУТ НАГРУЗКИ КАК ФАН-АУТ ОТВЕТСТВЕННОСТИ",
		"⚡ ПАРАЛЛЕЛИЗМ В КОДЕ = ПАРАЛЛЕЛИЗМ В ВОССТАНОВЛЕНИИ",
		"🛡️ sync.WaitGroup ДЕРЖИТ ТЕБЯ В ПРОЦЕССЕ",
		"🔗 КАНАЛЫ ПЕРЕДАЮТ РЕЗУЛЬТАТЫ, А НЕ ЖАЛОСТЬ",
		"📈 ГОША 38 ЛЕТ > ГОША С CAPCUT",
		"🎖️ ТВОЙ ЛУЧШИЙ КОММИТ — ЭТО КОММИТ СЕБЕ: 'Я НЕ БРОСИЛ ОБУЧЕНИЕ GO'",
	}
	return quotes[rand.Intn(len(quotes))]
}
