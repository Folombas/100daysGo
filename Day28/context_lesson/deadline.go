package main

import (
	"context"
	"fmt"
	"time"
)

func exampleWithDeadline() {
	deadline := time.Now().Add(30 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Выполнено")
	case <-ctx.Done():
		fmt.Println("Дедлайн превышен:", ctx.Err())
	}
}