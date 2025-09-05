package errors

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// DemoErrorWrapping демонстрирует обертывание ошибок
func DemoErrorWrapping() {
	// Попытка открыть несуществующий файл
	_, err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		
		// Распаковка ошибки
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Файл действительно не существует!")
		}
		
		// Получение оригинальной ошибки
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Printf("Ошибка в пути: %s, Операция: %s\n", 
				pathError.Path, pathError.Op)
		}
	}
}

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		// Обертывание ошибки с дополнительным контекстом
		return nil, fmt.Errorf("не удалось прочитать файл %s: %w", filename, err)
	}
	return data, nil
}