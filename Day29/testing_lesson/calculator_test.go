package main

import (
    "testing"
)

// Тестирование функции Add
func TestCalculator_Add(t *testing.T) {
    calc := Calculator{}
    
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Позитивные числа", 2, 3, 5},
        {"Отрицательные числа", -2, -3, -5},
        {"Смешанные числа", -2, 3, 1},
        {"Нули", 0, 0, 0},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := calc.Add(test.a, test.b)
            if result != test.expected {
                t.Errorf("Add(%d, %d) = %d; ожидалось %d", 
                    test.a, test.b, result, test.expected)
            }
        })
    }
}

// Тестирование функции Divide с обработкой ошибок
func TestCalculator_Divide(t *testing.T) {
    calc := Calculator{}
    
    t.Run("Успешное деление", func(t *testing.T) {
        result, err := calc.Divide(10, 2)
        if err != nil {
            t.Errorf("Неожиданная ошибка: %v", err)
        }
        if result != 5 {
            t.Errorf("Ожидалось 5, получено %d", result)
        }
    })
    
    t.Run("Деление на ноль", func(t *testing.T) {
        _, err := calc.Divide(10, 0)
        if err == nil {
            t.Error("Ожидалась ошибка при делении на ноль")
        }
    })
}