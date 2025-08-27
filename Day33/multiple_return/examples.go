package main

import (
	"errors"
	//"fmt"
	//"strconv"
)

// 1. Функция, возвращающая два значения
func swap(a, b int) (int, int) {
	return b, a
}

// 2. Функция, возвращающая три значения
func getUserInfo() (string, int, bool) {
	return "Анна", 30, true
}

// 3. Функция, возвращающая значения с ошибкой
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

// 4. Функция с именованными возвращаемыми значениями
func parseName(fullName string) (firstName, lastName string, err error) {
	// Простая логика разбиения имени
	names := splitName(fullName)
	if len(names) < 2 {
		return "", "", errors.New("неполное имя")
	}
	firstName = names[0]
	lastName = names[1]
	return // возвращаем именованные значения
}

func splitName(name string) []string {
	// Простая функция разделения имени
	var parts []string
	current := ""
	for _, char := range name {
		if char == ' ' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}

// 5. Структура и методы с множественными возвращаемыми значениями
type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Метод возвращает два значения
func (u User) GetInfo() (string, bool) {
	fullName := u.FirstName + " " + u.LastName
	isAdult := u.Age >= 18
	return fullName, isAdult
}

// Метод возвращает значение и ошибку
func (u User) GetEmail() (string, error) {
	if u.Email == "" {
		return "", errors.New("email не установлен")
	}
	return u.Email, nil
}

// 6. Функция с переменным числом аргументов и множественным возвратом
func calculateStats(numbers ...int) (min, max, sum, avg int) {
	if len(numbers) == 0 {
		return 0, 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	sum = 0
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}
	
	avg = sum / len(numbers)
	return
}

// 7. Функция, возвращающая функцию и ошибку
func createMultiplier(factor int) (func(int) int, error) {
	if factor == 0 {
		return nil, errors.New("множитель не может быть нулём")
	}
	
	multiplier := func(x int) int {
		return x * factor
	}
	
	return multiplier, nil
}