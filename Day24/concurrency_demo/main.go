package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	// Запуск демо-функций
	runGoroutineDemo()
	runChannelDemo()
	runWorkerPool()
	
	duration := time.Since(start)
	fmt.Printf("\nОбщее время выполнения: %.2f секунд\n", duration.Seconds())
}