package main

import "fmt"

// IncreaseSalary увеличивает зарплату сотрудника
func (e *Employee) IncreaseSalary(percent int) {
	increase := e.Salary * percent / 100
	e.Salary += increase
	fmt.Printf("Зарплата увеличева на %d%%: +%d руб.\n", percent, increase)
}

// Promote повышает сотрудника
func (e *Employee) Promote(newPosition string) {
	oldPosition := e.Position
	e.Position = newPosition
	fmt.Printf("Повышение: %s -> %s\n", oldPosition, newPosition)
}

// DisplayInfo отображает информацию о сотруднике
func (e Employee) DisplayInfo() {
	fmt.Println("=== Информация о сотруднике ===")
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Имя: %s\n", e.GetFullName())
	fmt.Printf("Должность: %s\n", e.Position)
	fmt.Printf("Зарплата: %d руб.\n", e.Salary)
	fmt.Println("===============================")
}

func demoMethods() {
	emp := NewEmployee(4, "Екатерина", "Иванова", "Junior Developer", 80000)
	emp.DisplayInfo()

	emp.IncreaseSalary(15)
	emp.Promote("Middle Developer")
	emp.DisplayInfo()
}
