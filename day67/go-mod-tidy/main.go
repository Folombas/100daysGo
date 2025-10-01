package main

import (
	"fmt"

	"golang.org/x/text/language"
	"github.com/pkg/errors"
)

func main() {
	// Используем только text и errors
	fmt.Println("Демонстрация используемых зависимостей:")

	// golang.org/x/text
	tag := language.MustParse("ru-RU")
	fmt.Printf("Язык: %s\n", tag)

	// github.com/pkg/errors
	err := errors.New("демо ошибка")
	fmt.Printf("Ошибка: %v\n", err)

	fmt.Println("✅ Только эти зависимости должны остаться после go mod tidy")
}
