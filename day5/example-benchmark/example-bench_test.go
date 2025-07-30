package main

import "testing"

func BenchmarkStringConcat(b *testing.B) {
	s1, s2 := "Hello", "World"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s1 + " " + s2 // Тестируем конкатенацию
	}
}
