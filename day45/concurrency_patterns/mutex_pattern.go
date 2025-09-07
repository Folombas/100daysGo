package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex Pattern
func MutexDemo() {
	fmt.Println("\n=== Mutex Pattern ===")
	
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)
	
	// Запускаем несколько горутин, которые изменяют общий ресурс
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 5; j++ {
				mu.Lock()
				counter++
				fmt.Printf("Горутина %d: counter = %d\n", id, counter)
				mu.Unlock()
				
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Итоговое значение: %d\n", counter)
}