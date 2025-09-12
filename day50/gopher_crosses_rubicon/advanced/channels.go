package advanced

import (
    "fmt"
    "time"
)

func DemoChannels() {
    fmt.Println("\n1. Продвинутая работа с каналами")
    fmt.Println("-------------------------------")

    // Каналы с буферизацией
    buffered := make(chan string, 3)
    buffered <- "первое"
    buffered <- "второе"
    buffered <- "третье"

    fmt.Println("Буферизированный канал:")
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)

    // Селекты и таймауты
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "сообщение из ch1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "сообщение из ch2"
    }()

    fmt.Println("\nМультиплексирование каналов:")
    for i := 0; i < 2; i++ {
        select {
        case msg := <-ch1:
            fmt.Println("Получено:", msg)
        case msg := <-ch2:
            fmt.Println("Получено:", msg)
        case <-time.After(3 * time.Second):
            fmt.Println("Таймаут")
        }
    }

    // Закрытие каналов и range
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Printf("Обработана задача %d\n", j)
            } else {
                fmt.Println("Все задачи обработаны")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
    }
    close(jobs)

    <-done
}
