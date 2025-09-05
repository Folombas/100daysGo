package examples

import (
	"fmt"
	"os"
)

// DemoFileOperations демонстрирует обработку ошибок при работе с файлами
func DemoFileOperations() {
	// Создание файла с обработкой ошибок
	if err := createFile("test.txt"); err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	
	// Чтение файла с обработкой ошибок
	if data, err := readFile("test.txt"); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	} else {
		fmt.Println("Содержимое файла:", string(data))
	}
	
	// Удаление файла
	if err := deleteFile("test.txt"); err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
	}
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()
	
	_, err = file.WriteString("Привет, мир!\nЭто тестовый файл.")
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	
	fmt.Println("Файл успешно создан:", filename)
	return nil
}

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return data, nil
}

func deleteFile(filename string) error {
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("ошибка удаления файла: %w", err)
	}
	fmt.Println("Файл успешно удален:", filename)
	return nil
}