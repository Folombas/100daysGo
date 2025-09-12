package internals

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func DemoScheduler() {
    fmt.Println("\n2. Работа планировщика Go")
    fmt.Println("------------------------")

    fmt.Printf("Количество CPU: %d\n", runtime.NumCPU())
    fmt.Printf("Количество горутин: %d\n", runtime.NumGoroutine())

    // Демонстрация работы планировщика
    var wg sync.WaitGroup
    start := time.Now()

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Имитация работы
            time.Sleep(1 * time.Millisecond)
            if id%100 == 0 {
                fmt.Printf("Горутина %d, активные горутины: %d\n",
                    id, runtime.NumGoroutine())
            }
        }(i)
    }

    wg.Wait()
    elapsed := time.Since(start)

    fmt.Printf("Время выполнения: %v\n", elapsed)
    fmt.Printf("Итоговое количество горутин: %d\n", runtime.NumGoroutine())

    // Настройка GOMAXPROCS
    fmt.Println("\nНастройка GOMAXPROCS:")
    fmt.Printf("Текущий GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

    // Демонстрация работы с разным количеством потоков
    testWithMaxProcs(1)
    testWithMaxProcs(2)
    testWithMaxProcs(4)
}

func testWithMaxProcs(procs int) {
    runtime.GOMAXPROCS(procs)
    start := time.Now()

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // Вычислительная задача
            for j := 0; j < 100000; j++ {
                _ = j * j
            }
        }()
    }

    wg.Wait()
    elapsed := time.Since(start)

    fmt.Printf("GOMAXPROCS=%d: %v\n", procs, elapsed)
}
