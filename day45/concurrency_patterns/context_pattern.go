package main

import (
	"context"
	"fmt"
	"time"
)

// Context Pattern
func ContextDemo() {
	fmt.Println("\n=== Context Pattern ===")
	
	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Запускаем длительную операцию
	resultCh := make(chan string)
	go longRunningOperation(ctx, resultCh)
	
	// Ждем результат или отмену
	select {
	case result := <-resultCh:
		fmt.Printf("Успешно: %s\n", result)
	case <-ctx.Done():
		fmt.Printf("Операция отменена: %v\n", ctx.Err())
	}
}

func longRunningOperation(ctx context.Context, resultCh chan<- string) {
	select {
	case <-time.After(5 * time.Second):
		resultCh <- "Операция завершена успешно"
	case <-ctx.Done():
		resultCh <- "Операция прервана"
	}
}