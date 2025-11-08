package main

import (
	"fmt"
)

// Глобальные переменные должны объявляться с var
var (
	appName    = "VarGo Demo"
	version    = "1.0.0"
	debugMode  = false
	maxRetries = 3
)

// Константы
const (
	AppAuthor = "100DaysGo Student"
	License   = "MIT"
	Pi        = 3.14159
)

func main() {
	fmt.Printf("🎯 %s v%s\n", appName, version)
	fmt.Println("=====================================")
	fmt.Printf("📝 Автор: %s\n", AppAuthor)
	fmt.Printf("🔧 Лицензия: %s\n", License)
	fmt.Printf("🧮 Значение Pi: %.5f\n", Pi)
	fmt.Println()

	demoVarDeclaration()
	demoShortDeclaration()
	demoConstants()
	demoBestPractices()
}

func demoVarDeclaration() {
	fmt.Println("📌 1. ДЕКЛАРАЦИЯ С VAR:")
	fmt.Println("-----------------------")

	// Явное объявление с типом
	var name string
	name = "Golang Student"
	fmt.Printf("   var name string = %q\n", name)

	// Объявление с инициализацией
	var age int = 25
	fmt.Printf("   var age int = %d\n", age)

	// Несколько переменных
	var x, y float64 = 10.5, 20.3
	fmt.Printf("   var x, y float64 = %.1f, %.1f\n", x, y)

	// Блок переменных
	var (
		isActive bool    = true
		salary   float64 = 50000.0
	)
	fmt.Printf("   var (isActive=%t, salary=%.2f)\n", isActive, salary)
	fmt.Println()
}

func demoShortDeclaration() {
	fmt.Println("⚡ 2. КОРОТКАЯ ДЕКЛАРАЦИЯ :=")
	fmt.Println("---------------------------")

	// Короткое объявление (только внутри функций)
	name := "Go Developer"
	age := 30
	score := 95.5
	isPassed := true

	fmt.Printf("   name := %q\n", name)
	fmt.Printf("   age := %d\n", age)
	fmt.Printf("   score := %.1f\n", score)
	fmt.Printf("   isPassed := %t\n", isPassed)

	// Множественное присваивание
	a, b, c := 1, "hello", true
	fmt.Printf("   a,b,c := %d, %q, %t\n", a, b, c)

	// Обмен значений
	x, y := 10, 20
	fmt.Printf("   До: x=%d, y=%d\n", x, y)
	x, y = y, x // Обмен без временной переменной
	fmt.Printf("   После: x=%d, y=%d\n", x, y)
	fmt.Println()
}
