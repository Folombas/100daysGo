package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// processText обрабатывает текст в соответствии с выбранной операцией
func processText(text string, operation string) string {
	switch operation {
	case "upper":
		return strings.ToUpper(text)
	case "lower":
		return strings.ToLower(text)
	case "title":
		return strings.Title(text)
	case "reverse":
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	default:
		return text
	}
}

// StartCLI запускает CLI инструмент
func StartCLI() {
	textPtr := flag.String("text", "", "Текст для обработки")
	opPtr := flag.String("op", "upper", "Операция: upper, lower, title, reverse")
	helpPtr := flag.Bool("help", false, "Показать справку")

	flag.Parse()

	if *helpPtr {
		fmt.Println("🐚 GoCLI - Простой CLI инструмент на Go")
		fmt.Println("Использование: gocli --text=\"ваш текст\" --op=operation")
		fmt.Println("Доступные операции: upper, lower, title, reverse")
		return
	}

	if *textPtr == "" {
		fmt.Println("❌ Ошибка: Не указан текст для обработки")
		fmt.Println("Используйте --help для справки")
		os.Exit(1)
	}

	result := processText(*textPtr, *opPtr)
	fmt.Printf("📝 Результат: %s\n", result)
}