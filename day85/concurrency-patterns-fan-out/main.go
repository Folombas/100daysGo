package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// FamilyTechSupport представляет смешанную вселенную Барбоскиных и Фиксиков
type FamilyTechSupport struct {
	barboskiny []string
	fixiki     []string
	problems   []string
}

func main() {
	fmt.Println("🏠 СМЕШАННАЯ ВСЕЛЕННАЯ: Барбоскины встречают Фиксиков!")
	fmt.Println("=====================================================")

	family := &FamilyTechSupport{
		barboskiny: []string{
			"Гена (папа)", "Мария (мама)", "Лиза", "Роза", "Дружок",
		},
		fixiki: []string{
			"Симка", "Нолик", "Файер", "Игрек", "Верта",
		},
		problems: []string{
			"сломался компьютер",
			"не работает Wi-Fi",
			"завис телевизор",
			"не печатает принтер",
			"сел аккумулятор у ноутбука",
			"глючит планшет",
			"не запускается микроволновка",
			"сломался холодильник",
		},
	}

	fmt.Println("📱 СИТУАЦИЯ: У Барбоскиных сломалась ВСЯ техника в доме!")
	fmt.Println("🆘 ПРОБЛЕМА: Один Фиксик не успеет всё починить!")
	fmt.Println("🎯 РЕШЕНИЕ: Используем Fan-out паттерн для распределения задач!")

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Создаем канал с проблемами
	problemsChannel := family.generateProblems(ctx)

	// Запускаем Fan-out паттерн
	fmt.Println("\n🔧 ЗАПУСК FAN-OUT ПАТТЕРНА...")
	fmt.Printf("🎪 Фиксики начинают работу! Распределяем %d проблем между %d фиксиками\n",
		len(family.problems), len(family.fixiki))

	// Запускаем worker-ов (фиксиков)
	results := family.fanOutWorkers(ctx, problemsChannel, 3) // 3 параллельных worker-а

	// Собираем результаты
	family.collectResults(results)

	fmt.Println("\n🎉 ВСЕ ПРОБЛЕМЫ РЕШЕНЫ! Техника снова работает!")
	fmt.Println("🏆 Барбоскины и Фиксики празднуют успех!")
}

// generateProblems генерирует поток проблем от Барбоскиных
func (f *FamilyTechSupport) generateProblems(ctx context.Context) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for _, problem := range f.problems {
			select {
			case <-ctx.Done():
				return
			case out <- problem:
				// Имитация времени между появлением проблем
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("   🚨 %s сообщает: '%s'\n",
					f.barboskiny[rand.Intn(len(f.barboskiny))], problem)
			}
		}
	}()

	return out
}

// fanOutWorkers реализует Fan-out паттерн: один канал -> несколько worker-ов
func (f *FamilyTechSupport) fanOutWorkers(ctx context.Context, in <-chan string, numWorkers int) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Запускаем несколько worker-ов
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go f.worker(ctx, &wg, in, out, f.fixiki[i])
	}

	// Закрываем выходной канал, когда все worker-ы завершатся
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// worker обрабатывает проблемы (ядро Fan-out паттерна)
func (f *FamilyTechSupport) worker(ctx context.Context, wg *sync.WaitGroup, in <-chan string, out chan<- string, fixikName string) {
	defer wg.Done()

	fmt.Printf("   🔧 %s начинает работу...\n", fixikName)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("   ⏰ %s: Время вышло, завершаю работу!\n", fixikName)
			return
		case problem, ok := <-in:
			if !ok {
				fmt.Printf("   ✅ %s: Все проблемы решены, завершаю работу!\n", fixikName)
				return
			}

			// Имитация времени на починку
			repairTime := time.Duration(500+rand.Intn(1000)) * time.Millisecond
			time.Sleep(repairTime)

			result := fmt.Sprintf("%s починил: %s (затратил %v)", fixikName, problem, repairTime)
			out <- result
		}
	}
}

// collectResults собирает и выводит результаты
func (f *FamilyTechSupport) collectResults(results <-chan string) {
	fmt.Println("\n📊 ОТЧЁТ О РАБОТЕ ФИКСИКОВ:")
	fmt.Println("==========================")

	count := 0
	for result := range results {
		count++
		fmt.Printf("   ✅ %s\n", result)
	}

	// ИСПРАВЛЕНИЕ: используем f.problems вместо family.problems
	fmt.Printf("\n📈 ИТОГО: Решено %d проблем из %d\n", count, len(f.problems))
}
