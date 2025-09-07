package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fan-out Fan-in Pattern
func FanOutFanInDemo() {
	fmt.Println("\n=== Fan-out Fan-in Pattern ===")
	
	numbers := generateNumbers(20)
	
	// Fan-out: несколько воркеров обрабатывают один канал
	workers := 4
	channels := make([]<-chan int, workers)
	
	for i := 0; i < workers; i++ {
		channels[i] = processNumbers(numbers, i+1)
	}
	
	// Fan-in: объединяем результаты из нескольких каналов
	results := merge(channels)
	
	// Вывод результатов
	for result := range results {
		fmt.Printf("Обработанный результат: %d\n", result)
	}
}

func processNumbers(in <-chan int, workerID int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			out <- num * workerID
		}
		close(out)
	}()
	return out
}

func merge(channels []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	
	// Запускаем горутину для каждого канала
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for num := range c {
				out <- num
			}
		}(ch)
	}
	
	// Закрываем out после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}