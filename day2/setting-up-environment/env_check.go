package main

import (
    "fmt"
    "os"
    "runtime"
)

func CheckGoEnvironment() {
    fmt.Println("\n🔍 Проверка среды Go:")
    fmt.Println("---------------------")

    fmt.Printf("✅ GOOS: %s\n", runtime.GOOS)
    fmt.Printf("✅ GOARCH: %s\n", runtime.GOARCH)
    fmt.Printf("✅ Version: %s\n", runtime.Version())

    if goPath := os.Getenv("GOPATH"); goPath != "" {
        fmt.Printf("✅ GOPATH: %s\n", goPath)
    }

    if goRoot := os.Getenv("GOROOT"); goRoot != "" {
        fmt.Printf("✅ GOROOT: %s\n", goRoot)
    }
}
