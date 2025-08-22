package main

import (
	"context"
	"fmt"
	"time"
)

func exampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Выполнено")
	case <-ctx.Done():
		fmt.Println("Таймаут:", ctx.Err())
	}
}