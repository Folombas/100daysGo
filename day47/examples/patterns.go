package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DemoPatterns демонстрирует паттерны работы с горутинами
func DemoPatterns() {
	fmt.Println("\n🔸 Worker Pool:")
	demoWorkerPool()

	fmt.Println("\n🔸 Fan-out Fan-in:")
	demoFanOutFanIn()

	fmt.Println("\n🔸 Rate Limiting:")
	demoRateLimiting()
}

// demoWorkerPool демонстрирует пул работников
func demoWorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// Создаем пул из 3 работников
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задания
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Получаем результаты
	go func() {
		wg.Wait()
		close(results)
	}()

	// Выводим результаты
	for result := range results {
		fmt.Printf("📊 Результат задания: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("👷 Работник %d начал задание %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		results <- job * 2
		fmt.Printf("👷 Работник %d завершил задание %d\n", id, job)
	}
}

// demoFanOutFanIn демонстрирует шаблон Fan-out Fan-in
func demoFanOutFanIn() {
	data := make(chan int)
	results := make(chan int)

	// Запускаем несколько обработчиков (Fan-out)
	for i := 1; i <= 3; i++ {
		go processor(i, data, results)
	}

	// Отправляем данные
	go func() {
		for i := 1; i <= 6; i++ {
			data <- i
		}
		close(data)
	}()

	// Собираем результаты (Fan-in)
	for i := 1; i <= 6; i++ {
		result := <-results
		fmt.Printf("🔄 Обработанный результат: %d\n", result)
	}
}

func processor(id int, data <-chan int, results chan<- int) {
	for item := range data {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		results <- item * 10
		fmt.Printf("⚙️  Процессор %d обработал: %d\n", id, item)
	}
}

// demoRateLimiting демонстрирует ограничение скорости
func demoRateLimiting() {
	requests := make(chan int, 5)
	limiter := time.Tick(200 * time.Millisecond) // Ограничение: 5 запросов в секунду

	// Отправляем запросы
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Обрабатываем с ограничением скорости
	for req := range requests {
		<-limiter
		fmt.Printf("⏱️  Обработан запрос %d (с ограничением скорости)\n", req)
	}
}