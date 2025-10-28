package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🕵️  КИБЕР-ДЕТЕКТИВ: ОХОТНИК ЗА БАГАМИ")
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Println("Привет, Гоша! Добро пожаловать в матрицу тестирования!")
	fmt.Println("Твоя миссия — находить и уничтожать баги.")
	fmt.Println("Каждый тест — это шаг к стабильности и уверенности!")
	fmt.Println()

	// Демонстрация работы функций
	fmt.Println("🔍 Тестируем калькулятор:")
	fmt.Printf("2 + 3 = %d\n", Add(2, 3))
	fmt.Printf("10 - 4 = %d\n", Subtract(10, 4))
	fmt.Printf("5 * 6 = %d\n", Multiply(5, 6))
	fmt.Printf("15 / 3 = %d\n", Divide(15, 3))

	fmt.Println()
	fmt.Println("🔍 Тестируем валидатор:")
	fmt.Printf("'hello' валидно? %v\n", ValidateString("hello"))
	fmt.Printf("'' валидно? %v\n", ValidateString(""))
	fmt.Printf("'a' валидно? %v\n", ValidateString("a"))

	fmt.Println()
	fmt.Println("💡 Запусти тесты: go test -v")
	fmt.Println("💡 Запусти конкретный тест: go test -v -run TestAdd")
	fmt.Println()
	fmt.Println("Помни: каждый написанный тест — это кирпичик в фундаменте")
	fmt.Println("твоей будущей карьеры Go-разработчика!")

	// Ждем нажатия Enter для удобства
	fmt.Println("\nНажми Enter для выхода...")
	fmt.Scanln()
	os.Exit(0)
}
