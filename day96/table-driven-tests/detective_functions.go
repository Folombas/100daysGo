package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// ValidatePassword проверяет пароль на соответствие требованиям безопасности
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("пароль должен содержать минимум 8 символов")
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasUpper {
		return fmt.Errorf("пароль должен содержать хотя бы одну заглавную букву")
	}
	if !hasLower {
		return fmt.Errorf("пароль должен содержать хотя бы одну строчную букву")
	}
	if !hasDigit {
		return fmt.Errorf("пароль должен содержать хотя бы одну цифру")
	}

	return nil
}

// PasswordStrength анализирует сложность пароля
func PasswordStrength(password string) string {
	length := len(password)

	var strength int
	if length >= 12 {
		strength++
	}
	if length >= 16 {
		strength++
	}

	// Проверяем разнообразие символов
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	if hasUpper && hasLower {
		strength++
	}
	if hasDigit {
		strength++
	}
	if hasSpecial {
		strength++
	}

	switch {
	case strength >= 5:
		return "ОЧЕНЬ_СИЛЬНЫЙ"
	case strength >= 3:
		return "СИЛЬНЫЙ"
	case strength >= 2:
		return "СРЕДНИЙ"
	default:
		return "СЛАБЫЙ"
	}
}

// GenerateSecurePassword генерирует безопасный пароль (упрощенная версия)
func GenerateSecurePassword(length int) (string, error) {
	if length < 8 {
		return "", fmt.Errorf("длина пароля должна быть не менее 8 символов")
	}

	// В реальном приложении здесь была бы сложная логика генерации
	// Для демонстрации возвращаем фиктивный пароль
	return "SecurePass123!", nil
}

// ValidateEmail проверяет валидность email адреса
func ValidateEmail(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		return fmt.Errorf("ошибка проверки email: %v", err)
	}
	if !matched {
		return fmt.Errorf("неверный формат email адреса")
	}
	return nil
}

// ValidateAge проверяет корректность возраста
func ValidateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("возраст не может быть отрицательным")
	}
	if age > 150 {
		return fmt.Errorf("возраст не может превышать 150 лет")
	}
	if age < 18 {
		return fmt.Errorf("требуется возраст 18+")
	}
	return nil
}

// PrintTestResults выводит результаты тестирования в детективном стиле
func PrintTestResults(functionName string, passed, total int) {
	fmt.Printf("\n🔍 РЕЗУЛЬТАТЫ ТЕСТИРОВАНИЯ %s:\n", strings.ToUpper(functionName))
	fmt.Printf("   ✅ Пройдено: %d/%d тестов\n", passed, total)
	fmt.Printf("   📊 Успешность: %.1f%%\n", float64(passed)*100/float64(total))

	if passed == total {
		fmt.Println("   🎉 ВСЕ ТЕСТЫ ПРОЙДЕНЫ УСПЕШНО!")
	} else {
		fmt.Printf("   ⚠️  Провалено: %d тестов\n", total-passed)
	}
}

