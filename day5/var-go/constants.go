package main

import "fmt"

func demoConstants() {
	fmt.Println("🌟 5. КОНСТАНТЫ:")
	fmt.Println("----------------")

	// Базовые константы
	const MaxUsers = 1000
	const AppName = "LearningGo"
	const DefaultTimeout = 30

	fmt.Printf("   MaxUsers = %d\n", MaxUsers)
	fmt.Printf("   AppName = %q\n", AppName)
	fmt.Printf("   DefaultTimeout = %d\n", DefaultTimeout)

	// Типизированные константы
	const Pi float64 = 3.1415926535
	const WelcomeMessage string = "Добро пожаловать в Go!"

	fmt.Printf("   Pi = %.4f\n", Pi)
	fmt.Printf("   WelcomeMessage = %q\n", WelcomeMessage)

	// Множественные константы
	const (
		Success = iota // авто-инкремент
		Failed
		Pending
	)

	fmt.Printf("   Status: Success=%d, Failed=%d, Pending=%d\n", Success, Failed, Pending)

	// Константы с выражениями
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	fmt.Printf("   Размеры: KB=%d, MB=%d, GB=%d\n", KB, MB, GB)
	fmt.Println()
}

func demoBestPractices() {
	fmt.Println("💡 6. ЛУЧШИЕ ПРАКТИКИ:")
	fmt.Println("----------------------")

	// Когда использовать var
	var configFile string // будет установлена позже
	if debugMode {
		configFile = "dev.config"
	} else {
		configFile = "prod.config"
	}
	fmt.Printf("   configFile = %q (объявлен с var)\n", configFile)

	// Когда использовать :=
	counter := 0 // локальная переменная с коротким жизненным циклом
	counter++
	fmt.Printf("   counter = %d (объявлен с :=)\n", counter)

	// Группировка связанных переменных
	var (
		userID      = 12345
		userName    = "john_doe"
		isVerified  = true
		accountType = "premium"
	)

	fmt.Printf("   Пользователь: ID=%d, Name=%q, Verified=%t, Type=%q\n",
		userID, userName, isVerified, accountType)

	fmt.Println("\n🎯 ИТОГИ:")
	fmt.Println("   • var: для глобальных переменных, нулевых значений, отложенной инициализации")
	fmt.Println("   • :=: для локальных переменных с немедленной инициализацией")
	fmt.Println("   • const: для значений, которые не должны изменяться")
	fmt.Println()

	// Демонстрация области видимости
	showScopeDemo()
}

func showScopeDemo() {
	// Локальная переменная - существует только в этой функции
	localVar := "Я локальная переменная"
	fmt.Printf("   📍 %s\n", localVar)
}
