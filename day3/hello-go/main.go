package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🎯 Day 3: Hello World in Go")
	fmt.Println("============================")

	// Базовая версия
	fmt.Println("\n1. 🚀 Классический Hello World:")
	fmt.Println("   Hello, World! Привет, мир!")

	// С задержкой для драматизма
	time.Sleep(1 * time.Second)

	// Вызов дополнительных функций
	ShowAdvancedHello()
	time.Sleep(1 * time.Second)
	ShowSpecialEffects()
}
