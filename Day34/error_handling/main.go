package main

import (
	"fmt"
	"log"

	"error_handling/examples" // Импортируем локальный пакет examples
)

func main() {
	fmt.Println("🛡️ Демонстрация обработки ошибок в Go")
	fmt.Println("=====================================")
	fmt.Println()

	// Базовые примеры
	fmt.Println("1. Базовые примеры обработки ошибок:")
	fmt.Println("-----------------------------------")
	if err := examples.BasicErrorHandling(); err != nil {
		log.Printf("Ошибка в базовых примерах: %v", err)
	}
	fmt.Println()

	// Кастомные ошибки
	fmt.Println("2. Кастомные ошибки:")
	fmt.Println("--------------------")
	if err := examples.CustomErrors(); err != nil {
		log.Printf("Ошибка в кастомных ошибках: %v", err)
	}
	fmt.Println()

	// Обертки ошибок
	fmt.Println("3. Обертки ошибок:")
	fmt.Println("------------------")
	if err := examples.ErrorWrapping(); err != nil {
		log.Printf("Ошибка в обертках ошибок: %v", err)
	}
	fmt.Println()

	// Запуск веб-сервера
	fmt.Println("4. Запуск веб-сервера с обработкой ошибок:")
	fmt.Println("------------------------------------------")
	fmt.Println("Запустите отдельно: go run handlers/web.go utils/errors.go")
	fmt.Println("И откройте http://localhost:8080 в браузере")
	fmt.Println()

	// Демонстрация паники и восстановления
	fmt.Println("5. Паника и восстановление:")
	fmt.Println("--------------------------")
	examples.PanicAndRecover()
	fmt.Println()
}