package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// 🎯 1. БАЗОВЫЕ ОШИБКИ
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrAccessDenied = errors.New("access denied")
)

// 🎯 2. КАСТОМНЫЕ ТИПЫ ОШИБОК
type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s (value: %v)",
		e.Field, e.Message, e.Value)
}

type NetworkError struct {
	URL        string
	StatusCode int
	RetryAfter time.Duration
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error [%d] %s - retry after %v",
		e.StatusCode, e.URL, e.RetryAfter)
}

func (e NetworkError) Timeout() bool {
	return e.StatusCode == 408 || e.StatusCode == 429
}

// 🎯 3. ФУНКЦИИ С ВОЗВРАТОМ ОШИБОК
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func parsePositiveNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parsePositiveNumber: %w", err)
	}
	if num <= 0 {
		return 0, ValidationError{
			Field:   "number",
			Value:   num,
			Message: "must be positive",
		}
	}
	return num, nil
}

func readConfig(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("readConfig: %w: %s", ErrNotFound, filename)
	}
	return nil
}

// 🎯 4. ОБЕРТЫВАНИЕ ОШИБОК
func processUserInput(input map[string]string) error {
	ageStr, ok := input["age"]
	if !ok {
		return fmt.Errorf("processUserInput: %w: missing age", ErrInvalidInput)
	}

	age, err := parsePositiveNumber(ageStr)
	if err != nil {
		return fmt.Errorf("processUserInput: %w", err)
	}

	if age < 18 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "must be at least 18",
		}
	}

	return nil
}

// 🎯 5. МНОЖЕСТВЕННЫЕ ОШИБКИ
type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	return fmt.Sprintf("%d errors occurred: %v", len(e.Errors), e.Errors)
}

func validateUser(user map[string]string) error {
	var errs []error

	// Validate name
	if name, ok := user["name"]; !ok || name == "" {
		errs = append(errs, ValidationError{
			Field:   "name",
			Message: "is required",
		})
	}

	// Validate age
	if ageStr, ok := user["age"]; ok {
		if _, err := parsePositiveNumber(ageStr); err != nil {
			errs = append(errs, err)
		}
	} else {
		errs = append(errs, ValidationError{
			Field:   "age",
			Message: "is required",
		})
	}

	// Validate email
	if email, ok := user["email"]; ok && len(email) > 0 {
		if !contains(email, "@") {
			errs = append(errs, ValidationError{
				Field:   "email",
				Value:   email,
				Message: "must contain @ symbol",
			})
		}
	}

	if len(errs) > 0 {
		return MultiError{Errors: errs}
	}
	return nil
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// 🎯 6. PANIC И RECOVER
func safeExecute(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	fn()
	return nil
}

func riskyOperation() {
	fmt.Println("🚀 Performing risky operation...")
	panic("something went terribly wrong!")
}

// 🎯 7. УТИЛИТЫ ДЛЯ РАБОТЫ С ОШИБКАМИ
func demonstrateErrorUtilities() {
	fmt.Println("\n🔧 УТИЛИТЫ ДЛЯ РАБОТЫ С ОШИБКАМИ:")

	// errors.Is()
	err := readConfig("missing_file.conf")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("✅ errors.Is() correctly identified ErrNotFound")
	}

	// errors.As()
	validationErr := ValidationError{Field: "test", Message: "test error"}
	wrappedErr := fmt.Errorf("context: %w", validationErr)

	var valErr ValidationError
	if errors.As(wrappedErr, &valErr) {
		fmt.Printf("✅ errors.As() extracted: %v\n", valErr)
	}

	// errors.Unwrap()
	if unwrapped := errors.Unwrap(wrappedErr); unwrapped != nil {
		fmt.Printf("✅ errors.Unwrap() found: %v\n", unwrapped)
	}
}

func main() {
	fmt.Println("🎯 ERROR VALUES IN GO - DEMONSTRATION")
	fmt.Println("=====================================")

	// 🎯 1. Базовые ошибки
	fmt.Println("\n1. 📝 БАЗОВЫЕ ОШИБКИ:")
	if result, err := divide(10, 0); err != nil {
		fmt.Printf("   ❌ Division error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Result: %.2f\n", result)
	}

	// 🎯 2. Кастомные типы ошибок
	fmt.Println("\n2. 🏗️ КАСТОМНЫЕ ТИПЫ ОШИБОК:")
	networkErr := NetworkError{
		URL:        "https://api.example.com",
		StatusCode: 429,
		RetryAfter: 30 * time.Second,
	}
	fmt.Printf("   🌐 Network error: %v\n", networkErr)
	fmt.Printf("   ⏰ Is timeout: %t\n", networkErr.Timeout())

	// 🎯 3. Обертывание ошибок
	fmt.Println("\n3. 🔄 ОБЕРТЫВАНИЕ ОШИБОК:")
	userInput := map[string]string{"age": "invalid"}
	if err := processUserInput(userInput); err != nil {
		fmt.Printf("   ❌ Processing error: %v\n", err)

		// Демонстрация извлечения оригинальной ошибки
		var numErr *strconv.NumError
		if errors.As(err, &numErr) {
			fmt.Printf("   🔍 Extracted NumError: %v\n", numErr)
		}
	}

	// 🎯 4. Множественные ошибки
	fmt.Println("\n4. 📊 МНОЖЕСТВЕННЫЕ ОШИБКИ:")
	invalidUser := map[string]string{
		"name":  "",
		"age":   "-5",
		"email": "invalid-email",
	}
	if err := validateUser(invalidUser); err != nil {
		fmt.Printf("   ❌ Validation failed: %v\n", err)

		var multiErr MultiError
		if errors.As(err, &multiErr) {
			fmt.Printf("   📈 Found %d validation errors:\n", len(multiErr.Errors))
			for i, e := range multiErr.Errors {
				fmt.Printf("      %d. %v\n", i+1, e)
			}
		}
	}

	// 🎯 5. Panic и Recover
	fmt.Println("\n5. 🛡️ PANIC И RECOVER:")
	if err := safeExecute(riskyOperation); err != nil {
		fmt.Printf("   ✅ Safely handled panic: %v\n", err)
	}

	// 🎯 6. Утилиты для работы с ошибками
	demonstrateErrorUtilities()

	// 🎯 7. ЛУЧШИЕ ПРАКТИКИ
	fmt.Println("\n💡 ЛУЧШИЕ ПРАКТИКИ РАБОТЫ С ОШИБКАМИ:")
	bestPractices := []string{
		"✅ Всегда проверяйте возвращаемые ошибки",
		"✅ Создавайте информативные сообщения об ошибках",
		"✅ Используйте errors.Is() для проверки конкретных ошибок",
		"✅ Используйте errors.As() для извлечения кастомных типов",
		"✅ Обертывайте ошибки с %w для сохранения контекста",
		"✅ Создавайте кастомные типы для сложных сценариев",
		"✅ Используйте defer/recover для обработки паник",
		"❌ Никогда не игнорируйте ошибки с _",
		"❌ Избегайте паник в обычном потоке выполнения",
		"❌ Не используйте строки для проверки ошибок",
	}

	for _, practice := range bestPractices {
		fmt.Println("   " + practice)
	}

	fmt.Println("\n🎉 Демонстрация завершена! Освоены ключевые аспекты работы с ошибками в Go.")
}
