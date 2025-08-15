package main

import "os"

// Записать данные в файл
func writeFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// Прочитать данные из файла
func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// Удалить тестовые файлы
func cleanup() {
	os.Remove("user.json")
	os.Remove("config.xml")
}