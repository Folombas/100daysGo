package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Обработчик главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// Обработчик определения типа данных
func detectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	
	// Получаем значение из формы
	input := strings.TrimSpace(r.FormValue("input"))
	
	// Определяем тип данных
	dataType, details := detectType(input)
	
	// Подготавливаем данные для шаблона
	data := struct {
		Input   string
		Type    string
		Details string
	}{
		Input:   input,
		Type:    dataType,
		Details: details,
	}
	
	// Рендерим результат
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}

// Функция для определения типа данных
func detectType(input string) (string, string) {
	if input == "" {
		return "Пустая строка", "Вы ничего не ввели"
	}
	
	// Проверяем на булево значение
	if strings.ToLower(input) == "true" || strings.ToLower(input) == "false" {
		return "Boolean", "Логическое значение (true/false)"
	}
	
	// Проверяем на целое число
	if intValue, err := strconv.Atoi(input); err == nil {
		return "Integer", fmt.Sprintf("Целое число: %d (размер: %d бит)", intValue, reflect.TypeOf(intValue).Bits())
	}
	
	// Проверяем на число с плавающей точкой
	if floatValue, err := strconv.ParseFloat(input, 64); err == nil {
		return "Float", fmt.Sprintf("Число с плавающей точкой: %f (размер: 64 бита)", floatValue)
	}
	
	// Проверяем на комплексное число
	if strings.Contains(input, "i") {
		if complexValue, err := strconv.ParseComplex(input, 128); err == nil {
			return "Complex", fmt.Sprintf("Комплексное число: %v (размер: 128 бит)", complexValue)
		}
	}
	
	// Проверяем на строку
	if len(input) > 0 {
		details := fmt.Sprintf("Длина строки: %d символов\n", len(input))
		details += fmt.Sprintf("Первый символ: %c (код: %d)\n", input[0], input[0])
		details += fmt.Sprintf("Последний символ: %c (код: %d)", input[len(input)-1], input[len(input)-1])
		return "String", details
	}
	
	return "Неизвестный тип", "Не удалось определить тип данных"
}