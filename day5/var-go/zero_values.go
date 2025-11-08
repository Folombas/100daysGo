package main

import "fmt"

func demoZeroValues() {
	fmt.Println("🔮 НУЛЕВЫЕ ЗНАЧЕНИЯ:")
	fmt.Println("────────────────────")

	// Группируем связанные переменные
	var (
		intValue     int
		floatValue   float64
		stringValue  string
		boolValue    bool
		pointerValue *int
		sliceValue   []int
		mapValue     map[string]int
	)

	// Компактный вывод
	fmt.Printf(`
   int:        %d
   float64:    %.1f
   string:     %q
   bool:       %t
   *int:       %v
   []int:      %v (len=%d)
   map:        %v
`, intValue, floatValue, stringValue, boolValue, pointerValue,
   sliceValue, len(sliceValue), mapValue)

	// Более идиоматические проверки
	if stringValue == "" {
		fmt.Println("✅ stringValue == \"\" (нулевое значение строки)")
	}

	if sliceValue == nil {
		fmt.Println("✅ sliceValue == nil (нулевое значение среза)")
	}

	if mapValue == nil {
		fmt.Println("✅ mapValue == nil (нулевое значение карты)")
	}
	fmt.Println()
}

func demoTypeInference() {
	fmt.Println("🎭 ВЫВОД ТИПОВ:")
	fmt.Println("───────────────")

	// Более наглядное сравнение
	explicitType := 42     // int
	inferredType := 42     // int
	shortDecl := 3.14      // float64
	number := 42           // int
	decimal := 42.0        // float64
	text := "42"           // string

	// Группированный вывод с выравниванием
	fmt.Printf(`
   explicitType: %-8T = %v
   inferredType: %-8T = %v
   shortDecl:    %-8T = %.2f
   number:       %-8T = %v
   decimal:      %-8T = %v
   text:         %-8T = %q
`, explicitType, explicitType,
   inferredType, inferredType,
   shortDecl, shortDecl,
   number, number,
   decimal, decimal,
   text, text)

	fmt.Println()
}

// Дополнительная функция для демонстрации лучших практик
func demoVariableGroups() {
	fmt.Println("📊 ГРУППИРОВКА ПЕРЕМЕННЫХ:")
	fmt.Println("──────────────────────────")

	// Логически связанные переменные группируем
	var (
		userID      = 1001
		userName    = "john_doe"
		isActive    = true
		lastLogin   = "2024-01-15"
	)

	// Конфигурационные параметры
	var (
		maxConnections = 100
		timeout        = 30
		debugEnabled   = false
	)

	fmt.Printf(`
   Пользователь:
      ID: %d, Имя: %s, Активен: %t, Последний вход: %s

   Конфигурация:
      Макс. подключений: %d, Таймаут: %dс, Отладка: %t
`, userID, userName, isActive, lastLogin,
   maxConnections, timeout, debugEnabled)
}


