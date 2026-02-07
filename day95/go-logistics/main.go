package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ================== Структуры данных ==================
type Cargo struct {
	ID          string
	Destination string
	Data        string
	Status      string
}

type Truck struct {
	ID      string
	Driver  string
	channel chan *Cargo
	quit    chan bool
}

type Dispatcher struct {
	sync.Mutex
	trucks    []*Truck
	pending   chan *Cargo
	delivered chan *Cargo
	failed    chan *Cargo
	log       []string
}

// ================== Инициализация ==================
func NewDispatcher(truckCount int) *Dispatcher {
	d := &Dispatcher{
		trucks:    make([]*Truck, 0, truckCount),
		pending:   make(chan *Cargo, 20),
		delivered: make(chan *Cargo, 10),
		failed:    make(chan *Cargo, 5),
	}

	drivers := []string{"Петрович", "Семёныч", "Михалыч", "Иваныч", "Степаныч"}

	for i := 0; i < truckCount; i++ {
		truck := &Truck{
			ID:      fmt.Sprintf("FU-%03d", i+1),
			Driver:  drivers[i%len(drivers)],
			channel: make(chan *Cargo, 1),
			quit:    make(chan bool),
		}
		d.trucks = append(d.trucks, truck)
		go truck.work(d.delivered, d.failed)
	}

	return d
}

// ================== Работа фуры ==================
func (t *Truck) work(delivered, failed chan<- *Cargo) {
	for {
		select {
		case cargo := <-t.channel:
			cargo.Status = "в пути"

			// Симуляция доставки
			delay := time.Duration(rand.Intn(3)+1) * time.Second

			// 20% шанс на проблему
			if rand.Intn(100) < 20 {
				delay *= 2
				fmt.Printf("⚠️  %s (%s): Попал в туман к %s\n",
					t.ID, t.Driver, cargo.Destination)
			}

			time.Sleep(delay)

			// 10% шанс на сбой
			if rand.Intn(100) < 10 {
				cargo.Status = "сбой"
				failed <- cargo
				fmt.Printf("🔥 %s: Сломался! Груз %s не доставлен\n", t.ID, cargo.ID)
			} else {
				cargo.Status = "доставлен"
				delivered <- cargo
				fmt.Printf("✅ %s: Доставил %s → %s\n",
					t.ID, cargo.ID, cargo.Destination)
			}

		case <-t.quit:
			return
		}
	}
}

// ================== Генерация грузов ==================
func (d *Dispatcher) generateCargo(count int) {
	destinations := []string{
		"auth-service", "user-service", "payment-service",
		"notification-service", "inventory-service",
	}

	for i := 0; i < count; i++ {
		cargo := &Cargo{
			ID:          fmt.Sprintf("CRG-%04d", i+1),
			Destination: destinations[rand.Intn(len(destinations))],
			Data:        fmt.Sprintf("Данные #%d", i+1),
			Status:      "ожидает",
		}
		d.pending <- cargo
		d.logEvent(fmt.Sprintf("📦 Создан груз %s → %s", cargo.ID, cargo.Destination))
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	close(d.pending)
}

// ================== Распределение грузов (Fan-Out) ==================
func (d *Dispatcher) distribute() {
	for cargo := range d.pending {
		assigned := false

		for _, truck := range d.trucks {
			select {
			case truck.channel <- cargo:
				d.logEvent(fmt.Sprintf("🚚 %s взят груз %s", truck.ID, cargo.ID))
				assigned = true
			default:
				continue
			}
			if assigned {
				break
			}
		}

		if !assigned {
			// Если все фуры заняты, ждем и пробуем снова
			go func(c *Cargo) {
				time.Sleep(500 * time.Millisecond)
				d.pending <- c
				d.logEvent(fmt.Sprintf("⏳ Груз %s ждет свободную фуру", c.ID))
			}(cargo)
		}
	}
}

// ================== Сбор результатов (Fan-In) ==================
func (d *Dispatcher) collectResults(done chan<- bool) {
	delivered, failed := 0, 0

	for {
		select {
		case cargo := <-d.delivered:
			delivered++
			d.logEvent(fmt.Sprintf("🎉 Груз %s доставлен", cargo.ID))

		case cargo := <-d.failed:
			failed++
			d.logEvent(fmt.Sprintf("❌ Сбой: %s → %s", cargo.ID, cargo.Destination))

			// Повторная отправка (Retry Pattern)
			go func(c *Cargo) {
				time.Sleep(time.Second * 2)
				c.Status = "повтор"
				d.pending <- c
			}(cargo)

		case <-time.After(2 * time.Second):
			// Таймаут - все грузы обработаны
			d.Lock()
			fmt.Printf("\n📊 ИТОГИ:\n")
			fmt.Printf("   ✅ Доставлено: %d\n", delivered)
			fmt.Printf("   ❌ Сбоев: %d\n", failed)
			fmt.Printf("   🚛 Фур в пуле: %d\n", len(d.trucks))
			d.Unlock()

			done <- true
			return
		}
	}
}

// ================== Вспомогательные методы ==================
func (d *Dispatcher) logEvent(event string) {
	d.Lock()
	timestamp := time.Now().Format("15:04:05")
	entry := fmt.Sprintf("[%s] %s", timestamp, event)
	d.log = append(d.log, entry)
	d.Unlock()
}

func (d *Dispatcher) showLog() {
	fmt.Println("\n📋 ЖУРНАЛ СОБЫТИЙ:")
	fmt.Println(strings.Repeat("─", 50))
	for _, entry := range d.log {
		fmt.Println(entry)
	}
}

func (d *Dispatcher) shutdown() {
	fmt.Println("\n🌙 Завершаем работу...")

	// Останавливаем фуры
	for _, truck := range d.trucks {
		truck.quit <- true
	}

	close(d.delivered)
	close(d.failed)
	time.Sleep(time.Second)

	fmt.Println("✅ Диспетчерская завершила работу")
}

// ================== Основная программа ==================
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("🚚 GO-LOGISTICS: Паттерны конкурентности")
	fmt.Println(strings.Repeat("═", 60))

	// 1. Инициализация
	dispatcher := NewDispatcher(4)
	fmt.Println("🚛 Создан пул из 4 фур")

	// 2. Запуск процессов
	done := make(chan bool)

	go dispatcher.generateCargo(12)
	go dispatcher.distribute()
	go dispatcher.collectResults(done)

	// 3. Ожидание завершения
	<-done

	// 4. Отчет
	dispatcher.showLog()

	// 5. Завершение
	dispatcher.shutdown()

	// 6. Итог
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("🎯 РЕАЛИЗОВАННЫЕ ПАТТЕРНЫ:")
	fmt.Println("   • Worker Pool (4 фуры-горутины)")
	fmt.Println("   • Producer-Consumer (генерация/обработка грузов)")
	fmt.Println("   • Fan-Out (распределение по фурам)")
	fmt.Println("   • Fan-In (сбор результатов)")
	fmt.Println("   • Retry (повтор при сбое)")
	fmt.Println("   • Graceful Shutdown (корректное завершение)")
	fmt.Println(strings.Repeat("═", 60))

	fmt.Println("\n📝 DISCLAIMER: Все персонажи вымышлены.")
	fmt.Println("   Образовательная программа для изучения Go.")
}

