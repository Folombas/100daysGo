package main

import (
	"fmt"
	"sync"
	"time"
)

// ProgrammingZoo представляет зоопарк языков программирования
type ProgrammingZoo struct {
	foodBowl    int // общая миска с едой
	mu          sync.Mutex
	animals     []*Animal
}

type Animal struct {
	name     string
	species  string // вид животного (язык программирования)
	speed    time.Duration
	foodEaten int
}

func main() {
	fmt.Println("🐾 ЗООПАРК ЯЗЫКОВ ПРОГРАММИРОВАНИЯ: Обнаружение гонок!")
	fmt.Println("=====================================================")

	zoo := &ProgrammingZoo{
		foodBowl: 100, // начальное количество еды
		animals: []*Animal{
			{name: "Питоша", species: "Python", speed: 300 * time.Millisecond},
			{name: "Перлуша", species: "Perl", speed: 500 * time.Millisecond},
			{name: "Гофер", species: "Golang", speed: 100 * time.Millisecond},
			{name: "Слоник", species: "PHP", speed: 200 * time.Millisecond},
			{name: "Крабик", species: "Rust", speed: 150 * time.Millisecond},
		},
	}

	fmt.Println("🍎 СИТУАЦИЯ: В зоопарке одна миска с едой на всех!")
	fmt.Println("🎯 ЗАДАЧА: Животные должны есть БЕЗ конфликтов и потери данных!")

	// Демонстрация проблемы: гонка данных без синхронизации
	fmt.Println("\n❌ ДЕМОНСТРАЦИЯ ПРОБЛЕМЫ: Гонка данных (Race Condition)")
	zoo.foodBowl = 100 // сбрасываем миску
	zoo.demoRaceCondition()

	fmt.Printf("\n🥣 После гонки в миске осталось: %d единиц еды\n", zoo.foodBowl)
	zoo.printAnimalStats()

	// Демонстрация решения: с синхронизацией
	fmt.Println("\n✅ ДЕМОНСТРАЦИЯ РЕШЕНИЯ: Синхронизация (Mutex)")
	zoo.foodBowl = 100 // сбрасываем миску
	zoo.demoWithSync()

	fmt.Printf("\n🥣 После синхронизации в миске осталось: %d единиц еды\n", zoo.foodBowl)
	zoo.printAnimalStats()

	fmt.Println("\n🔍 КАК ОБНАРУЖИТЬ ГОНКУ:")
	fmt.Println("   go run -race main.go")
	fmt.Println("   go build -race")
	fmt.Println("   go test -race")

	fmt.Println("\n🎉 ВЫВОД: Go Race Detector - наш лучший друг для поиска скрытых гонок!")
}

// demoRaceCondition демонстрирует гонку данных
func (z *ProgrammingZoo) demoRaceCondition() {
	var wg sync.WaitGroup

	fmt.Println("   🐍🐫🐹🐘🦀 Животные начинают есть БЕЗ синхронизации...")

	for _, animal := range z.animals {
		wg.Add(1)
		go z.animalEatWithoutSync(&wg, animal)
	}

	wg.Wait()
}

// animalEatWithoutSync - животное ест без синхронизации (опасно!)
func (z *ProgrammingZoo) animalEatWithoutSync(wg *sync.WaitGroup, animal *Animal) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// ❌ ОПАСНО: гонка данных! Несколько горутин читают и пишут foodBowl
		if z.foodBowl > 0 {
			time.Sleep(animal.speed)

			// Критическая секция без защиты
			currentFood := z.foodBowl
			time.Sleep(10 * time.Millisecond) // имитация задержки
			z.foodBowl = currentFood - 1
			animal.foodEaten++

			fmt.Printf("   %s съел кусочек! Осталось: %d\n", animal.name, z.foodBowl)
		}
	}
}

// demoWithSync демонстрирует решение с синхронизацией
func (z *ProgrammingZoo) demoWithSync() {
	var wg sync.WaitGroup

	// Сбрасываем статистику животных
	for _, animal := range z.animals {
		animal.foodEaten = 0
	}

	fmt.Println("   🐍🐫🐹🐘🦀 Животные начинают есть С синхронизацией...")

	for _, animal := range z.animals {
		wg.Add(1)
		go z.animalEatWithSync(&wg, animal)
	}

	wg.Wait()
}

// animalEatWithSync - животное ест с синхронизацией (безопасно!)
func (z *ProgrammingZoo) animalEatWithSync(wg *sync.WaitGroup, animal *Animal) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// ✅ БЕЗОПАСНО: используем мьютекс для защиты критической секции
		z.mu.Lock()

		if z.foodBowl > 0 {
			// Имитация времени на еду
			time.Sleep(animal.speed)

			z.foodBowl--
			animal.foodEaten++

			fmt.Printf("   %s съел кусочек! Осталось: %d\n", animal.name, z.foodBowl)
		}

		z.mu.Unlock()

		// Небольшая пауза между подходами к миске
		time.Sleep(50 * time.Millisecond)
	}
}

// printAnimalStats выводит статистику по животным
func (z *ProgrammingZoo) printAnimalStats() {
	fmt.Println("\n📊 СТАТИСТИКА ЖИВОТНЫХ:")
	for _, animal := range z.animals {
		fmt.Printf("   %s (%s): съел %d кусочков\n",
			animal.name, animal.species, animal.foodEaten)
	}
}

// Дополнительная функция для демонстрации скрытой гонки
func hiddenRaceDemo() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ❌ Скрытая гонка: несколько горутин пишут в counter
			counter++
		}()
	}

	wg.Wait()
	fmt.Printf("\n🧪 Скрытая гонка: counter = %d (может быть меньше 10!)\n", counter)
}
