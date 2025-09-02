package main

import (
	"testing"
)

// Бенчмарк для Fibonacci (Go)
func BenchmarkFibonacciGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateFibonacci(1000, make(chan int, 1))
	}
}

// Бенчмарк для демонстрации скорости горутин
func BenchmarkGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		for j := 0; j < 1000; j++ {
			go calculateFibonacci(100, ch)
		}
		for j := 0; j < 1000; j++ {
			<-ch
		}
	}
}

// Тест для проверки корректности вычислений
func TestFibonacci(t *testing.T) {
	ch := make(chan int)
	go calculateFibonacci(10, ch)
	result := <-ch
	
	// 10-е число Фибоначчи
	expected := 55
	if result != expected {
		t.Errorf("Ожидалось %d, получено %d", expected, result)
	}
}