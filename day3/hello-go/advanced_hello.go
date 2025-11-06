package main

import "fmt"

func ShowAdvancedHello() {
	fmt.Println("\n2. 💫 Продвинутый Hello World:")

	// Разные способы вывода
	name := "Gopher"
	version := 1.21

	fmt.Printf("   Hello, %s! \n", name)
	fmt.Printf("   Go version: %.2f \n", float64(version))
	fmt.Printf("   Сегодняшняя дата: 2025 \n")

	// Многострочный вывод
	message := `
   🎉 Поздравляю!
   Ты написал свою первую программу на Go!
   Это только начало великого пути!`
	fmt.Println(message)
}
