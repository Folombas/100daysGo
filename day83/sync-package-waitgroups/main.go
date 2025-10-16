package main

import (
	"fmt"
	"sync"
	"time"
)

// Волк и Заяц из "Ну, погоди!" попали в мир Белоснежки
type FairyTaleAdventure struct {
	wg          sync.WaitGroup
	gnomes      []string
	tasks       []string
	completed   int
}

func main() {
	fmt.Println("🎭 СМЕШАННАЯ ВСЕЛЕННАЯ: Белоснежка встречает 'Ну, погоди!'")
	fmt.Println("=======================================================")

	adventure := &FairyTaleAdventure{
		gnomes: []string{"Умник", "Ворчун", "Весельчак", "Соня", "Скромник", "Чихун", "Простачок"},
		tasks: []string{
			"собрать алмазы в шахте",
			"приготовить ужин для Белоснежки",
			"починить мост через реку",
			"научить Зайца программировать на Go",
			"спрятать яблоки от Волка",
			"настроить телепорт для возвращения",
			"организовать праздник спасения",
		},
	}

	fmt.Println("🐺 Волк гонится за Зайцем по сказочному лесу...")
	fmt.Println("🚨 ВНЕЗАПНО: Появляется Злая Королева с отравленным яблоком!")
	fmt.Println("📢 Белоснежка кричит: 'Гномы, помогите! Нужно выполнить все задачи до заката!'")

	// Демонстрация проблемы без WaitGroup
	fmt.Println("\n❌ СИТУАЦИЯ: Гномы работают БЕЗ координации...")
	adventure.workWithoutWaitGroup()

	fmt.Println("\n💥 КАТАСТРОФА: Волк почти поймал Зайца, а Королева нашла Белоснежку!")
	fmt.Println("😱 Завершено задач:", adventure.completed, "из", len(adventure.tasks))

	// Сброс для следующего эксперимента
	adventure.completed = 0

	// Демонстрация решения с WaitGroup
	fmt.Println("\n✅ РЕШЕНИЕ: Используем магию WaitGroup для координации!")
	adventure.workWithWaitGroup()

	fmt.Println("\n🎉 ПОБЕДА: Все задачи выполнены! Волк подружился с Зайцем, Королева раскаялась!")
	fmt.Println("🏆 Завершено задач:", adventure.completed, "из", len(adventure.tasks))
	fmt.Println("✨ Белоснежка, Заяц и Волк программируют на Go вместе!")
}

// workWithoutWaitGroup - работа без синхронизации (хаос)
func (a *FairyTaleAdventure) workWithoutWaitGroup() {
	for i, gnome := range a.gnomes {
		go func(id int, name string, task string) {
			// Имитация работы гнома (разное время)
			workTime := time.Duration(500*(id+1)) * time.Millisecond
			time.Sleep(workTime)

			fmt.Printf("   %s: '%s' - ВЫПОЛНЕНО! (затратил %v)\n", name, task, workTime)
			a.completed++
		}(i, gnome, a.tasks[i])
	}

	// ❌ ПРОБЛЕМА: Не ждём завершения горутин
	time.Sleep(1 * time.Second) // Ждём только секунду
	fmt.Println("\n⏰ ВРЕМЯ ВЫШЛО! Главная горутина завершает работу...")
}

// workWithWaitGroup - работа с синхронизацией (порядок)
func (a *FairyTaleAdventure) workWithWaitGroup() {
	// Устанавливаем счётчик на количество задач
	a.wg.Add(len(a.gnomes))

	for i, gnome := range a.gnomes {
		go func(id int, name string, task string) {
			// Гарантируем уменьшение счётчика при завершении
			defer a.wg.Done()

			// Имитация работы гнома
			workTime := time.Duration(500*(id+1)) * time.Millisecond
			time.Sleep(workTime)

			fmt.Printf("   %s: '%s' - ВЫПОЛНЕНО! (затратил %v)\n", name, task, workTime)
			a.completed++
		}(i, gnome, a.tasks[i])
	}

	fmt.Println("🔄 WaitGroup ждёт завершения ВСЕХ задач...")
	a.wg.Wait() // ✅ Ждём завершения ВСЕХ горутин
	fmt.Println("🎯 WaitGroup: Все горутины завершены!")
}
