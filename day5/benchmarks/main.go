package main

import "strings"

// Функция конкатенации строк через +
func concatPlus(a, b string) string {
	return a + b
}

// Функция конкатенации через strings.Builder
func concatBuilder(a, b string) string {
	var builder strings.Builder
	builder.WriteString(a)
	builder.WriteString(b)
	return builder.String()
}

// Функция сложения чисел
func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total

}
