package main

import (
	"testing"
	"fmt"
)

// TestAdd тестирует функцию сложения
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"положительные числа", 2, 3, 5},
		{"отрицательные числа", -2, -3, -5},
		{"смешанные числа", -5, 10, 5},
		{"ноль", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; ожидалось %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract тестирует функцию вычитания
func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; ожидалось %d", result, expected)
	}
}

// TestMultiply тестирует функцию умножения
func TestMultiply(t *testing.T) {
	result := Multiply(5, 6)
	expected := 30
	if result != expected {
		t.Errorf("Multiply(5, 6) = %d; ожидалось %d", result, expected)
	}
}

// TestDivide тестирует функцию деления
func TestDivide(t *testing.T) {
	t.Run("нормальное деление", func(t *testing.T) {
		result := Divide(15, 3)
		expected := 5
		if result != expected {
			t.Errorf("Divide(15, 3) = %d; ожидалось %d", result, expected)
		}
	})

	t.Run("деление на ноль", func(t *testing.T) {
		result := Divide(10, 0)
		if result != 0 {
			t.Errorf("Divide(10, 0) = %d; ожидалось 0", result)
		}
	})
}

// TestPower тестирует возведение в степень
func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exp      int
		expected int
	}{
		{"2^3", 2, 3, 8},
		{"5^0", 5, 0, 1},
		{"10^1", 10, 1, 10},
		{"3^2", 3, 2, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.base, tt.exp)
			if result != tt.expected {
				t.Errorf("Power(%d, %d) = %d; ожидалось %d", tt.base, tt.exp, result, tt.expected)
			}
		})
	}
}

// ExampleAdd демонстрирует использование функции Add
func ExampleAdd() {
	result := Add(3, 4)
	fmt.Println(result)
	// Output: 7
}

// BenchmarkAdd тестирует производительность функции Add
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(5, 10)
	}
}
