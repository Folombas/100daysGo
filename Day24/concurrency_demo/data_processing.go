package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Имитация обработки данных
func processData(data string) {
	// Случайная задержка от 0.5 до 2 секунд
	delay := time.Duration(500+rand.Intn(1500)) * time.Millisecond
	time.Sleep(delay)
}

// Демонстрация горутин
func RunGoroutineDemo() {
	fmt.Println("\nДемо горутин:")
	
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Горутина %d запущена\n", id)
			processData(fmt.Sprintf("Горутина %d", id))
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}
	
	// Даем время на выполнение
	time.Sleep(3 * time.Second)
}

// Демонстрация каналов
func RunChannelDemo() {
	fmt.Println("\nДемо каналов:")
	
	ch := make(chan string)
	
	go func() {
		processData("Канальная задача")
		ch <- "Результат обработки"
	}()
	
	result := <-ch
	fmt.Println("Получено из канала:", result)
}