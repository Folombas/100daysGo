package main

import (
    "fmt"
    //"os"
)

// PrintToStdout - безопасный вывод кириллицы
func PrintToStdout(messages ...string) {
    for _, msg := range messages {
        fmt.Println(msg)
    }
}