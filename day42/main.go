package main

import (
	"day42/errors"
	"day42/examples"
	"fmt"
)

func main() {
	fmt.Println("Day 42: Error Handling in Go")
	fmt.Println("=============================")

	// Демонстрация пользовательских ошибок
	fmt.Println("\n1. Custom Errors:")
	errors.DemoCustomErrors()

	// Демонстрация обёртывания ошибок
	fmt.Println("\n2. Error Wrapping:")
	errors.DemoErrorWrapping()

	// Демонстрация валидации ошибок
	fmt.Println("\n3. Error Validation:")
	errors.DemoValidationErrors()

	// Демонстрация работы с файлами
	fmt.Println("\n4. File Operations:")
	examples.DemoFileOperations()

	// Демонстрация API клиента
	fmt.Println("\n5. API Client Example:")
	examples.DemoAPIClient()

	fmt.Println("\n🎉 Все примеры завершены!")
}
