package main

import (
	"testing"
)

// Тест для проверки корректности функций
func TestConcat(t *testing.T) {
	if result := concatPlus("Hello", "World"); result != "HelloWorld" {
		t.Errorf("concatPlus failed: got %s", result)
	}

	if result := concatBuilder("Hello", "World"); result != "HelloWorld" {
		t.Errorf("concatBuilder failed: got %s", result)
	}

}

// Бунчмарк конкатенации через +
func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concatPlus("Hello", "World")
	}

}

// Бенчмарк конкатенации через Builder
func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concatBuilder("Hello", "World")
	}
}

// Бенчмарк с настройкой
func BenchmarkSum100(b *testing.B) {
	benchmarkSum(100, b)
}

func Benchmark1000(b *testing.B) {
	benchmarkSum(1000, b)
}

func benchmarkSum(n int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sum(n)
	}
}

// Бенчмарк с измерениями памяти
func BenchmarkConcatMemory(b *testing.B) {
	b.ReportAllocs() //Включаем отчет по памяти

	testCases := []struct {
		a, b string
	}{
		{"a", "b"},
		{"hello", "world"},
		{"go", "lang"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_ = concatPlus(tc.a, tc.b)
		}
	}
}
