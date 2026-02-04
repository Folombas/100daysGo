package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("🚚 СИМУЛЯТОР ДОСТАВКИ GO-КУРЬЕРА")
	fmt.Println("══════════════════════════════════")

	// Ситуация 1: Хаос без синхронизации (утренняя депрессия)
	fmt.Println("\n🎭 ЭТАП 1: ХАОС БЕЗ SYNС (как Гоша утром)")
	fmt.Println("----------------------------------------")

	var chaoticMoney int32 = 0
	var wgChaos sync.WaitGroup

	// 5 горутин-курьеров пытаются одновременно добавить деньги
	for i := 1; i <= 5; i++ {
		wgChaos.Add(1)
		go func(courierID int) {
			defer wgChaos.Done()

			// Имитируем доставку
			deliveryTime := time.Duration(rand.Intn(300)+100) * time.Millisecond
			time.Sleep(deliveryTime)

			// ПРОБЛЕМА: гонка данных!
			current := chaoticMoney
			time.Sleep(10 * time.Millisecond) // Искусственная задержка для демонстрации
			chaoticMoney = current + 900

			fmt.Printf("   Курьер %d доставил заказ (+900 руб). Счетчик: %d руб\n",
				courierID, chaoticMoney)
		}(i)
	}

	wgChaos.Wait()
	fmt.Printf("📊 ИТОГО в хаосе: %d руб (должно быть: 4500 руб)\n", chaoticMoney)
	fmt.Println("💥 ПРОБЛЕМА: потерянные деньги из-за гонки данных!")

	// Ситуация 2: Порядок с Mutex (как мамин звонок наводит порядок)
	fmt.Println("\n🔒 ЭТАП 2: ПОРЯДОК С SYNC.MUTEX (мамин звонок)")
	fmt.Println("----------------------------------------------")

	var syncedMoney int32 = 0
	var mu sync.Mutex
	var wgSync sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wgSync.Add(1)
		go func(courierID int) {
			defer wgSync.Done()

			deliveryTime := time.Duration(rand.Intn(300)+100) * time.Millisecond
			time.Sleep(deliveryTime)

			// РЕШЕНИЕ: мьютекс защищает общий ресурс
			mu.Lock()
			current := syncedMoney
			time.Sleep(10 * time.Millisecond)
			syncedMoney = current + 900
			mu.Unlock()

			fmt.Printf("   [Mutex] Курьер %d доставил заказ. Безопасный счетчик: %d руб\n",
				courierID, syncedMoney)
		}(i)
	}

	wgSync.Wait()
	fmt.Printf("📊 ИТОГО с Mutex: %d руб (ВЕРНО!)\n", syncedMoney)

	// Ситуация 3: Atomic операции (как горячий душ — быстро и эффективно)
	fmt.Println("\n⚛️  ЭТАП 3: АТОМАРНОСТЬ С SYNC/ATOMIC (горячий душ)")
	fmt.Println("-------------------------------------------------")

	var atomicMoney int32 = 0
	var wgAtomic sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wgAtomic.Add(1)
		go func(courierID int) {
			defer wgAtomic.Done()

			deliveryTime := time.Duration(rand.Intn(300)+100) * time.Millisecond
			time.Sleep(deliveryTime)

			// СУПЕР-РЕШЕНИЕ: атомарные операции
			atomic.AddInt32(&atomicMoney, 900)

			fmt.Printf("   [Atomic] Курьер %d доставил заказ. Atomic счетчик: %d руб\n",
				courierID, atomic.LoadInt32(&atomicMoney))
		}(i)
	}

	wgAtomic.Wait()
	fmt.Printf("📊 ИТОГО с Atomic: %d руб (САМОЕ БЫСТРОЕ РЕШЕНИЕ!)\n", atomicMoney)

	// Демонстрация sync.Once (как ужин, который готовится только один раз)
	fmt.Println("\n🍲 ЭТАП 4: SYNC.ONCE (мамин ужин один раз за вечер)")
	fmt.Println("-----------------------------------------------")

	var once sync.Once
	dinnerReady := false

	// Несколько горутин пытаются "приготовить ужин"
	for i := 1; i <= 3; i++ {
		go func(personID int) {
			once.Do(func() {
				time.Sleep(500 * time.Millisecond)
				dinnerReady = true
				fmt.Printf("   👩🍳 Человек %d приготовил ужин (выполнено 1 раз)\n", personID)
			})

			if dinnerReady {
				fmt.Printf("   🍽️  Человек %d кушает ужин\n", personID)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)

	// Финальная мудрость
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("🎓 ВЫВОДЫ ДНЯ:")
	fmt.Println(strings.Repeat("═", 50))

	lessons := []string{
		"1. Без sync — хаос и потерянные данные (как утро без плана)",
		"2. sync.Mutex — надёжно, но может создавать очереди",
		"3. sync/atomic — молниеносно для простых операций",
		"4. sync.Once — гарантия однократного выполнения",
		"5. sync.WaitGroup — ждём завершения всех горутин",
	}

	for _, lesson := range lessons {
		fmt.Printf("   %s\n", lesson)
		time.Sleep(400 * time.Millisecond)
	}

	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("🏁 СИНХРОНИЗАЦИЯ ПРОЙДЕНА. КОД И ЖИЗНЬ В ПОРЯДКЕ.")
	fmt.Println(strings.Repeat("═", 50))
}
