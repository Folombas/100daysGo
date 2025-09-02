package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"runtime"
	"strings"
)

// getWindowsDetails возвращает детальную информацию о версии Windows
func getWindowsDetails() string {
	if runtime.GOOS != "windows" {
		return "Windows"
	}

	// Получаем информацию из реестра
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "Windows (неизвестная версия)"
	}
	defer k.Close()

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return "Windows (неизвестная версия)"
	}

	// Очищаем название продукта от лишних слов
	productName = strings.ReplaceAll(productName, "Microsoft", "")
	productName = strings.ReplaceAll(productName, "(R)", "")
	productName = strings.ReplaceAll(productName, "(TM)", "")
	productName = strings.TrimSpace(productName)

	// Получаем номер сборки
	currentBuild, _, err := k.GetStringValue("CurrentBuild")
	if err == nil {
		displayVersion, _, err := k.GetStringValue("DisplayVersion")
		if err == nil {
			return fmt.Sprintf("%s %s (сборка %s)", productName, displayVersion, currentBuild)
		}
		return fmt.Sprintf("%s (сборка %s)", productName, currentBuild)
	}

	return productName
}