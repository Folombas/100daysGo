package main

import (
	"context"
	"fmt"
	"time"
)

func exampleWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Операция отменена:", ctx.Err())
		}
	}()

	time.Sleep(10 * time.Millisecond)
	cancel()
	time.Sleep(10 * time.Millisecond)
}