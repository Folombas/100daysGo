package errors

import (
	"fmt"
)

// CustomError - пользовательская ошибка с дополнительными полями
type CustomError struct {
	Code    int
	Message  string
	Datails  string
}

// Error реализует интерфейс error
func (e *CustomError) Error() string {
	return fmt.Sprintf("Ошибка %d: %s (Детали: %s)", e.Code, e.Message, e.Datails)
}

// NewCustomError - создаёт новую пользовательскую ошибку
func NewCustomError(code int, message, details string) error {
	return &CustomError{
		Code:    code,
		Message:  message,
		Datails:  details,
	}
}

// DemoCustomErrors демонстрирует использование пользовательских ошибок
func DemoCustomErrors() {
	err := NewCustomError(404, "Ресурс не найден", "Файл config.txt отсутствует")
	fmt.Println("Пользовательская ошибка:", err)

	// Проверка типа ошибки
	if customErr, ok := err.(*CustomError); ok {
		fmt.Printf("Код ошибки: %d, Сообщение: %s\n", customErr.Code, customErr.Message)
	}
}