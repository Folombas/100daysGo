package main

import (
	"fmt"
	"sync"
)

// Запуск всех демо-функций с ожиданием
func runGoroutineDemo() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		RunGoroutineDemo()
	}()

	wg.Wait()
}

func runChannelDemo() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		RunChannelDemo()
	}()

	wg.Wait()
}

func runWorkerPool() {
	fmt.Println("\nДемо пула воркеров:")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		RunWorkerPool()
	}()

	wg.Wait()
}
