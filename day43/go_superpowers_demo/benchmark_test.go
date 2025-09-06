package main

import (
	"sync"
	"testing"
	"time"
)

// BenchmarkConcurrent выполняет задачи конкурентно
func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		tasks := make(chan int, 100)
		results := make(chan int, 100)
		
		// Запускаем 10 воркеров
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for task := range tasks {
					// Имитация работы
					time.Sleep(1 * time.Millisecond)
					results <- task * 2
				}
			}(j)
		}
		
		// Отправляем задачи
		for j := 0; j < 100; j++ {
			tasks <- j
		}
		close(tasks)
		
		wg.Wait()
		close(results)
	}
}

// BenchmarkSequential выполняет задачи последовательно
func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results := make(chan int, 100)
		
		for j := 0; j < 100; j++ {
			// Имитация работы
			time.Sleep(1 * time.Millisecond)
			results <- j * 2
		}
		close(results)
	}
}