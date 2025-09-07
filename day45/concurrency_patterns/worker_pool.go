package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker Pool Pattern
func WorkerPoolDemo() {
	fmt.Println("=== Worker Pool Pattern ===")
	
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Printf("Результат: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Воркер %d начал задачу %d\n", id, job)
		time.Sleep(time.Second) // Имитация работы
		fmt.Printf("Воркер %d завершил задачу %d\n", id, job)
		results <- job * 2
	}
}