package main

import (
    "os"
    "log"
)

// Создать тестовые файлы
func createTestFiles() {
    files := []string{"test.txt", "user.json", "my_dir/file1.txt", "my_dir/file2.txt"}
    for _, file := range files {
        if _, err := os.Stat(file); err == nil {
            os.Remove(file) // Удалить если существует
        }
    }
    
    // Удалить демо-каталог
    os.RemoveAll("my_dir")
}

// Обработчик ошибок
func checkErr(err error, message string) {
    if err != nil {
        log.Fatal(message, ": ", err)
    }
}