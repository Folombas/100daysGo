// calculator/calc.go
package calculator

// Add возвращает сумму двух чисел
func Add(a, b float64) float64 {
	return a + b
}

// Subtract возвращает разность двух чисел
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply возвращает произведение двух чисел
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide возвращает частное двух чисел
// Если b == 0, возвращает 0 и false
func Divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
