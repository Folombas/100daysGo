package main

import "fmt"

// Employee представляет структуру сотрудника
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
}

// NewEmployee создает нового сотрудника (конструктор)
func NewEmployee(id int, firstName, lastName, position string, salary int) *Employee {
	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
	}
}

// GetFullName возвращает полное имя сотрудника
func (e Employee) GetFullName() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}

// Manager представляет менеджера с встроенной структурой Employee
type Manager struct {
	Employee
	Department string
	TeamSize   int
}
