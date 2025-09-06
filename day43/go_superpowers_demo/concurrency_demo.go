package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task представляет задачу для выполнения
type Task struct {
	ID   int
	Data string
}

// Result представляет результат выполнения задачи
type Result struct {
	TaskID int
	Output string
	Time   time.Duration
}

// worker обрабатывает задачи из канала и отправляет результаты
func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for task := range tasks {
		start := time.Now()
		
		// Имитация обработки (случайная задержка)
		processingTime := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(processingTime)
		
		// "Обработка" данных
		output := fmt.Sprintf("Обработано воркером %d: %s", id, task.Data)
		
		results <- Result{
			TaskID: task.ID,
			Output: output,
			Time:   time.Since(start),
		}
	}
}

// StartWorkerPool демонстрирует паттерн "Worker Pool"
func StartWorkerPool() {
	fmt.Println("🎯 Демонстрация Worker Pool в Go")
	
	var wg sync.WaitGroup
	numWorkers := 5
	numTasks := 15
	
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)
	
	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}
	
	// Отправляем задачи
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{
			ID:   i,
			Data: fmt.Sprintf("Задача %d", i),
		}
	}
	close(tasks)
	
	// Горутина для закрытия results после завершения всех воркеров
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Собираем результаты
	totalTime := time.Duration(0)
	for result := range results {
		fmt.Printf("✅ %s (время: %v)\n", result.Output, result.Time)
		totalTime += result.Time
	}
	
	fmt.Printf("📊 Все задачи завершены. Общее время: %v\n", totalTime)
	fmt.Printf("⚡ Среднее время на задачу: %v\n", totalTime/time.Duration(numTasks))
}