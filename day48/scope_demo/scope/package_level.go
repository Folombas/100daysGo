package scope

import "fmt"

// Пакетные переменные (доступны во всем пакете)
var (
	GlobalVariable     = "Я глобальная переменная пакета"
	packageLevelSecret = "Я секретная переменная пакета (не экспортируется)"
)

// ExportedFunction демонстрирует доступ к пакетным переменным
func ExportedFunction() {
	fmt.Println("📦 Из ExportedFunction:")
	fmt.Println("   ", GlobalVariable)
	fmt.Println("   ", packageLevelSecret)
}

// DemoPackageLevelScope демонстрирует область видимости на уровне пакета
func DemoPackageLevelScope() {
	fmt.Println("🌍 Глобальная переменная:", GlobalVariable)
	
	// packageLevelSecret доступна только внутри пакета
	fmt.Println("🔒 Секретная переменная пакета:", packageLevelSecret)
	
	// Вызов функции из того же пакета
	ExportedFunction()
}