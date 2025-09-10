package scope

import "fmt"

// DemoFunctionLevelScope демонстрирует область видимости на уровне функций
func DemoFunctionLevelScope() {
	functionVariable := "Я переменная функции DemoFunctionLevelScope"
	fmt.Println("🏠 Локальная переменная функции:", functionVariable)
	
	// Вложенный блок (if)
	if true {
		blockVariable := "Я переменная блока if"
		fmt.Println("   📦 Переменная блока if:", blockVariable)
		
		// Доступ к переменной функции из блока
		fmt.Println("   🏠 Доступ к переменной функции из блока:", functionVariable)
	}
	
	// Ошибка! blockVariable не доступна вне блока if
	// fmt.Println(blockVariable) // Эта строка вызовет ошибку компиляции
	
	// Вызов другой функции
	anotherFunction()
}

func anotherFunction() {
	anotherVar := "Я переменная другой функции"
	fmt.Println("🏠 Локальная переменная anotherFunction:", anotherVar)
	
	// Ошибка! functionVariable не доступна в этой функции
	// fmt.Println(functionVariable) // Эта строка вызовет ошибку компиляции
}