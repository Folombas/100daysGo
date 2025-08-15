package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// Структура для демонстрационных данных
type Data struct {
	Name    string   `json:"name" xml:"name"`
	Age     int      `json:"age" xml:"age"`
	Email   string   `json:"email,omitempty" xml:"email,omitempty"`
	Hobbies []string `json:"hobbies" xml:"hobbies>hobby"`
}

// Конвертация в JSON
func convertToJSON(input string) (string, error) {
	data := parseInput(input)
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// Конвертация в XML
func convertToXML(input string) (string, error) {
	data := parseInput(input)
	result, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(result), nil
}

// Парсинг ввода пользователя
func parseInput(input string) Data {
	parts := strings.Split(input, ",")
	if len(parts) < 3 {
		return Data{}
	}
	
	hobbies := strings.Split(parts[3], ";")
	return Data{
		Name:    strings.TrimSpace(parts[0]),
		Age:     parseInt(parts[1]),
		Email:   strings.TrimSpace(parts[2]),
		Hobbies: hobbies,
	}
}

// Безопасный парсинг чисел
func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}