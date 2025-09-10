package main

import (
	"fmt"
	"scope_demo/examples"
	"scope_demo/scope"
)

func main() {
	fmt.Println("🔍 Day 48: Блоки и область видимости в Go")
	fmt.Println("==========================================")

	// Демонстрация областей видимости
	fmt.Println("\n1. Область видимости на уровне пакета:")
	scope.DemoPackageLevelScope()

	fmt.Println("\n2. Область видимости на уровне функций:")
	scope.DemoFunctionLevelScope()

	fmt.Println("\n3. Область видимости на уровне блоков:")
	scope.DemoBlockLevelScope()

	// Демонстрация замыканий
	fmt.Println("\n4. Замыкания и область видимости:")
	examples.DemoClosures()

	fmt.Println("\n🎉 Изучение областей видимости завершено!")
}
