package main

import (
	"fmt"
	"sync"
	"time"
)

// Демонстрация горутин и WaitGroup
func demonstrateGoroutines() {
	fmt.Println("=== Демонстрация горутин и WaitGroup ===")
	
	var wg sync.WaitGroup
	start := time.Now()
	
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Все горутины завершены за: %v\n", time.Since(start))
}

// Демонстрация мьютексов для безопасного доступа к общим данным
func demonstrateMutex() {
	fmt.Println("\n=== Демонстрация мьютексов ===")
	
	type SafeCounter struct {
		mu    sync.Mutex
		count int
	}
	
	counter := SafeCounter{}
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.mu.Lock()
			counter.count++
			counter.mu.Unlock()
		}()
	}
	
	wg.Wait()
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.count)
}

// Демонстрация работы с каналами
func demonstrateChannels() {
	fmt.Println("\n=== Демонстрация каналов ===")
	
	// Небуферизированный канал
	messageChan := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		messageChan <- "Привет из канала!"
	}()
	
	msg := <-messageChan
	fmt.Println("Получено:", msg)
	
	// Буферизированный канал
	bufferedChan := make(chan int, 3)
	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	
	fmt.Println("Буферизированные значения:", <-bufferedChan, <-bufferedChan, <-bufferedChan)
	
	// Select с каналами
	fmt.Println("\n--- Select с каналами ---")
	
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		chan1 <- "из первого канала"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		chan2 <- "из второго канала"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Получено", msg1)
		case msg2 := <-chan2:
			fmt.Println("Получено", msg2)
		}
	}
}