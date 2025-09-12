package features

import (
    "errors"
    "fmt"
    "time"
)

// Кастомный тип ошибки
type MyError struct {
    When time.Time
    What string
}

func (e MyError) Error() string {
    return fmt.Sprintf("%v: %s", e.When.Format("2006-01-02 15:04:05"), e.What)
}

// Функция, возвращающая кастомную ошибку
func riskyOperation() error {
    return MyError{
        When: time.Now(),
        What: "произошла непредвиденная ошибка",
    }
}

func DemoErrorHandling() {
    fmt.Println("🚨 Продвинутая обработка ошибок в Go")
    fmt.Println("-----------------------------------")

    // Стандартная обработка ошибок
    if err := riskyOperation(); err != nil {
        fmt.Printf("Ошибка: %s\n", err)

        // Проверка типа ошибки
        var myErr MyError
        if errors.As(err, &myErr) {
            fmt.Printf("Детали ошибки: время=%v, сообщение=%s\n",
                myErr.When.Format("15:04:05"), myErr.What)
        }
    }

    // Обертывание ошибок
    originalErr := errors.New("оригинальная ошибка")
    wrappedErr := fmt.Errorf("дополнительный контекст: %w", originalErr)

    fmt.Printf("Обернутая ошибка: %s\n", wrappedErr)

    // Распаковка ошибок
    if unwrapped := errors.Unwrap(wrappedErr); unwrapped != nil {
        fmt.Printf("Распакованная ошибка: %s\n", unwrapped)
    }

    // Множественные ошибки
    errs := []error{
        errors.New("первая ошибка"),
        errors.New("вторая ошибка"),
        nil, // ошибки могут быть nil
    }

    fmt.Println("Обработка множественных ошибок:")
    for i, err := range errs {
        if err != nil {
            fmt.Printf("  Ошибка %d: %s\n", i+1, err)
        }
    }

    fmt.Println()
}
