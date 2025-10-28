package main

// Add возвращает сумму двух чисел
func Add(a, b int) int {
	return a + b
}

// Subtract возвращает разность двух чисел
func Subtract(a, b int) int {
	return a - b
}

// Multiply возвращает произведение двух чисел
func Multiply(a, b int) int {
	return a * b
}

// Divide возвращает результат деления a на b
// Если b == 0, возвращает 0 и false в будущих расширениях
func Divide(a, b int) int {
	if b == 0 {
		return 0 // В реальном проекте здесь должно быть другое поведение
	}
	return a / b
}

// Power возвращает a в степени b
func Power(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
