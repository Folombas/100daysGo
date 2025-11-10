package main

import "fmt"

// Константы времени суток
const (
    Morning = iota  // 0
    Day             // 1
    Evening         // 2
    Night           // 3
)

// Константы статусов заказа с использованием iota и сдвига
const (
    StatusPending = iota + 1  // 1
    StatusProcessing          // 2
    StatusShipped             // 3
    StatusDelivered           // 4
    StatusCancelled           // 5
)

// Константы с битовыми масками
const (
    ReadPermission = 1 << iota  // 1 << 0 = 1
    WritePermission             // 1 << 1 = 2
    ExecutePermission           // 1 << 2 = 4
    AdminPermission             // 1 << 3 = 8
)

// Константы размеров в байтах с вычислениями
const (
    _  = iota
    KB = 1 << (10 * iota)  // 1 << 10 = 1024
    MB                     // 1 << 20 = 1048576
    GB                     // 1 << 30 = 1073741824
    TB                     // 1 << 40 = 1099511627776
)

// Константы с пропуском значений
const (
    Monday = iota * 2    // 0
    Tuesday              // 2
    Wednesday            // 4
    _                    // пропускаем 6
    Friday               // 8
)

func main() {
    fmt.Println("🎯 Day 7: Const and Iota - Мощь перечислений в Go!")
    fmt.Println("==================================================")

    // Демонстрация базового использования iota
    fmt.Println("\n📅 Время суток:")
    fmt.Printf("Morning: %d\n", Morning)
    fmt.Printf("Day: %d\n", Day)
    fmt.Printf("Evening: %d\n", Evening)
    fmt.Printf("Night: %d\n", Night)

    // Статусы заказа
    fmt.Println("\n📦 Статусы заказа:")
    fmt.Printf("Pending: %d\n", StatusPending)
    fmt.Printf("Processing: %d\n", StatusProcessing)
    fmt.Printf("Delivered: %d\n", StatusDelivered)

    // Битовые маски разрешений
    fmt.Println("\n🔐 Система разрешений:")
    userPermissions := ReadPermission | WritePermission
    fmt.Printf("Read: %b (%d)\n", ReadPermission, ReadPermission)
    fmt.Printf("Write: %b (%d)\n", WritePermission, WritePermission)
    fmt.Printf("User permissions: %b (%d)\n", userPermissions, userPermissions)

    // Размеры файлов
    fmt.Println("\n💾 Размеры файлов:")
    fmt.Printf("KB: %d bytes\n", KB)
    fmt.Printf("MB: %d bytes\n", MB)
    fmt.Printf("GB: %d bytes\n", GB)

    // Дни недели с пропусками
    fmt.Println("\n📆 Дни недели (с пропусками):")
    fmt.Printf("Monday: %d\n", Monday)
    fmt.Printf("Tuesday: %d\n", Tuesday)
    fmt.Printf("Friday: %d\n", Friday)

    // Практический пример использования
    fmt.Println("\n💡 Практический пример:")
    processOrder(StatusProcessing)
    checkPermissions(ReadPermission | WritePermission)

    fmt.Println("\n🎉 Вывод: Iota - это мощный инструмент для создания")
    fmt.Println("   последовательных констант и перечислений в Go!")
}

func processOrder(status int) {
    switch status {
    case StatusPending:
        fmt.Println("Заказ ожидает обработки...")
    case StatusProcessing:
        fmt.Println("Заказ в процессе обработки!")
    case StatusShipped:
        fmt.Println("Заказ отправлен!")
    case StatusDelivered:
        fmt.Println("Заказ доставлен!")
    default:
        fmt.Println("Неизвестный статус заказа")
    }
}

func checkPermissions(perms int) {
    if perms&ReadPermission != 0 {
        fmt.Println("✅ Есть право на чтение")
    }
    if perms&WritePermission != 0 {
        fmt.Println("✅ Есть право на запись")
    }
    if perms&AdminPermission != 0 {
        fmt.Println("✅ Есть админские права")
    }
}
