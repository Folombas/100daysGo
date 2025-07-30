package examplebenchmark_test

import "testing"

func example_bench(b *testing.B) {
	// Подготовка данных (не учитывается в замерах)

	b.ResetTimer() // Сброс таймера
	for i := 0; i < b.N; i++ {
		// Тестируемый код
	}
}
