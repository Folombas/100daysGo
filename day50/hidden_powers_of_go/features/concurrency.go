package features

import (
    "fmt"
    "sync"
    "time"
)

func DemoConcurrency() {
    fmt.Println("🚀 Продвинутая конкурентность в Go")
    fmt.Println("----------------------------------")

    // WaitGroup для ожидания завершения горутин
    var wg sync.WaitGroup
    results := make(chan string, 5)

    // Запускаем несколько горутин с разными задачами
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            time.Sleep(time.Duration(id*100) * time.Millisecond)
            results <- fmt.Sprintf("Горутина %d завершена", id)
        }(i)
    }

    // Отслеживаем завершение в отдельной горутине
    go func() {
        wg.Wait()
        close(results)
    }()

    // Читаем результаты
    fmt.Println("Результаты выполнения горутин:")
    for result := range results {
        fmt.Println("  -", result)
    }

    // Мьютексы для безопасного доступа к общим данным
    var counter int
    var mu sync.Mutex

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Printf("Безопасный счетчик: %d\n\n", counter)
}
