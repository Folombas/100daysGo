package main

import (
    "fmt"
    "runtime"
    "time"
)

// 🗑️ Демонстрация работы сборщика мусора в Go
func main() {
    // Приветственное сообщение
    fmt.Println("🧹 Добро пожаловать в мир Garbage Collection в Go! 🇷🇺")

    // Выделяем много памяти с помощью среза
    fmt.Println("Создаём большое количество объектов...")
    data := make([][]byte, 0, 1000)
    for i := 0; i < 1000; i++ {
        data = append(data, make([]byte, 1024*1024)) // 1MB на объект
    }
    fmt.Println("Объекты созданы! 📦")

    // Показываем текущее состояние памяти
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Память до очистки мусора: %d MB\n", memStats.Alloc/1024/1024)

    // Удаляем ссылки на объекты, чтобы они стали 'мусором'
    data = nil
    fmt.Println("Ссылки на объекты удалены, можно вызывать сборщик мусора 🧹")

    // Запускаем сборщик мусора вручную
    runtime.GC()

    // Пауза, чтобы сборщик мусора сделал работу
    time.Sleep(2 * time.Second)

    // Проверяем сколько осталось занятой памяти
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Память после очистки мусора: %d MB\n", memStats.Alloc/1024/1024)

    fmt.Println("Спасибо за внимание! Надеюсь, теперь сборка мусора — не загадка! 🤓")
}
