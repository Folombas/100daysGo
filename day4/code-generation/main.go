package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Саша"
	code := fmt.Sprintf(`package main

import "fmt"

func Hello() {
	fmt.Println("Привет, %s!")
}`, name)

	os.WriteFile("hello.go", []byte(code), 0644)
	fmt.Println("Файл hello.go создан!")
}
