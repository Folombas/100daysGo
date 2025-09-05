package errors

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// ValidationError представляет ошибку валидации
type ValidationError struct {
	Field   string
	Message string
}

// Error реализует интерфейс error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Ошибка валидации поля %s: %s", e.Field, e.Message)
}

// User представляет пользователя системы
type User struct {
	Username string
	Email    string
	Age      int
}

// Validate проверяет валидность данных пользователя
func (u *User) Validate() error {
	var errs []error
	
	if utf8.RuneCountInString(u.Username) < 3 {
		errs = append(errs, &ValidationError{
			Field:   "Username",
			Message: "должен содержать 3 символа",
		})
	}
	
	if u.Age < 0 || u.Age > 150 {
		errs = append(errs, &ValidationError{
			Field:   "Age",
			Message: "должен быть между 0 и 85",
		})
	}
	
	// Объединение ошибок
	if len(errs) > 0 {
		return fmt.Errorf("найдено %d ошибок валидации: %w", 
			len(errs), errors.Join(errs...))
	}
	
	return nil
}

// DemoValidationErrors демонстрирует валидацию ошибок
func DemoValidationErrors() {
	user := User{
		Username: "A", // Невалидное имя
		Email:    "test@example.com",
		Age:      -5, // Невалидный возраст
	}
	
	if err := user.Validate(); err != nil {
		fmt.Println("Ошибки валидации:", err)
		
		// Извлечение отдельных ошибок
		var joinedErr interface{ Unwrap() []error }
		if errors.As(err, &joinedErr) {
			for _, e := range joinedErr.Unwrap() {
				fmt.Println(" -", e)
			}
		}
	}
}