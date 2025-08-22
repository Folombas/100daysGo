package main

import (
	"context"
	"fmt"
)

type contextKey string

func exampleWithValue() {
	key := contextKey("userID")
	ctx := context.WithValue(context.Background(), key, 12345)

	if value := ctx.Value(key); value != nil {
		fmt.Printf("Значение из контекста: %v\n", value)
	} else {
		fmt.Println("Значение не найдено")
	}

	// Попытка доступа к несуществующему ключу
	unknownKey := contextKey("unknown")
	if value := ctx.Value(unknownKey); value == nil {
		fmt.Println("Неизвестный ключ не возвращает значений")
	}
}