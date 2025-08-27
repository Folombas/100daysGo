package main

import (
	"fmt"
	//"os"
)

func main() {
	fmt.Println("📋 Демонстрация множественных возвращаемых значений в Go")
	fmt.Println("======================================================")
	fmt.Println()

	// Базовые примеры
	fmt.Println("1. Базовые примеры:")
	fmt.Println("-------------------")
	example1()
	fmt.Println()

	// Пример с ошибками
	fmt.Println("2. Пример с обработкой ошибок:")
	fmt.Println("-----------------------------")
	example2()
	fmt.Println()

	// Пример с методами
	fmt.Println("3. Пример с методами структур:")
	fmt.Println("-----------------------------")
	example3()
	fmt.Println()

	// Запуск веб-демо
	fmt.Println("4. Запуск веб-демонстрации:")
	fmt.Println("--------------------------")
	fmt.Println("Запустите отдельно: go run demo/web_demo.go")
	fmt.Println("И откройте http://localhost:8080 в браузере")
}

func example1() {
	// Простое возвращение двух значений
	a, b := swap(10, 20)
	fmt.Printf("swap(10, 20) = %d, %d\n", a, b)

	// Возвращение трех значений
	name, age, active := getUserInfo()
	fmt.Printf("Пользователь: %s, возраст: %d, активен: %t\n", name, age, active)

	// Игнорирование одного из значений
	firstName, _, _ := parseName("Иван Иванов")
	fmt.Printf("Имя: %s\n", firstName)
}

func example2() {
	// Работа с ошибками
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", result)
	}

	// Ошибка деления на ноль
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.1f\n", result)
	}
}

func example3() {
	// Создаем пользователя
	user := User{FirstName: "Петр", LastName: "Петров", Age: 25}
	
	// Метод возвращает несколько значений
	fullName, isAdult := user.GetInfo()
	fmt.Printf("Пользователь: %s, совершеннолетний: %t\n", fullName, isAdult)
	
	// Метод с ошибкой
	email, err := user.GetEmail()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Email: %s\n", email)
	}
}