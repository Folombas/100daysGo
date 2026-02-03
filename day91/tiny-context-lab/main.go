package main

import (
    "context"
    "fmt"
    "math/rand"
    "time"
)

// fetchDataFromSlowAPI имитирует долгий API-запрос
func fetchDataFromSlowAPI(ctx context.Context, apiName string) (string, error) {
    fmt.Printf("[%s] Запрос начат...\n", apiName)
    
    // Имитируем случайную длительность запроса от 1 до 7 секунд
    delay := time.Duration(1+rand.Intn(7)) * time.Second
    
    select {
    case <-time.After(delay):
        result := fmt.Sprintf("Данные от '%s' (заняло %v)", apiName, delay)
        fmt.Printf("[%s] Успешно завершен!\n", apiName)
        return result, nil
        
    case <-ctx.Done():
        fmt.Printf("[%s] Отменено! Причина: %v\n", apiName, ctx.Err())
        return "", ctx.Err()
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("🚀 Запуск лаборатории контекстов Go!")
    fmt.Println("=====================================")
    
    // Создаем контекст с таймаутом в 3 секунды
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel() // Важно: освобождаем ресурсы
    
    fmt.Println("⏱  Установлен лимит: 3 секунды")
    fmt.Println("📞 Параллельно запускаем 3 'тяжелых' API-запроса...")
    fmt.Println()
    
    // Запускаем запросы "параллельно" (в этой упрощенной модели)
    resultChan := make(chan string, 3)
    
    // Имитируем 3 параллельных запроса
    apis := []string{"Платежный шлюз", "Гео-сервис", "База пользователей"}
    
    for _, api := range apis {
        go func(name string) {
            if data, err := fetchDataFromSlowAPI(ctx, name); err == nil {
                resultChan <- data
            } else {
                resultChan <- fmt.Sprintf("ОШИБКА '%s': %v", name, err)
            }
        }(api)
    }
    
    // Собираем результаты (не более 3 секунд из-за контекста)
    for i := 0; i < len(apis); i++ {
        select {
        case res := <-resultChan:
            fmt.Printf("✅ Результат %d: %s\n", i+1, res)
        case <-ctx.Done():
            fmt.Printf("\n⛔ Все операции отменены! Контекст истек: %v\n", ctx.Err())
            return
        }
    }
    
    fmt.Println("\n🎉 Все запросы завершены в рамках дедлайна!")
}
