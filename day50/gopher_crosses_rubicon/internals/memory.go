package internals

import (
    "fmt"
    "runtime"
    "time"
)

func DemoMemoryManagement() {
    fmt.Println("\n1. Управление памятью в Go")
    fmt.Println("-------------------------")

    // Статистика памяти до выделения
    var m1 runtime.MemStats
    runtime.ReadMemStats(&m1)

    fmt.Printf("Выделено памяти до: %d KB\n", m1.HeapAlloc/1024)

    // Создаем большую структуру данных
    data := make([]byte, 10*1024*1024) // 10 MB
    for i := range data {
        data[i] = byte(i % 256)
    }

    // Статистика памяти после выделения
    var m2 runtime.MemStats
    runtime.ReadMemStats(&m2)

    fmt.Printf("Выделено памяти после: %d KB\n", m2.HeapAlloc/1024)
    fmt.Printf("Разница: %d KB\n", (m2.HeapAlloc-m1.HeapAlloc)/1024)

    // Принудительный вызов сборщика мусора
    runtime.GC()
    time.Sleep(100 * time.Millisecond)

    var m3 runtime.MemStats
    runtime.ReadMemStats(&m3)
    fmt.Printf("Память после GC: %d KB\n", m3.HeapAlloc/1024)

    // Анализ стека и кучи
    fmt.Println("\nРазмеры памяти:")
    fmt.Printf("  Stack: %d bytes\n", runtime.Stack(buf, false))

    // Использование указателей для контроля над памятью
    fmt.Println("\nИспользование указателей для контроля памяти:")
    withoutPointer()
    withPointer()
}

var buf = make([]byte, 1024)

func withoutPointer() {
    start := time.Now()
    var data [1000000]int64
    for i := range data {
        data[i] = int64(i)
    }
    elapsed := time.Since(start)
    fmt.Printf("Без указателей: %v, размер: %d bytes\n", elapsed, len(data)*8)
}

func withPointer() {
    start := time.Now()
    data := make([]*int64, 1000000)
    for i := range data {
        val := int64(i)
        data[i] = &val
    }
    elapsed := time.Since(start)
    fmt.Printf("С указателями: %v, размер указателей: %d bytes\n", elapsed, len(data)*8)
}
