package main

import "fmt"

// Employee - композиция структур (вместо наследования)
type Employee struct {
	Person  // Встроенная структура
	salary  float64
	company string
}

// NewEmployee - конструктор
func NewEmployee(name string, age int, salary float64, company string) *Employee {
	return &Employee{
		Person:  Person{name: name, age: age},
		salary:  salary,
		company: company,
	}
}

// Собственный метод
func (e Employee) WorkInfo() string {
	return fmt.Sprintf("%s работает в %s", e.name, e.company)
}

// Переопределение метода
func (e Employee) Introduce() string {
	return fmt.Sprintf("%s. Зарплата: %.2f рублей", e.Person.Introduce(), e.salary)
}
