package main

import (
    "fmt"
    "os"
    "os/exec"
)

// GenerateCoverageReport создает HTML-отчет о покрытии
func GenerateCoverageReport() error {
    // Создаем coverage-профиль
    cmd := exec.Command("go", "test", "-coverprofile=coverage.out", "./...")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("тесты не прошли: %w", err)
    }

    // Генерируем HTML отчет
    cmd = exec.Command("go", "tool", "cover", "-html=coverage.out", "-o=coverage.html")
    return cmd.Run()
}