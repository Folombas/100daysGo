package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🚀 Day 45: Паттерны параллельного программирования в Go")
	fmt.Println("======================================================")
	
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	
	switch os.Args[1] {
	case "worker":
		WorkerPoolDemo()
	case "pipeline":
		PipelineDemo()
	case "fan":
		FanOutFanInDemo()
	case "context":
		ContextDemo()
	case "mutex":
		MutexDemo()
	case "all":
		runAllDemos()
	default:
		fmt.Printf("❌ Неизвестная команда: %s\n", os.Args[1])
		printUsage()
	}
}

func runAllDemos() {
	WorkerPoolDemo()
	PipelineDemo()
	FanOutFanInDemo()
	ContextDemo()
	MutexDemo()
}

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  go run . worker     - Worker Pool Pattern")
	fmt.Println("  go run . pipeline   - Pipeline Pattern")
	fmt.Println("  go run . fan        - Fan-out Fan-in Pattern")
	fmt.Println("  go run . context    - Context Pattern")
	fmt.Println("  go run . mutex      - Mutex Pattern")
	fmt.Println("  go run . all        - Все демонстрации")
}