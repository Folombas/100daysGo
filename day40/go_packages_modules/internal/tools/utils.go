package tools

import (
    "github.com/fatih/color"
	"fmt"
)

// PrintError выводит сообщение об ошибке красным цветом
func PrintError(message string) {
    red := color.New(color.FgRed).SprintFunc()
    fmt.Printf("%s: %s\n", red("ОШИБКА"), message)
}

// PrintSuccess выводит сообщение об успехе зеленым цветом
func PrintSuccess(message string) {
    green := color.New(color.FgGreen).SprintFunc()
    fmt.Printf("%s: %s\n", green("УСПЕХ"), message)
}

// PrintInfo выводит информационное сообщение синим цветом
func PrintInfo(message string) {
    blue := color.New(color.FgBlue).SprintFunc()
    fmt.Printf("%s: %s\n", blue("ИНФО"), message)
}