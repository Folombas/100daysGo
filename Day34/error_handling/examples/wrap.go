package examples

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// ErrorWrapping демонстрирует обертывание ошибок
func ErrorWrapping() error {
	fmt.Println("Демонстрация обертывания ошибок...")

	// Пример 1: Обертывание ошибок
	if err := readConfig(); err != nil {
		fmt.Printf("Полная цепочка ошибок: %v\n", err)
		fmt.Printf("Развернутая ошибка: %+v\n", err)
		
		// Проверка конкретной ошибки в цепочке
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Ошибка вызвана отсутствием файла")
		}
		
		// Извлечение оригинальной ошибки
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("Оригинальная ошибка конфигурации: %s\n", configErr.Message)
		}
	}

	// Пример 2: Множественное обертывание
	if err := complexOperation(); err != nil {
		fmt.Printf("Сложная операция завершилась ошибкой: %v\n", err)
	}

	return nil
}

// ConfigError представляет ошибку конфигурации
type ConfigError struct {
	File    string
	Message string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("ошибка конфигурации в %s: %s", e.File, e.Message)
}

// readConfig имитирует чтение конфигурационного файла
func readConfig() error {
	if err := readFile("config.yaml"); err != nil {
		return fmt.Errorf("не удалось прочитать конфиг: %w", err)
	}
	return nil
}

// readFile имитирует чтение файла
func readFile(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return fmt.Errorf("ошибка доступа к файлу %s: %w", filename, err)
	}
	return &ConfigError{File: filename, Message: "неверный формат файла"}
}

// complexOperation демонстрирует множественное обертывание ошибок
func complexOperation() error {
	if err := step1(); err != nil {
		return fmt.Errorf("комплексная операция провалена: %w", err)
	}
	return nil
}

func step1() error {
	if err := step2(); err != nil {
		return fmt.Errorf("шаг 1 не выполнен: %w", err)
	}
	return nil
}

func step2() error {
	return fmt.Errorf("шаг 2 не выполнен: %w", errors.New("таймаут соединения"))
}