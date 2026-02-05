package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// Delivery представляет доставку документов
type Delivery struct {
	From      string
	To        string
	Price     int
	Completed bool
}

// DeliveryService имитирует службу доставки
type DeliveryService struct {
	balance      int
	deliveries   []Delivery
	mu           sync.Mutex  // Ключевой мьютекс дня
	balanceMu    sync.Mutex  // Дополнительный мьютекс для баланса
}

// NewDeliveryService создает новую службу доставки
func NewDeliveryService() *DeliveryService {
	return &DeliveryService{
		balance: 0,
		deliveries: []Delivery{
			{"Речной Вокзал", "Центр", 600, false},
			{"Химки", "Зюзино", 900, false},
			{"Тверская", "Арбат", 450, false},
		},
	}
}

// processDeliveryWithoutMutex имитирует обработку без синхронизации
func (ds *DeliveryService) processDeliveryWithoutMutex(deliveryNum int) {
	// ПРОБЛЕМА: Гонка данных!
	currentBalance := ds.balance
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // Имитация работы
	ds.balance = currentBalance + ds.deliveries[deliveryNum].Price
	ds.deliveries[deliveryNum].Completed = true
	
	fmt.Printf("   🚚 Доставка %d выполнена (+%d руб). Баланс: %d руб\n",
		deliveryNum+1, ds.deliveries[deliveryNum].Price, ds.balance)
}

// processDeliveryWithMutex обрабатывает доставку с мьютексом
func (ds *DeliveryService) processDeliveryWithMutex(deliveryNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	ds.mu.Lock() // БЛОКИРОВКА КРИТИЧЕСКОЙ СЕКЦИИ
	defer ds.mu.Unlock()
	
	// Безопасная работа с общими данными
	ds.balance += ds.deliveries[deliveryNum].Price
	ds.deliveries[deliveryNum].Completed = true
	
	// Имитация времени доставки
	deliveryTime := time.Duration(100+rand.Intn(200)) * time.Millisecond
	time.Sleep(deliveryTime)
	
	fmt.Printf("   🔒 [Mutex] Доставка %d: %s → %s за %d руб (заняло %v)\n",
		deliveryNum+1, 
		ds.deliveries[deliveryNum].From,
		ds.deliveries[deliveryNum].To,
		ds.deliveries[deliveryNum].Price,
		deliveryTime)
	fmt.Printf("      Баланс: %d руб | Выполнено: %d/%d\n",
		ds.balance, ds.countCompleted(), len(ds.deliveries))
}

// countCompleted считает выполненные доставки
func (ds *DeliveryService) countCompleted() int {
	completed := 0
	for _, d := range ds.deliveries {
		if d.Completed {
			completed++
		}
	}
	return completed
}

// tryDoubleBooking демонстрирует попытку двойного использования ресурса
func (ds *DeliveryService) tryDoubleBooking() {
	fmt.Println("\n🎭 СИТУАЦИЯ: Попытка 'двойной брони' доставки")
	fmt.Println("   (Как попытка одновременно монтировать и писать код)")
	
	var mu sync.Mutex
	resourceInUse := false
	
	// Первая горутина (монтаж видео)
	go func() {
		mu.Lock()
		if resourceInUse {
			fmt.Println("   ⚠️  МОНТАЖ: Ресурс (внимание) уже занят! Жду...")
		}
		resourceInUse = true
		fmt.Println("   🎬 МОНТАЖ: Начал рендеринг видео (блокирую внимание)")
		time.Sleep(800 * time.Millisecond)
		resourceInUse = false
		mu.Unlock()
		fmt.Println("   ✅ МОНТАЖ: Рендеринг завершен, внимание свободно")
	}()
	
	// Вторая горутина (код на Go)
	go func() {
		time.Sleep(100 * time.Millisecond) // Небольшая задержка
		mu.Lock()
		resourceInUse = true
		fmt.Println("   💻 GO: Начал писать код (блокирую внимание)")
		time.Sleep(500 * time.Millisecond)
		resourceInUse = false
		mu.Unlock()
		fmt.Println("   ✅ GO: Код написан, внимание свободно")
	}()
	
	time.Sleep(2 * time.Second)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("🔐 DAY 93: МЬЮТЕКСЫ ДОСТАВКИ")
	fmt.Println("══════════════════════════════════")
	
	fmt.Println("\n🎯 СИТУАЦИЯ: Гоша и его служба доставки")
	fmt.Println("   Баланс: 0 руб | Заказов: 3 | Состояние: депрессия")
	
	// Часть 1: Проблема без мьютекса
	fmt.Println("\n💥 ЧАСТЬ 1: ХАОС БЕЗ МЬЮТЕКСА")
	fmt.Println("   (Как депрессия — мысли наперегонки)")
	
	chaoticService := NewDeliveryService()
	
	// Запускаем горутины без синхронизации
	fmt.Println("   Запускаем 3 доставки параллельно...")
	for i := 0; i < 3; i++ {
		go chaoticService.processDeliveryWithoutMutex(i)
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("\n   📊 ИТОГО в хаосе: %d руб (должно быть: %d руб)\n",
		chaoticService.balance, 600+900+450)
	fmt.Println("   ❌ Деньги потерялись в гонке данных!")
	
	// Часть 2: Решение с мьютексом
	fmt.Println("\n🔒 ЧАСТЬ 2: ПОРЯДОК С МЬЮТЕКСОМ")
	fmt.Println("   (Как решение заниматься и монтажом, и кодом)")
	
	orderedService := NewDeliveryService()
	var wg sync.WaitGroup
	
	fmt.Println("   Запускаем доставки с sync.Mutex...")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go orderedService.processDeliveryWithMutex(i, &wg)
	}
	
	wg.Wait()
	fmt.Printf("\n   📊 ИТОГО с мьютексом: %d руб (ВЕРНО!)\n", orderedService.balance)
	
	// Часть 3: Аналогия с монтажом и кодом
	orderedService.tryDoubleBooking()
	
	// Финал
	fmt.Println("\n══════════════════════════════════")
	fmt.Println("🎓 ВЫВОДЫ ДНЯ:")
	fmt.Println("   1. Mutex даёт эксклюзивный доступ к ресурсу")
	fmt.Println("   2. Без mutex — гонка данных и потеря информации")
	fmt.Println("   3. С mutex — порядок, но возможны очереди")
	fmt.Println("   4. Как в жизни: можно делать либо код, либо монтаж")
	fmt.Println("      в момент времени, но переключаться грамотно")
	fmt.Println("\n   Гоша сегодня: 600 руб доставка + изучение мьютексов")
	fmt.Println("   Завтра: баланс между рендерингом и компиляцией")
	fmt.Println("══════════════════════════════════")
}
