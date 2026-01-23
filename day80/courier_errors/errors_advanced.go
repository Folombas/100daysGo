package main

import (
	"fmt"
	"strings"
)

// AdvancedError демонстрирует создание кастомных типов ошибок
type AdvancedError struct {
	Code    int
	Message string
	Context map[string]string
}

func (e AdvancedError) Error() string {
	return fmt.Sprintf("Код ошибки %d: %s", e.Code, e.Message)
}

// AddContext добавляет контекст к ошибке
func (e AdvancedError) AddContext(key, value string) AdvancedError {
	if e.Context == nil {
		e.Context = make(map[string]string)
	}
	e.Context[key] = value
	return e
}

// CourierAdvancedError - пример ошибки конкретного домена
type CourierAdvancedError struct {
	PackageID   string
	Action      string // "scan", "load", "deliver"
	Reason      string
	Recoverable bool
}

func (e CourierAdvancedError) Error() string {
	recoverable := "неисправимая"
	if e.Recoverable {
		recoverable = "временная"
	}
	return fmt.Sprintf("Ошибка курьера (%s): %s посылки %s - %s", 
		recoverable, e.Action, e.PackageID, e.Reason)
}

// Проверка типов ошибок - важный паттерн
func handleCourierError(err error) {
	fmt.Println("\n🔧 Специализированная обработка ошибки:")
	
	// Проверяем, является ли ошибка нашим кастомным типом
	if courierErr, ok := err.(CourierAdvancedError); ok {
		fmt.Printf("   Тип: CourierAdvancedError\n")
		fmt.Printf("   Посылка: %s\n", courierErr.PackageID)
		fmt.Printf("   Действие: %s\n", courierErr.Action)
		
		if courierErr.Recoverable {
			fmt.Println("   ✅ Действие: повторить через 5 минут")
		} else {
			fmt.Println("   ❌ Действие: вернуть на склад")
		}
		return
	}
	
	// Проверяем на AdvancedError
	if advErr, ok := err.(AdvancedError); ok {
		fmt.Printf("   Тип: AdvancedError (код: %d)\n", advErr.Code)
		for k, v := range advErr.Context {
			fmt.Printf("   Контекст: %s = %s\n", k, v)
		}
		return
	}
	
	// Общая обработка через анализ текста
	if strings.Contains(err.Error(), "штрих-код") {
		fmt.Println("   💡 Рекомендация: запросить фото посылки у отправителя")
	}
	
	fmt.Println("   ℹ️  Стандартная обработка ошибки")
}

// Пример использования кастомных ошибок
func demonstrateAdvancedErrors() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("РАСШИРЕННАЯ ОБРАБОТКА ОШИБОК")
	fmt.Println(strings.Repeat("=", 50))
	
	// Создаем кастомную ошибку с контекстом
	err1 := AdvancedError{
		Code:    429,
		Message: "Превышен лимит запросов к API навигации",
	}.AddContext("Время", "12:30").AddContext("Местоположение", "ТТК, 38 км")
	
	fmt.Printf("Ошибка 1: %v\n", err1)
	handleCourierError(err1)
	
	// Ошибка специфичная для курьерской логистики
	err2 := CourierAdvancedError{
		PackageID:   "PKG-2025-123",
		Action:      "deliver",
		Reason:      "адресат недоступен",
		Recoverable: true,
	}
	
	fmt.Printf("\nОшибка 2: %v\n", err2)
	handleCourierError(err2)
	
	// Обычная ошибка, обернутая с контекстом
	baseErr := fmt.Errorf("навигатор не отвечает")
	wrappedErr := fmt.Errorf("сбой при построении маршрута в %s: %w", 
		"БЦ 'Остров'", baseErr)
	
	fmt.Printf("\nОшибка 3: %v\n", wrappedErr)
	handleCourierError(wrappedErr)
}
