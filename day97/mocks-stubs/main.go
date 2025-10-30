package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("🔄 СИМУЛЯТОР БИЗНЕС-ЛОГИКИ С МОКАМИ И СТАБАМИ")
    fmt.Println("=============================================")
    fmt.Println()

    // Демонстрация работы с реальными сервисами
    fmt.Println("🎯 ДЕМО: РАБОТА С РЕАЛЬНЫМИ СЕРВИСАМИ")
    realPaymentService := NewRealPaymentService("real_api_key")
    realUserRepo := NewRealUserRepository()

    realService := NewOrderService(realPaymentService, realUserRepo)

    paymentID, err := realService.ProcessOrder("user_123", 100.0)
    if err != nil {
        log.Printf("Ошибка обработки заказа: %v", err)
    } else {
        fmt.Printf("✅ Заказ обработан. ID платежа: %s\n", paymentID)
    }

    fmt.Println()

    // Демонстрация работы со стабами
    fmt.Println("🎯 ДЕМО: РАБОТА СО СТАБАМИ")
    paymentStub := &StubPaymentService{
        ProcessPaymentFunc: func(amount float64, currency string) (string, error) {
            return "stub_payment_456", nil
        },
    }

    userRepoStub := &StubUserRepository{
        FindByIDFunc: func(id string) (*User, error) {
            return &User{
                ID:      id,
                Name:    "Гоша (стаб)",
                Email:   "gosha-stub@example.com",
                Balance: 500.0,
            }, nil
        },
    }

    stubService := NewOrderService(paymentStub, userRepoStub)

    stubPaymentID, err := stubService.ProcessOrder("user_456", 50.0)
    if err != nil {
        log.Printf("Ошибка обработки заказа (стаб): %v", err)
    } else {
        fmt.Printf("✅ Заказ обработан через стаб. ID платежа: %s\n", stubPaymentID)
    }

    fmt.Println()
    fmt.Println("💡 ЗАПУСК ТЕСТОВ:")
    fmt.Println("go test -v")
    fmt.Println()
    fmt.Println("📚 ЧТО ИЗУЧИЛИ:")
    fmt.Println("• Создание интерфейсов для абстракции")
    fmt.Println("• Реализация моков с testify/mock")
    fmt.Println("• Создание стабов с кастомной логикой")
    fmt.Println("• Написание unit-тестов с изоляцией зависимостей")
    fmt.Println("• Table-driven tests с моками")
    fmt.Println()
    fmt.Println("🚀 Каждый написанный тест с моками — это шаг к профессиональному Go-разработчику!")
}
