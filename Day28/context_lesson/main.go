package main

import (
	//"context"
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Демонстрация работы с контекстом в Go")
	fmt.Println("=====================================")

	// Демонстрация WithCancel
	fmt.Println("\n1. WithCancel:")
	exampleWithCancel()

	// Демонстрация WithTimeout
	fmt.Println("\n2. WithTimeout:")
	exampleWithTimeout()

	// Демонстрация WithDeadline
	fmt.Println("\n3. WithDeadline:")
	exampleWithDeadline()

	// Демонстрация WithValue
	fmt.Println("\n4. WithValue:")
	exampleWithValue()
}