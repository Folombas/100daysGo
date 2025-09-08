package main

import (
	"fmt"
	"time"
)

// RiverMetaphorDemo демонстрирует речную метафору
func RiverMetaphorDemo() {
	fmt.Println("\n🌊 Речная метафора: Горутины как реки, Каналы как кораблики")
	fmt.Println("==========================================================")
	
	fmt.Println(`
Визуализация:
- Горутины - это быстрые реки, текущие параллельно
- Каналы - это кораблики, перевозящие грузы между реками
- Данные - это грузы на корабликах
	`)
	
	// Создаем каналы-реки
	nile := make(chan string)    // Река Нил
	amazon := make(chan string)  // Река Амазонка
	volga := make(chan string)   // Река Волга
	
	fmt.Println("🌍 Создано 3 реки-горутины:")
	fmt.Println("   - Река Нил (канал nile)")
	fmt.Println("   - Река Амазонка (канал amazon)") 
	fmt.Println("   - Река Волга (канал volga)")
	fmt.Println()
	
	// Запускаем горутины-реки
	go riverFlow("Нил", nile, 1000)
	go riverFlow("Амазонка", amazon, 1500)
	go riverFlow("Волга", volga, 1200)
	
	// Запускаем кораблики по рекам
	go sailShip("Грузовик с книгами", nile, amazon, 3)
	go sailShip("Ящик с яблоками", amazon, volga, 2)
	go sailShip("Контейнер с компьютерами", volga, nile, 4)
	
	// Даем время поработать демонстрации
	time.Sleep(10 * time.Second)
	
	fmt.Println("\n🎯 Вывод: Горутины как реки обеспечивают 'путь' для данных,")
	fmt.Println("          а каналы как кораблики перемещают данные между этими путями")
}

func riverFlow(name string, river chan string, speed time.Duration) {
	for {
		select {
		case cargo := <-river:
			fmt.Printf("📦 Река %s приняла груз: %s\n", name, cargo)
			time.Sleep(speed * time.Millisecond)
			fmt.Printf("🌊 Река %s доставила груз: %s\n", name, cargo)
		default:
			// Река течет даже без грузов
			time.Sleep(speed * 2 * time.Millisecond)
			fmt.Printf("~ Река %s течет...\n", name)
		}
	}
}

func sailShip(cargo string, from, to chan string, trips int) {
	for i := 0; i < trips; i++ {
		fmt.Printf("⛵ Кораблик с %s отправляется в плавание\n", cargo)
		from <- cargo
		time.Sleep(500 * time.Millisecond)
		received := <-to
		fmt.Printf("🏁 Кораблик с %s завершил плавание\n", received)
		time.Sleep(800 * time.Millisecond)
	}
}