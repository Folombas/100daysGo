package main

import (
	"fmt"
	"time"
)

// ExampleStruct для демонстрации форматирования структур
type ExampleStruct struct {
	Name     string
	Age      int
	Salary   float64
	HiredAt  time.Time
	Skills   []string
	Metadata map[string]interface{}
}

// GetFormattingExamples возвращает примеры форматирования
func GetFormattingExamples() []map[string]string {
	now := time.Now()
	person := ExampleStruct{
		Name:    "Иван Иванов",
		Age:      30,
		Salary:   12345.6789,
		HiredAt:  now.AddDate(-2, -3, -15),
		Skills:   []string{"Go", "Python", "JavaScript"},
		Metadata: map[string]interface{}{"team": "backend", "level": "senior"},
	}

	number := 42
	pi := 3.14159265359
	message := "Привет, мир!"
	enabled := true

	examples := []map[string]string{
		{
			"verb":        "%v",
			"description": "Значение в формате по умолчанию",
			"example":     fmt.Sprintf("person = %v", person),
			"result":      fmt.Sprintf("%v", person),
		},
		{
			"verb":        "%+v",
			"description": "Значение с именами полей (для структур)",
			"example":     fmt.Sprintf("person = %+v", person),
			"result":      fmt.Sprintf("%+v", person),
		},
		{
			"verb":        "%#v",
			"description": "Синтаксис Go, воспроизводящий значение",
			"example":     fmt.Sprintf("person = %#v", person),
			"result":      fmt.Sprintf("%#v", person),
		},
		{
			"verb":        "%T",
			"description": "Тип значения",
			"example":     fmt.Sprintf("person = %T", person),
			"result":      fmt.Sprintf("%T", person),
		},
		{
			"verb":        "%t",
			"description": "Булево значение",
			"example":     fmt.Sprintf("enabled = %t", enabled),
			"result":      fmt.Sprintf("%t", enabled),
		},
		{
			"verb":        "%d",
			"description": "Целое число в десятичной системе",
			"example":     fmt.Sprintf("number = %d", number),
			"result":      fmt.Sprintf("%d", number),
		},
		{
			"verb":        "%b",
			"description": "Целое число в двоичной системе",
			"example":     fmt.Sprintf("number = %b", number),
			"result":      fmt.Sprintf("%b", number),
		},
		{
			"verb":        "%x",
			"description": "Целое число в шестнадцатеричной системе",
			"example":     fmt.Sprintf("number = %x", number),
			"result":      fmt.Sprintf("%x", number),
		},
		{
			"verb":        "%f",
			"description": "Число с плавающей точкой",
			"example":     fmt.Sprintf("pi = %f", pi),
			"result":      fmt.Sprintf("%f", pi),
		},
		{
			"verb":        "%.2f",
			"description": "Число с плавающей точкой (2 знака после запятой)",
			"example":     fmt.Sprintf("pi = %.2f", pi),
			"result":      fmt.Sprintf("%.2f", pi),
		},
		{
			"verb":        "%e",
			"description": "Научная нотация (e)",
			"example":     fmt.Sprintf("pi = %e", pi),
			"result":      fmt.Sprintf("%e", pi),
		},
		{
			"verb":        "%E",
			"description": "Научная нотация (E)",
			"example":     fmt.Sprintf("pi = %E", pi),
			"result":      fmt.Sprintf("%E", pi),
		},
		{
			"verb":        "%s",
			"description": "Строка",
			"example":     fmt.Sprintf("message = %s", message),
			"result":      fmt.Sprintf("%s", message),
		},
		{
			"verb":        "%q",
			"description": "Строка в кавычках",
			"example":     fmt.Sprintf("message = %q", message),
			"result":      fmt.Sprintf("%q", message),
		},
		{
			"verb":        "%p",
			"description": "Указатель (адрес памяти)",
			"example":     fmt.Sprintf("&number = %p", &number),
			"result":      fmt.Sprintf("%p", &number),
		},
		{
			"verb":        "%15s",
			"description": "Ширина 15 символов, выравнивание по правому краю",
			"example":     fmt.Sprintf("message = '%15s'", message),
			"result":      fmt.Sprintf("'%15s'", message),
		},
		{
			"verb":        "%-15s",
			"description": "Ширина 15 символов, выравнивание по левому краю",
			"example":     fmt.Sprintf("message = '%-15s'", message),
			"result":      fmt.Sprintf("'%-15s'", message),
		},
		{
			"verb":        "%+d",
			"description": "Всегда показывать знак числа",
			"example":     fmt.Sprintf("number = %+d", number),
			"result":      fmt.Sprintf("%+d", number),
		},
		{
			"verb":        "%02d",
			"description": "Дополнение нулями до 2 символов",
			"example":     fmt.Sprintf("number = %02d", 5),
			"result":      fmt.Sprintf("%02d", 5),
		},
		{
			"verb":        "%-10.4f",
			"description": "Ширина 10, 4 знака после запятой, выравнивание влево",
			"example":     fmt.Sprintf("pi = '%-10.4f'", pi),
			"result":      fmt.Sprintf("'%-10.4f'", pi),
		},
	}

	// Добавляем категорию к каждому примеру
	for i, example := range examples {
		examples[i]["category"] = getCategory(example["verb"])
	}

	return examples
}
