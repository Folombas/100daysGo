package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	//"strings"
	"fmt"
)

// Информация о файле
type FileInfo struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
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

// Чтение файла с ограничением размера
func readFile(path string) (string, error) {
	// Ограничиваем размер читаемых файлов (2 МБ)
	const maxSize = 2 << 20 // 2 MB

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Проверяем размер файла
	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	if info.Size() > maxSize {
		return fmt.Sprintf("Файл слишком большой (%s)", formatSize(info.Size())), nil
	}

	data, err := ioutil.ReadAll(file)
	return string(data), err
}

// Форматирование размера
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d Б", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cБ", float64(bytes)/float64(div), "КМГТПЕ"[exp])
}
