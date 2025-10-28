package main

import "testing"

func TestValidateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"валидная строка", "Hello", true},
		{"строка на русском", "Привет", true},
		{"слишком короткая", "H", false},
		{"пустая строка", "", false},
		{"только пробелы", "   ", false},
		{"начинается с цифры", "1Hello", false},
		{"начинается с пробела", " Hello", false},
		{"максимальная длина", "Это очень длинная строка которая должна быть валидной", true},
		{"слишком длинная", "Эта строка definitely слишком длинная для нашей валидации и должна быть отвергнута", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateString(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateString(%q) = %v; ожидалось %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"валидный email", "test@example.com", true},
		{"простой email", "a@b.c", true},
		{"без @", "invalid.com", false},
		{"только локальная часть", "@example.com", false},
		{"только домен", "user@", false},
		{"пустой email", "", false},
		{"слишком длинный", "verylongemailaddress@verylongdomainnamethatmakestheentireemailtoolongandinvalid.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.email)
			if result != tt.expected {
				t.Errorf("ValidateEmail(%q) = %v; ожидалось %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		char     rune
		expected bool
	}{
		{'a', true},
		{'Z', true},
		{'А', true},
		{'я', true},
		{'1', false},
		{' ', false},
		{'@', false},
	}

	for _, tt := range tests {
		t.Run(string(tt.char), func(t *testing.T) {
			result := isLetter(tt.char)
			if result != tt.expected {
				t.Errorf("isLetter(%q) = %v; ожидалось %v", tt.char, result, tt.expected)
			}
		})
	}
}
