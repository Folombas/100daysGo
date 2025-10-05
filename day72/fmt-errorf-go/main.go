package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// FarmError представляет кастомную ошибку для фермерских операций
type FarmError struct {
	Operation string
	Cause     string
	FieldSize int
}

func (e *FarmError) Error() string {
	return fmt.Sprintf("Фермерская ошибка: %s (причина: %s, размер поля: %dм²)",
		e.Operation, e.Cause, e.FieldSize)
}

// Траншейные работы на картофельном поле
func digTrench(length, depth int) error {
	if length <= 0 {
		return fmt.Errorf("копание траншеи: неверная длина %dм - должна быть положительной", length)
	}
	if depth > 2 {
		return fmt.Errorf("копание траншеи: глубина %dм слишком большая - максимум 2м", depth)
	}
	return nil
}

// Обработка картофельного урожая
func processPotatoes(area int) error {
	if area > 50 {
		return &FarmError{
			Operation: "обработка картофеля",
			Cause:     "слишком большая площадь",
			FieldSize: area,
		}
	}

	if area < 10 {
		return fmt.Errorf("обработка картофеля: площадь %dм² слишком мала для эффективной работы", area)
	}

	return nil
}

// Комплексная фермерская задача
func performFarmTasks() error {
	// Задача 1: Копание траншеи
	if err := digTrench(-5, 1); err != nil {
		return fmt.Errorf("не удалось выполнить задание по траншее: %w", err)
	}

	// Задача 2: Обработка картофеля
	if err := processPotatoes(100); err != nil {
		return fmt.Errorf("проблемы с картошкой: %w", err)
	}

	return nil
}

// Демонстрация различных вариантов fmt.Errorf
func demonstrateErrorFeatures() {
	fmt.Println("🚜 Демонстрация fmt.Errorf в фермерских условиях!")
	fmt.Println(strings.Repeat("=", 50))

	// Базовое использование fmt.Errorf
	err1 := fmt.Errorf("лопата сломалась во время копки")
	fmt.Printf("1. Базовая ошибка: %v\n", err1)

	// Ошибка с форматированием
	trenchLength := 150
	err2 := fmt.Errorf("траншея длиной %dм превышает допустимую длину 100м", trenchLength)
	fmt.Printf("2. Ошибка с параметрами: %v\n", err2)

	// Оборачивание существующей ошибки
	originalErr := errors.New("закончилась солярка в тракторе")
	wrappedErr := fmt.Errorf("не удалось вспахать поле: %w", originalErr)
	fmt.Printf("3. Оборачивание ошибки: %v\n", wrappedErr)

	// Проверка обернутой ошибки
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("4. ✓ Ошибка успешно обернута и распознана")
	}

	// Работа с кастомными ошибками
	farmErr := &FarmError{
		Operation: "полив грядок",
		Cause:     "сломался шланг",
		FieldSize: 25,
	}
	fmt.Printf("5. Кастомная ошибка: %v\n", farmErr)

	fmt.Println(strings.Repeat("=", 50))
}

func main() {
	// Демонстрация возможностей fmt.Errorf
	demonstrateErrorFeatures()

	// Симуляция рабочих задач
	fmt.Println("\n🏗️ Симуляция рабочих задач на картофельном поле:")

	if err := performFarmTasks(); err != nil {
		fmt.Printf("\n❌ Произошла ошибка во время работы:\n%v\n\n", err)

		// Разбор цепочки ошибок
		fmt.Println("🔍 Анализ цепочки ошибок:")
		var farmErr *FarmError
		if errors.As(err, &farmErr) {
			fmt.Printf("Найдена фермерская ошибка: %v\n", farmErr)
		}

		// Дополнительная информация об ошибке
		fmt.Printf("Подробности: %+v\n", err)
	} else {
		fmt.Println("✅ Все задачи выполнены успешно!")
	}

	// Практический пример с файлами (симуляция)
	fmt.Println("\n💾 Практический пример (симуляция работы с файлами):")
	if err := simulateFileOperation(); err != nil {
		fmt.Printf("Ошибка файловой операции: %v\n", err)
	}
}

// Дополнительная функция для демонстрации
func simulateFileOperation() error {
	// Симуляция попытки открыть несуществующий файл
	fileName := "план_посадки_картофеля.txt"

	// Имитация ошибки "файл не найден"
	fileErr := fmt.Errorf("открытие файла %s: %w", fileName, os.ErrNotExist)

	return fmt.Errorf("работа с документацией: %w", fileErr)
}
