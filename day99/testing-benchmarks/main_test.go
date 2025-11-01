package main

import "testing"

// Бенчмарк для функции fibonacci
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(30)
	}
}

// Бенчмарк для RunBenchmark с разным количеством итераций
func BenchmarkRunBenchmark1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunBenchmark(1000)
	}
}

func BenchmarkRunBenchmark10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunBenchmark(10000)
	}
}
