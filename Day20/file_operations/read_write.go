package main

import (
    "os"
    "io/ioutil"
)

// Запись строки в файл
func writeToFile(filename, content string) error {
    return os.WriteFile(filename, []byte(content), 0644)
}

// Чтение файла в строку
func readFile(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    return string(data), err
}

// Демо: запись и чтение
func runReadWriteDemo() {
    writeToFile("test.txt", "Привет, файловая система!\nКириллица работает!")
    content, _ := readFile("test.txt")
    println("Содержимое файла:\n" + content)
}