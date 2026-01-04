package main

import (
	"fmt"
	"reflect"
)

// TypeExplorer исследует типы через empty interface
type TypeExplorer struct{}

func (te TypeExplorer) Explore(value interface{}) {
	fmt.Printf("\n⛏️  Исследую значение:\n")
	fmt.Printf("   Тип: %T\n", value)
	fmt.Printf("   Значение: %v\n", value)
	fmt.Printf("   Reflect тип: %v\n", reflect.TypeOf(value))
	fmt.Printf("   Kind: %v\n", reflect.ValueOf(value).Kind())

	switch v := value.(type) {
	case int:
		fmt.Printf("   Type switch: целое число (%d)\n", v)
	case string:
		fmt.Printf("   Type switch: строка (%s), длина: %d\n", v, len(v))
	case []interface{}:
		fmt.Printf("   Type switch: срез с %d элементами\n", len(v))
	case map[string]interface{}:
		fmt.Printf("   Type switch: мапа с %d ключами\n", len(v))
	case float64:
		fmt.Printf("   Type switch: число с плавающей точкой (%.2f)\n", v)
	case bool:
		fmt.Printf("   Type switch: булево значение (%v)\n", v)
	default:
		fmt.Printf("   Type switch: неизвестный тип\n")
	}
}

// JSONParser демонстрирует реальное использование interface{}
func JSONParser() {
	fmt.Printf("\n📦 Парсинг JSON (реальный кейс):\n")

	jsonData := map[string]interface{}{
		"name":    "Гоша Гофер",
		"age":     38,
		"skills":  []interface{}{"Go", "Concurrency", "Backend"},
		"active":  true,
		"rating":  4.8,
		"meta": map[string]interface{}{
			"projects": 3,
			"lines":    15678,
		},
	}

	te := TypeExplorer{}
	for key, value := range jsonData {
		fmt.Printf("\n  Ключ: %s", key)
		te.Explore(value)
	}
}

// SafeExtractor безопасно извлекает значения
func SafeExtractor(data interface{}, key string) (interface{}, bool) {
	if m, ok := data.(map[string]interface{}); ok {
		val, exists := m[key]
		return val, exists
	}
	return nil, false
}

// GenericAlternative показывает альтернативу через дженерики
func GenericAlternative[T any](value T) {
	fmt.Printf("\n🎯 Дженерик альтернатива:\n")
	fmt.Printf("   Тип: %T\n", value)
	fmt.Printf("   Значение: %v\n", value)
}

func main() {
	fmt.Println("==============================================")
	fmt.Println("    🐹 EMPTY INTERFACE EXPLORER - День 62")
	fmt.Println("==============================================")
	fmt.Println("  Дата: 04.01.2026 | Тема: Empty Interface")
	fmt.Println("==============================================")

	te := TypeExplorer{}

	fmt.Printf("\n🧪 1. Разные типы через interface{}:\n")

	te.Explore(42)
	te.Explore("Вкусняшки к чаю: сухарики с изюмом")
	te.Explore(3.14159)
	te.Explore([]interface{}{"Венская сдоба", "Слойка Воскресенская", "Маффины"})
	te.Explore(true)

	JSONParser()

	fmt.Printf("\n🛡️  Безопасное извлечение:\n")

	data := map[string]interface{}{
		"task":      "Изучить empty interface",
		"priority":  1,
		"completed": false,
	}

	if val, ok := SafeExtractor(data, "task"); ok {
		if str, ok := val.(string); ok {
			fmt.Printf("   Задача: %s\n", str)
		}
	}

	fmt.Printf("\n⚖️  Сравнение подходов:\n")
	fmt.Println("   1. Empty interface:   гибко, но небезопасно")
	fmt.Println("   2. Type assertion:    безопаснее, но verbose")
	fmt.Println("   3. Generics (1.18+):  типобезопасно, современно")

	GenericAlternative("Пряники классические")
	GenericAlternative(15678)

	fmt.Printf("\n💡 Вывод:\n")
	fmt.Println("   Empty interface — мощный инструмент, но:")
	fmt.Println("   • Используй для динамических данных (JSON, YAML)")
	fmt.Println("   • Избегай в бизнес-логике")
	fmt.Println("   • Переходи на дженерики где возможно")
	fmt.Println("   • Всегда проверяй типы через type assertion")

	fmt.Println("==============================================")
	fmt.Println("   Глубина изучения: Уровень 2/10")
	fmt.Println("   Следующая тема: reflect package")
	fmt.Println("==============================================")
}
