package main

import (
	"errors"
	"fmt"
)

// Пользовательская ошибка
var ErrNegativeNumber = errors.New("отрицательное число недопустимо")

// Функция с возвращением ошибки
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeNumber
	}
	
	// Простая реализация (для демонстрации)
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// Демонстрация обработки ошибок
func demonstrateErrorHandling() {
	fmt.Println("=== Демонстрация обработки ошибок ===")
	
	// Проверка на ошибку
	if result, err := sqrt(16); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Квадратный корень: %.2f\n", result)
	}
	
	// Ошибка
	if _, err := sqrt(-4); err != nil {
		fmt.Println("Ошибка:", err)
	}
	
	// Defer, panic и recover
	fmt.Println("\n--- Defer, Panic и Recover ---")
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено после паники:", r)
		}
	}()
	
	fmt.Println("Выполняем обычный код")
	
	// Имитация паники (в реальном коде избегайте panic)
	panic("Что-то пошло не так!")
	
	// Этот код не выполнится
	fmt.Println("Этот код не будет выполнен")
}