package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Умножение 2x3", 2, 3, 6},
		{"Умножение 5x5", 5, 5, 25},
		{"Умножение на 0", 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestHelperFunction(t *testing.T) {
	msg := HelperFunction()
	if msg == "" {
		t.Error("HelperFunction() вернула пустую строку")
	}
}
