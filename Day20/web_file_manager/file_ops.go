package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	//"strings"
)

// Информация о файле
type FileInfo struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
}

// Чтение файла
func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

// Запись в файл
func writeToFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// Создание директории
func createDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// Удаление файла/папки
func deletePath(path string) error {
	return os.RemoveAll(path)
}

// Список файлов в директории
func listDir(path string) ([]byte, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []FileInfo
	for _, entry := range entries {
		info, _ := entry.Info()
		files = append(files, FileInfo{
			Name:  entry.Name(),
			Size:  info.Size(),
			IsDir: entry.IsDir(),
		})
	}
	return json.Marshal(files)
}

// Получение абсолютного пути
func absPath(path string) string {
	abs, _ := filepath.Abs(path)
	return abs
}
