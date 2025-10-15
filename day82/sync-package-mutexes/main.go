package main

import (
	"fmt"
	"sync"
	"time"
)

// DiamondMine - шахта алмазов, где работают гномы
type DiamondMine struct {
	diamonds    int
	mutex       sync.Mutex
	gnomeNames  []string
}

// GnomeWorker - гном-работник
type GnomeWorker struct {
	name   string
	speed  time.Duration
	mine   *DiamondMine
	wg     *sync.WaitGroup
}

func main() {
	fmt.Println("🚪 Мы с Пахомычем прошли через портал и оказались...")
	fmt.Println("💎 В АЛМАЗНОЙ ШАХТЕ СЕМИ ГНОМОВ!")
	fmt.Println("================================================")

	// Создаем алмазную шахту
	mine := &DiamondMine{
		diamonds: 0,
		gnomeNames: []string{
			"Умник", "Ворчун", "Весельчак", "Соня",
			"Скромник", "Чихун", "Простачок",
		},
	}

	var wg sync.WaitGroup

	fmt.Println("👷 Гномы начинают работу в шахте...")
	fmt.Println("⚠️  Но что это? Без мьютексов возникает ПУТАНИЦА!")

	// Запускаем гномов БЕЗ синхронизации (демонстрация проблемы)
	startTime := time.Now()
	runWithoutMutex(mine, &wg)

	fmt.Printf("\n💥 РЕЗУЛЬТАТ БЕЗ МЬЮТЕКСОВ: %d алмазов (должно быть 70)\n", mine.diamonds)
	fmt.Printf("⏱️  Время работы: %v\n", time.Since(startTime))

	// Сбрасываем счетчик для следующего эксперимента
	mine.diamonds = 0

	fmt.Println("\n🔧 Теперь используем МЬЮТЕКСЫ для синхронизации...")

	// Запускаем гномов С синхронизацией
	startTime = time.Now()
	runWithMutex(mine, &wg)

	fmt.Printf("\n✅ РЕЗУЛЬТАТ С МЬЮТЕКСАМИ: %d алмазов (ВСЁ ВЕРНО!)\n", mine.diamonds)
	fmt.Printf("⏱️  Время работы: %v\n", time.Since(startTime))

	fmt.Println("\n🎉 Леший впечатлен! Он говорит:")
	fmt.Println("«Вы справились с самой сложной магией - синхронизацией!»")
	fmt.Println("🏆 Теперь вы знаете силу sync.Mutex!")
}

// runWithoutMutex - запуск без синхронизации (демонстрация проблемы)
func runWithoutMutex(mine *DiamondMine, wg *sync.WaitGroup) {
	for i, name := range mine.gnomeNames {
		wg.Add(1)
		gnome := &GnomeWorker{
			name:  name,
			speed: time.Duration(100*(i+1)) * time.Millisecond,
			mine:  mine,
			wg:    wg,
		}
		go gnome.workWithoutMutex()
	}
	wg.Wait()
}

// runWithMutex - запуск с синхронизацией (правильное решение)
func runWithMutex(mine *DiamondMine, wg *sync.WaitGroup) {
	for i, name := range mine.gnomeNames {
		wg.Add(1)
		gnome := &GnomeWorker{
			name:  name,
			speed: time.Duration(100*(i+1)) * time.Millisecond,
			mine:  mine,
			wg:    wg,
		}
		go gnome.workWithMutex()
	}
	wg.Wait()
}

// workWithoutMutex - работа гнома БЕЗ синхронизации
func (g *GnomeWorker) workWithoutMutex() {
	defer g.wg.Done()

	for i := 0; i < 10; i++ {
		// Имитация работы
		time.Sleep(g.speed)

		// ❌ ОПАСНО: конкурентный доступ без защиты!
		current := g.mine.diamonds
		time.Sleep(10 * time.Millisecond) // Имитация задержки
		g.mine.diamonds = current + 1

		fmt.Printf("%s добыл алмаз! Всего: %d\n", g.name, g.mine.diamonds)
	}
}

// workWithMutex - работа гнома С синхронизацией
func (g *GnomeWorker) workWithMutex() {
	defer g.wg.Done()

	for i := 0; i < 10; i++ {
		// Имитация работы
		time.Sleep(g.speed)

		// ✅ БЕЗОПАСНО: используем мьютекс для защиты доступа
		g.mine.mutex.Lock()
		g.mine.diamonds++
		current := g.mine.diamonds
		g.mine.mutex.Unlock()

		fmt.Printf("%s добыл алмаз! Всего: %d\n", g.name, current)
	}
}
