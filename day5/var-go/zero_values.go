package main

import "fmt"

func demoZeroValues() {
	fmt.Println("🔮 3. НУЛЕВЫЕ ЗНАЧЕНИЯ:")
	fmt.Println("----------------------")

	// Go автоматически инициализирует нулевыми значениями
	var intValue int
	var floatValue float64
	var stringValue string
	var boolValue bool
	var pointerValue *int
	var sliceValue []int
	var mapValue map[string]int

	fmt.Printf("   int: %d\n", intValue)
	fmt.Printf("   float64: %.1f\n", floatValue)
	fmt.Printf("   string: %q\n", stringValue)
	fmt.Printf("   bool: %t\n", boolValue)
	fmt.Printf("   *int: %v\n", pointerValue)
	fmt.Printf("   []int: %v (len=%d)\n", sliceValue, len(sliceValue))
	fmt.Printf("   map[string]int: %v\n", mapValue)

	// Проверка на нулевое значение
	if stringValue == "" {
		fmt.Println("   ✅ stringValue является нулевым значением")
	}

	if sliceValue == nil {
		fmt.Println("   ✅ sliceValue является nil")
	}
	fmt.Println()
}

func demoTypeInference() {
	fmt.Println("🎭 4. ВЫВОД ТИПОВ:")
	fmt.Println("------------------")

	// Go автоматически выводит тип
	var explicitType int = 42
	var inferredType = 42 // тип выведен как int
	shortDecl := 3.14     // тип выведен как float64

	fmt.Printf("   explicitType: %T = %d\n", explicitType, explicitType)
	fmt.Printf("   inferredType: %T = %d\n", inferredType, inferredType)
	fmt.Printf("   shortDecl: %T = %.2f\n", shortDecl, shortDecl)

	// Разные типы на основе литералов
	number := 42    // int
	decimal := 42.0 // float64
	text := "42"    // string

	fmt.Printf("   number: %T = %v\n", number, number)
	fmt.Printf("   decimal: %T = %v\n", decimal, decimal)
	fmt.Printf("   text: %T = %v\n", text, text)
	fmt.Println()
}
