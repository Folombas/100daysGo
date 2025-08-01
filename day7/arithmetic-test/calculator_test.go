package main 

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{100, -50, 50},
	}

	for _, tc := range tests {
		got := Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 2},
		{10, 15, -5},
		{0, 0, 0},
		{-5, -3, -2},
	}

	for _, tc := range tests {
		got := Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Subtract(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 6},
		{-2, 4, -8},
		{0, 100, 0},
		{7, 0, 0},
	}

	for _, tc := range tests {
		got := Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
    t.Run("ValidDivision", func(t *testing.T) {
        got, ok := Divide(10, 2)
        if !ok {
            t.Error("Expected successful division")
        }
        if got != 5.0 {
            t.Errorf("Divide(10, 2) = %f; want 5.0", got)
        }
    })

    t.Run("DivisionByZero", func(t *testing.T) {
        _, ok := Divide(10, 0)
        if ok {
            t.Error("Expected division failure")
        }
    })

    t.Run("FractionResult", func(t *testing.T) {
        got, ok := Divide(3, 2)
        if !ok {
            t.Error("Expected successful division")
        }
        if got != 1.5 {
            t.Errorf("Divide(3, 2) = %f; want 1.5", got)
        }
    })
}

func TestIsPrime(t *testing.T) {
    primeTests := []struct {
        n    int
        want bool
    }{
        {2, true},
        {3, true},
        {4, false},
        {17, true},
        {1, false},
        {0, false},
        {-5, false},
        {997, true}, // большое простое число
    }

    for _, tt := range primeTests {
        got := IsPrime(tt.n)
        if got != tt.want {
            t.Errorf("IsPrime(%d) = %v; want %v", tt.n, got, tt.want)
        }
    }
}