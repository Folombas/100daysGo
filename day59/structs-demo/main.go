package main

import (
	"fmt"
)

func main() {
	fmt.Println("🎯 День 59: Структуры в Go - Фундамент данных")
	fmt.Println("============================================")
	fmt.Println()

	// Демонстрация базовых структур
	fmt.Println("1. Базовые структуры:")
	demoBasicStructs()

	fmt.Println("\n2. Методы структур:")
	demoMethods()

	fmt.Println("\n3. Интерфейсы и структуры:")
	demoInterfaces()

	fmt.Println("\n4. Встроенные структуры:")
	demoEmbeddedStructs()

	fmt.Println("\n✨ Демонстрация завершена! Изучайте структуры - это фундамент Go!")
}

func demoBasicStructs() {
	// Создаем экземпляр структуры Employee
	emp := Employee{
		ID:        1,
		FirstName: "Иван",
		LastName:  "Петров",
		Position:  "Разработчик Go",
		Salary:    150000,
	}

	fmt.Printf("Сотрудник: %s %s\n", emp.FirstName, emp.LastName)
	fmt.Printf("Должность: %s\n", emp.Position)
	fmt.Printf("Зарплата: %d руб.\n", emp.Salary)

	// Создаем через указатель
	emp2 := &Employee{
		ID:        2,
		FirstName: "Мария",
		LastName:  "Сидорова",
		Position:  "Team Lead",
		Salary:    250000,
	}

	fmt.Printf("Сотрудник 2: %s %s\n", emp2.FirstName, emp2.LastName)
}

func demoEmbeddedStructs() {
	// Встроенная структура
	manager := Manager{
		Employee: Employee{
			ID:        3,
			FirstName: "Алексей",
			LastName:  "Кузнецов",
			Position:  "Менеджер проекта",
			Salary:    300000,
		},
		Department: "Разработка",
		TeamSize:   10,
	}

	fmt.Printf("Менеджер: %s %s\n", manager.FirstName, manager.LastName)
	fmt.Printf("Департамент: %s\n", manager.Department)
	fmt.Printf("Размер команды: %d\n", manager.TeamSize)
}
