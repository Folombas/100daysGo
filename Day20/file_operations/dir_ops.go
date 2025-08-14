package main

import (
    "os"
    "path/filepath"
    "fmt"
)

// Создать каталог
func createDir(dirname string) error {
    return os.Mkdir(dirname, 0755)
}

// Рекурсивный обход каталога
func walkDir(root string) {
    filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            fmt.Printf("Файл: %s (%d байт)\n", path, info.Size())
        }
        return nil
    })
}

// Демо: операции с каталогами
func runDirOperations() {
    createDir("my_dir")
    writeToFile("my_dir/file1.txt", "Файл 1")
    writeToFile("my_dir/file2.txt", "Файл 2")
    println("\nСодержимое my_dir:")
    walkDir("my_dir")
}