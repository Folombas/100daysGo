package main

// Add возвращает сумму двух чисел
func Add(a, b int) int {
	return a + b
}

// Subtract возвращает разность a - b
func Subtract(a, b int) int {
	return a - b
}

// Multiply возвращает произведение чисел 
func Multiply(a, b int) int {
	return a * b
}

// Divide возвращает результат деления a / b
// Возвращает частное и флаг успеха операции
func Divide(a, b int) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return float64(a) / float64(b), true
}

// IsPrime проверяет, является ли число простым
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}