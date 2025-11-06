package main

import (
	"fmt"
	"runtime"
	"time"
)

func ShowAdvancedHello() {
	fmt.Println("\n2. 💫 Продвинутый Hello World:")

	// Разные способы вывода
	name := "Gopher"
	currentTime := time.Now()

	// Получаем реальную версию Go
	goVersion := runtime.Version()

	fmt.Printf("   Hello, %s! \n", name)
	fmt.Printf("   Go version: %s \n", goVersion)
	fmt.Printf("   Сегодняшняя дата: %s\n", currentTime.Format("02.01.2006"))
	fmt.Printf("   Время запуска: %s\n", currentTime.Format("15:04:05"))

	// Многострочный вывод
	message := `
   🎉 Поздравляю!
   Ты написал свою первую программу на Go!
   Это только начало великого пути!`
	fmt.Println(message)
}
