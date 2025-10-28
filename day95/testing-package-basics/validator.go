package main

import (
	"strings"
	"unicode/utf8"
)

// ValidateString проверяет валидность строки
func ValidateString(s string) bool {
	// Проверяем длину строки (2-50 символов)
	length := utf8.RuneCountInString(s)
	if length < 2 || length > 50 {
		return false
	}

	// Проверяем, что строка не состоит только из пробелов
	if strings.TrimSpace(s) == "" {
		return false
	}

	// Проверяем, что первый символ - буква
	firstChar, _ := utf8.DecodeRuneInString(s)
	if !isLetter(firstChar) {
		return false
	}

	return true
}

// isLetter проверяет, является ли руна буквой
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= 'А' && r <= 'Я') || (r >= 'а' && r <= 'я')
}

// ValidateEmail проверяет валидность email (упрощенная версия)
func ValidateEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	if !strings.Contains(email, "@") {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if parts[0] == "" || parts[1] == "" {
		return false
	}

	return true
}
