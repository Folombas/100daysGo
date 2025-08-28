package utils

import (
	"fmt"
	"net/http"
)

// Divide выполняет деление с проверкой ошибок
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

// ValidateUser проверяет корректность данных пользователя
func ValidateUser(name string, age int) error {
	if name == "" {
		return fmt.Errorf("имя не может быть пустым")
	}
	if age < 0 || age > 150 {
		return fmt.Errorf("возраст должен быть от 0 до 150 лет")
	}
	if age < 18 {
		return fmt.Errorf("возраст должен быть не менее 18 лет")
	}
	return nil
}

// RecoverFromPanic восстанавливает после паники в веб-обработчиках
func RecoverFromPanic(w http.ResponseWriter, r *http.Request) {
	if rec := recover(); rec != nil {
		errorMsg := fmt.Sprintf("Внутренняя ошибка сервера: %v", rec)
		SendError(w, http.StatusInternalServerError, errorMsg)
	}
}

// SendError отправляет ошибку клиенту
func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, `<div class="error">Ошибка %d: %s</div><p><a href="/">Назад</a></p>`, statusCode, message)
}