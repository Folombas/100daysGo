package examples

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// BasicErrorHandling демонстрирует базовые подходы к обработке ошибок
func BasicErrorHandling() error {
	fmt.Println("Проверка базовой обработки ошибок...")

	// Пример 1: Проверка ошибок после вызова функции
	result, err := divide(10, 2)
	if err != nil {
		return fmt.Errorf("ошибка деления: %w", err)
	}
	fmt.Printf("10 / 2 = %.1f\n", result)

	// Пример 2: Деление на ноль
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.1f\n", result)
	}

	// Пример 3: Чтение несуществующего файла
	content, err := readFileContent("nonexistent.txt") // Изменили имя функции
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
	} else {
		fmt.Printf("Содержимое файла: %s\n", content)
	}

	// Пример 4: Преобразование строки в число
	num, err := strconv.Atoi("не число")
	if err != nil {
		fmt.Printf("Ошибка преобразования: %v\n", err)
	} else {
		fmt.Printf("Число: %d\n", num)
	}

	return nil
}

// divide выполняет деление двух чисел с проверкой на ноль
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

// readFileContent читает содержимое файла (переименовали функцию)
func readFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать файл %s: %w", filename, err)
	}
	return string(content), nil
}

// PanicAndRecover демонстрирует обработку паники
func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Паника перехвачена: %v\n", r)
		}
	}()

	fmt.Println("Запуск функции с паникой...")
	causePanic()
	fmt.Println("Это сообщение не будет показано")
}

func causePanic() {
	panic("критическая ошибка в системе")
}