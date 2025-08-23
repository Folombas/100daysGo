package main

import (
    "testing"
	"fmt"
)

// MockCalculator мок-объект для тестирования
type MockCalculator struct {
    AddResult      int
    SubtractResult int
    AddCalled      bool
}

func (m *MockCalculator) Add(a, b int) int {
    m.AddCalled = true
    return m.AddResult
}

func (m *MockCalculator) Subtract(a, b int) int {
    return m.SubtractResult
}

func (m *MockCalculator) Multiply(a, b int) int {
    return a * b
}

func (m *MockCalculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Тестирование с использованием мок-объекта
func TestWithMock(t *testing.T) {
    mock := &MockCalculator{
        AddResult: 100,
        SubtractResult: 50,
    }
    
    result := mock.Add(2, 3)
    if result != 100 {
        t.Errorf("Ожидалось 100, получено %d", result)
    }
    
    if !mock.AddCalled {
        t.Error("Метод Add не был вызван")
    }
}