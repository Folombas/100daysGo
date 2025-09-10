package scope

import "fmt"

// DemoBlockLevelScope демонстрирует область видимости на уровне блоков
func DemoBlockLevelScope() {
	// Переменная уровня функции
	outerVar := "Я внешняя переменная"
	fmt.Println("🏠 Внешняя переменная:", outerVar)
	
	// Блок 1: if statement
	if condition := true; condition {
		innerVar := "Я переменная блока if"
		fmt.Println("   📦 Переменная блока if:", innerVar)
		fmt.Println("   🏠 Доступ к внешней переменной из if:", outerVar)
		
		// Переменная condition доступна только внутри блока if
		fmt.Println("   📦 Переменная condition:", condition)
	}
	
	// Блок 2: for loop
	for i := 0; i < 2; i++ {
		loopVar := "Я переменная цикла"
		fmt.Printf("   🔁 Итерация %d: %s\n", i, loopVar)
		fmt.Printf("   🏠 Доступ к внешней переменной из цикла: %s\n", outerVar)
		
		// Переменная i доступна только внутри цикла
		fmt.Printf("   🔁 Счетчик цикла i: %d\n", i)
	}
	
	// Блок 3: switch statement
	switch value := 42; value {
	case 42:
		caseVar := "Я переменная case"
		fmt.Println("   🔀 Переменная case:", caseVar)
		fmt.Println("   🏠 Доступ к внешней переменной из case:", outerVar)
	default:
		fmt.Println("Другой случай")
	}
	
	// Ошибка! Все эти переменные не доступны вне своих блоков
	// fmt.Println(innerVar)  // Недоступна
	// fmt.Println(loopVar)   // Недоступна  
	// fmt.Println(caseVar)   // Недоступна
	// fmt.Println(i)         // Недоступна
	// fmt.Println(condition) // Недоступна
	// fmt.Println(value)     // Недоступна
}