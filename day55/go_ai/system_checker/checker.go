// system_checker/checker.go
package system_checker

import (
	"fmt"
	"runtime"
)

// PrintSystemInfo выводит информацию о системе — для мотивации и понимания контекста
func PrintSystemInfo() {
	fmt.Println("=== 🖥️  Информация о системе ===")
	fmt.Printf("ОС: %s\n", runtime.GOOS)
	fmt.Printf("Архитектура: %s\n", runtime.GOARCH)
	fmt.Printf("Количество CPU: %d\n", runtime.NumCPU())
	fmt.Printf("Go версия: %s\n", runtime.Version())
	fmt.Println("💡 Совет: Go не требует GPU — ты можешь учиться, развиваться и строить backend для ИИ даже на скромном железе!")
}
