package main

import (
	"fmt"
	"sync"
	"time"
)

// FactoryMetaphorDemo демонстрирует фабричную метафору
func FactoryMetaphorDemo() {
	fmt.Println("\n🏭 Фабричная метафора: Горутины как рабочие, Каналы как конвейеры")
	fmt.Println("===============================================================")
	
	fmt.Println(`
Визуализация:
- Горутины - это рабочие на фабрике
- Каналы - это конвейерные ленты между рабочими
- Данные - это детали, moving по конвейеру
	`)
	
	// Создаем конвейерные ленты
	rawMaterials := make(chan string)
	assemblyLine := make(chan string)
	qualityControl := make(chan string)
	finishedProducts := make(chan string)
	
	fmt.Println("🔧 Создана фабрика с конвейерами:")
	fmt.Println("   - Сырьевая лента (rawMaterials)")
	fmt.Println("   - Сборочная линия (assemblyLine)")
	fmt.Println("   - Контроль качества (qualityControl)")
	fmt.Println("   - Готовые изделия (finishedProducts)")
	fmt.Println()
	
	var wg sync.WaitGroup
	
	// Запускаем рабочих-горутин
	wg.Add(4)
	go worker("Сборщик", rawMaterials, assemblyLine, &wg, 1000)
	go worker("Монтажник", assemblyLine, qualityControl, &wg, 1500)
	go worker("Контролер", qualityControl, finishedProducts, &wg, 1200)
	go qualitySupervisor(finishedProducts, &wg)
	
	// Поставка сырья
	go func() {
		parts := []string{"Деталь A", "Деталь B", "Деталь C", "Деталь D"}
		for _, part := range parts {
			fmt.Printf("📦 Поставка сырья: %s\n", part)
			rawMaterials <- part
			time.Sleep(800 * time.Millisecond)
		}
		close(rawMaterials)
	}()
	
	// Ждем завершения
	wg.Wait()
	
	fmt.Println("\n🎯 Вывод: Горутины как рабочие обрабатывают данные,")
	fmt.Println("          а каналы как конвейеры перемещают данные между этапами обработки")
}

func worker(name string, in <-chan string, out chan<- string, wg *sync.WaitGroup, speed time.Duration) {
	defer wg.Done()
	
	for part := range in {
		fmt.Printf("🔧 %s обрабатывает: %s\n", name, part)
		time.Sleep(speed * time.Millisecond)
		processed := fmt.Sprintf("Обработанная %s", part)
		out <- processed
		fmt.Printf("✅ %s завершил обработку: %s\n", name, processed)
	}
	close(out)
}

func qualitySupervisor(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for product := range in {
		fmt.Printf("🔍 Контроль качества проверяет: %s\n", product)
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("🏆 Продукт прошел контроль: %s\n", product)
	}
}