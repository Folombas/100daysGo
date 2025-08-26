package main

import (
	"fmt"
	//"os"
	"os/exec"
	//"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("🛠️  Инструменты разработки Go")
	fmt.Println("==============================")
	fmt.Printf("Версия Go: %s\n", runtime.Version())
	fmt.Printf("Платформа: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()

	// Демонстрация различных инструментов
	demoGoFmt()
	demoGoVet()
	demoGoTest()
	demoGoMod()
	demoGoGet()
	demoGoRun()
	demoGoBuild()

	fmt.Println("\n✅ Все инструменты продемонстрированы!")
	fmt.Println("Подробнее смотрите в файле README.md")
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = "."
	fmt.Printf("$ %s", name)
	for _, a := range arg {
		fmt.Printf(" %s", a)
	}
	fmt.Println()

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Ошибка: %s\n", err)
	} else if len(output) > 0 {
		fmt.Printf("%s\n", string(output))
	}
}