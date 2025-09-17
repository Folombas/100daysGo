// formatter/format.go
package formatter

import "fmt"

// PrintResult выводит отформатированный результат операции на русском языке
func PrintResult(operation string, a, b, result float64) {
	fmt.Printf("🔹 %s: %.2f и %.2f = %.2f\n", operation, a, b, result)
}

// FormatWithLabel возвращает строку с меткой и значением (для веб-интерфейса)
func FormatWithLabel(label, operation string, a, b, result float64) string {
	return fmt.Sprintf("%s: %.2f %s %.2f = %.2f", label, a, operation, b, result)
}
